package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

var (
	hashtable1 = map[string]string{}
	hashtable2 = map[string]string{}
)

type (
	HIndex struct {
		hash string
		i, l int
		c    rune
	}
)

func NewHIndex(i int) HIndex {
	return HIndex{
		i: i,
	}
}

func (h *HIndex) copy() HIndex {
	return HIndex{
		hash: h.hash,
		i:    h.i,
		l:    h.l,
		c:    h.c,
	}
}

func (h *HIndex) genHash(s string) {
	j := s + strconv.Itoa(h.i)
	if val, ok := hashtable1[j]; ok {
		h.hash = val
	} else {
		h.hash = fmt.Sprintf("%x", md5.Sum([]byte(j)))
		hashtable1[j] = h.hash
	}
}

func (h *HIndex) keyStretch(k int) {
	if val, ok := hashtable2[h.hash]; ok {
		h.hash = val
	} else {
		j := h.hash
		for i := 0; i < k; i++ {
			h.hash = fmt.Sprintf("%x", md5.Sum([]byte(h.hash)))
		}
		hashtable2[j] = h.hash
	}
}

func (h *HIndex) hasNTuple(length int) bool {
tupleloop:
	for n := length - 1; n < len(h.hash); n++ {
		ch := h.hash[n]
		for j := length - 1; j > 0; j-- {
			if h.hash[n-j] != ch {
				continue tupleloop
			}
		}
		h.l = length
		h.c = rune(ch)
		return true
	}
	return false
}

func (h *HIndex) sameChar(other *HIndex) bool {
	return h.c == other.c
}

const (
	salt = "ihaygndm"
)

func main() {
	start := time.Now()

	count := 0
	for i := 0; count < 64; i++ {
		k := NewHIndex(i)
		k.genHash(salt)
		if !k.hasNTuple(3) {
			continue
		}
		dist := i + 1001
		for j := i + 1; j < dist; j++ {
			l := NewHIndex(j)
			l.genHash(salt)
			if !l.hasNTuple(5) {
				continue
			}
			if k.sameChar(&l) {
				count++
				fmt.Println(count, k.i, k.hash)
				break
			}
		}
	}

	count = 0
	for i := 0; count < 64; i++ {
		k := NewHIndex(i)
		k.genHash(salt)
		k.keyStretch(2016)
		if !k.hasNTuple(3) {
			continue
		}
		dist := i + 1001
		for j := i + 1; j < dist; j++ {
			l := NewHIndex(j)
			l.genHash(salt)
			l.keyStretch(2016)
			if !l.hasNTuple(5) {
				continue
			}
			if k.sameChar(&l) {
				count++
				fmt.Println(count, k.i, k.hash)
				break
			}
		}
	}

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
