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
	i := 0

	for i = 0; len(password) < 8; i++ {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(input+strconv.Itoa(i))))
		if checkHash(hash) {
			password += string(hash[5])
			fmt.Println(hash)
		}
	}

	password2 := [8]byte{0, 0, 0, 0, 0, 0, 0, 0}
	j := 0

	for j = 0; checkPass(password2[:]); j++ {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(input+strconv.Itoa(j))))
		if checkHash(hash) {
			k := hash[5] - '0'
			if k < 8 && password2[k] == 0 {
				password2[k] = hash[6]
				fmt.Println(hash)
			}
		}
	}

	fmt.Println(fmt.Sprintf("\npassword: %s in %d iterations", password, i))
	fmt.Println(fmt.Sprintf("\npassword2: %c in %d iterations", password2, j))

	fmt.Println(fmt.Sprintf("\ntime elapsed: %s", time.Since(start)))
}
