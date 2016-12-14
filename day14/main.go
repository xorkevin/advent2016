package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
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
	h.hash = fmt.Sprintf("%x", md5.Sum([]byte(s+strconv.Itoa(h.i))))
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

func (h *HIndex) within(dist int, other *HIndex) bool {
	return other.i-h.i <= dist
}

func (h *HIndex) lessThanEq(other *HIndex) bool {
	return h.i <= other.i
}

func (h *HIndex) sameChar(other *HIndex) bool {
	return h.c == other.c
}

func IndexGen(s string, stopchan <-chan bool) <-chan HIndex {
	outchan := make(chan HIndex, 128)
	go func() {
		defer close(outchan)
		i := 0
		for {
			k := NewHIndex(i)
			k.genHash(s)
			select {
			case outchan <- k:
				i++
			case <-stopchan:
				return
			}
		}
	}()
	return outchan
}

func SearchTup35(inchan <-chan HIndex, stopchan <-chan bool) (outchan3, outchan5 <-chan HIndex) {
	out3 := make(chan HIndex, 128)
	out5 := make(chan HIndex, 128)
	go func() {
		defer close(out3)
		defer close(out5)
		for i := range inchan {
			if i.hasNTuple(3) {
				select {
				case out3 <- i.copy():
				case <-stopchan:
					return
				}
			}
			if i.hasNTuple(5) {
				select {
				case out5 <- i.copy():
				case <-stopchan:
					return
				}
			}
		}
	}()
	return out3, out5
}

func FindHashes(dist int, inchan3, inchan5 <-chan HIndex, stopchan <-chan bool) <-chan HIndex {
	outchan := make(chan HIndex, 64)
	go func() {
		defer close(outchan)
		list5 := []HIndex{}
		for i := range inchan3 {
			markfordelete := 0
			notFound := true
			for n, j := range list5 {
				if j.lessThanEq(&i) {
					markfordelete = n
				} else if i.sameChar(&j) && i.within(dist, &j) {
					select {
					case outchan <- i:
					case <-stopchan:
					}
					notFound = false
					break
				}
			}
			if markfordelete > 0 {
				list5 = list5[markfordelete+1:]
			}
			if notFound {
				for j := range inchan5 {
					if i.within(dist, &j) {
						if i.sameChar(&j) {
							select {
							case outchan <- i:
							case <-stopchan:
							}
							break
						}
					} else {
						break
					}
				}
			}
		}
	}()
	return outchan
}

const (
	salt = "ihaygndm"
)

func main() {
	start := time.Now()

	fmt.Println(fmt.Sprintf("time elapsed: %s", time.Since(start)))
}
