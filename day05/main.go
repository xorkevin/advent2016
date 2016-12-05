package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"sync"
	"time"
)

type (
	Vec2 struct {
		X byte
		Y byte
	}
)

const (
	input = "cxdnnyjw"
)

func checkHash(hash string) bool {
	for j := 0; j < 5; j++ {
		if hash[j] != '0' {
			return false
		}
	}
	return true
}

func numGen(done <-chan bool) <-chan int {
	send := make(chan int, 128)
	go func() {
		defer close(send)
		for i := 0; true; i++ {
			select {
			case <-done:
				return
			case send <- i:
			}
		}
	}()
	return send
}

func hashWorker(done <-chan bool, receive <-chan int, send chan<- Vec2) {
	for i := range receive {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(input+strconv.Itoa(i))))
		if checkHash(hash) {
			select {
			case <-done:
				return
			case send <- Vec2{hash[5], hash[6]}:
			}
		}
	}
}

func multiSendVec(done <-chan bool, receive <-chan Vec2, send ...chan<- Vec2) {
	for i := range receive {
		for _, j := range send {
			select {
			case <-done:
				return
			case j <- i:
			}
		}
	}
}

func checkPass(pass []byte) bool {
	for _, i := range pass {
		if i == 0 {
			return true
		}
	}
	return false
}

func main() {
	start := time.Now()

	done := make(chan bool)

	n := numGen(done)
	hashChan := make(chan Vec2, 128)
	for i := 0; i < 4; i++ {
		go hashWorker(done, n, hashChan)
	}

	passSend := make(chan Vec2, 128)
	pass2Send := make(chan Vec2, 128)

	go multiSendVec(done, hashChan, passSend, pass2Send)

	var wg sync.WaitGroup

	password := ""
	wg.Add(1)
	go func() {
		defer wg.Done()
		j := 0
		for j = 0; j < 8; j++ {
			hash := <-passSend
			password += string(hash.X)
			fmt.Println(password)
		}
	}()

	password2 := [8]byte{0, 0, 0, 0, 0, 0, 0, 0}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for checkPass(password2[:]) {
			hash := <-pass2Send
			k := hash.X - '0'
			if k < 8 && password2[k] == 0 {
				password2[k] = hash.Y
				fmt.Println(fmt.Sprintf("%c", password2))
			}
		}
	}()

	wg.Wait()
	close(done)

	fmt.Println(fmt.Sprintf("\npassword: %s", password))
	fmt.Println(fmt.Sprintf("\npassword2: %c", password2))

	fmt.Println(fmt.Sprintf("\ntime elapsed: %s", time.Since(start)))
}
