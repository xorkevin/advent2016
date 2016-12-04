package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	file_name = "input.txt"
)

func checkRoom(s string) int {
	r := strings.Split(s, "[")
	room := r[0]
	temp := strings.Split(room, "-")
	id, _ := strconv.Atoi(temp[len(temp)-1])
	checksum := strings.Trim(r[1], "]")

	commonChars := [26]string{}
	for i := 0; i < 26; i++ {
		commonChars[i] = fmt.Sprintf("%02d", strings.Count(room, string('a'+i))) + string('z'-i)
	}
	sort.Strings(commonChars[:])
	l := len(commonChars)

	a := string('z' - commonChars[l-1][2] + 'a')
	b := string('z' - commonChars[l-2][2] + 'a')
	c := string('z' - commonChars[l-3][2] + 'a')
	d := string('z' - commonChars[l-4][2] + 'a')
	e := string('z' - commonChars[l-5][2] + 'a')

	chars := [5]string{a, b, c, d, e}

	cs := strings.Join(chars[:], "")

	if cs == checksum {

		k := temp[0 : len(temp)-1]

		name := []string{}

		for _, i := range k {
			word := []rune{}
			for _, j := range i {
				word = append(word, rune('a'+(int(j)-'a'+id)%26))
			}
			name = append(name, string(word))
		}

		roomname := strings.Join(name, " ")

		if strings.Contains(roomname, "north") {
			fmt.Println(roomname, id)
		}

		return id
	} else {
		return 0
	}
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

	sectorIds := 0

	for s.Scan() {
		sectorIds += checkRoom(s.Text())
	}

	fmt.Println(sectorIds)

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
