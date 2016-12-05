package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

const (
	input = "cxdnnyjw"
)

func checkPass(pass []byte) bool {
	for _, i := range pass {
		if i == 0 {
			return true
		}
	}
	return false
}

func checkHash(hash string) bool {
	for j := 0; j < 5; j++ {
		if hash[j] != '0' {
			return false
		}
	}
	return true
}

func main() {
	start := time.Now()

	password := ""
	passChan := make(chan string, 8)
	passDone := make(chan bool)
	iterations := 0

	go func() {
		j := 0
		i := 0
		for i = 0; j < 8; i++ {
			hash := fmt.Sprintf("%x", md5.Sum([]byte(input+strconv.Itoa(i))))
			if checkHash(hash) {
				passChan <- string(hash[5])
				j++
			}
		}
		iterations = i
		close(passChan)
	}()

	go func() {
		for i := range passChan {
			password += i
			fmt.Println(password)
		}
		passDone <- true
	}()

	password2 := [8]byte{0, 0, 0, 0, 0, 0, 0, 0}
	pass2Done := make(chan bool)
	iterations2 := 0
	go func() {
		i := 0
		for i = 0; checkPass(password2[:]); i++ {
			hash := fmt.Sprintf("%x", md5.Sum([]byte(input+strconv.Itoa(i))))
			if checkHash(hash) {
				k := hash[5] - '0'
				if k < 8 && password2[k] == 0 {
					password2[k] = hash[6]
					fmt.Println(fmt.Sprintf("%c", password2))
				}
			}
		}
		iterations2 = i
		pass2Done <- true
	}()

	<-passDone
	<-pass2Done

	fmt.Println(fmt.Sprintf("\npassword: %s in %d iterations", password, iterations))
	fmt.Println(fmt.Sprintf("\npassword2: %c in %d iterations", password2, iterations2))

	fmt.Println(fmt.Sprintf("\ntime elapsed: %s", time.Since(start)))
}
