package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	file_name = "input.txt"
)

func decompress(s string) int {
	uncompressed := 0
	for len(s) > 0 {
		k := strings.SplitN(s, "(", 2)
		if len(k) > 1 {
			uncompressed += len(k[0])
			k = strings.SplitN(k[1], ")", 2)
			marker := strings.Split(k[0], "x")
			num, _ := strconv.Atoi(marker[0])
			rep, _ := strconv.Atoi(marker[1])
			s = k[1][num:]
			uncompressed += num * rep
		} else {
			uncompressed += len(k[0])
			s = ""
		}
	}
	return uncompressed
}

func decompressv2(s string) int {
	uncompressed := 0
	for len(s) > 0 {
		k := strings.SplitN(s, "(", 2)
		if len(k) > 1 {
			uncompressed += len(k[0])
			k = strings.SplitN(k[1], ")", 2)
			marker := strings.Split(k[0], "x")
			num, _ := strconv.Atoi(marker[0])
			rep, _ := strconv.Atoi(marker[1])
			sub := decompressv2(k[1][:num])
			s = k[1][num:]
			uncompressed += rep * sub
		} else {
			uncompressed += len(k[0])
			s = ""
		}
	}
	return uncompressed
}

func main() {
	start := time.Now()

	f, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	uncompressed := 0
	uncompressedv2 := 0

	for s.Scan() {
		uncompressed += decompress(s.Text())
		uncompressedv2 += decompressv2(s.Text())
	}
	// uncompressedv2 += decompressv2("(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN")

	fmt.Println("uncompressed length: ", uncompressed)
	fmt.Println("uncompressedv2 length: ", uncompressedv2)

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
