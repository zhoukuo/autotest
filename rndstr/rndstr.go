package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func RandomCharacter(count int) string {
	var chars []rune
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < count; i++ {
		if r.Intn(2) == 0 {
			// 中文字符
			chars = append(chars, rune(r.Intn(20902)+0x4E00))
		} else {
			// ASCII字符
			chars = append(chars, rune(r.Intn(94)+33))
		}
	}
	return string(chars)
}

func RandomCharacter2(count int) string {
	var chars bytes.Buffer
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < count; i++ {
		if r.Intn(2) == 0 {
			// 中文字符
			chars.WriteRune(rune(r.Intn(20902) + 0x4E00))

		} else {
			// ASCII字符
			chars.WriteRune(rune(r.Intn(94) + 33))
		}
	}
	fmt.Println(chars.String())
	return chars.String()
}

func main() {
	start := time.Now()
	RandomCharacter(2000000)
	end := time.Now()
	// Elapsed Time: 160.0091ms
	fmt.Println("Elapsed Time:", end.Sub(start))

	start2 := time.Now()
	RandomCharacter2(2000000)
	end2 := time.Now()
	// Elapsed Time: 145.0082ms
	fmt.Println("Elapsed Time:", end2.Sub(start2))
}
