package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	var input string
	ans := randSeq(4, false)
	log.Println(ans)
	for {
		fmt.Scan(&input)
		if result := play1A2B(input, ans); result == "4A0B" {
			break
		} else {
			log.Println(result)
		}
	}
}

// play1A2B 1A2B
func play1A2B(input string, ans string) (result string) {
	var a, b int
	var abMap = make(map[int]struct{})
	// A
	for i := 0; i < len(input); i++ {
		if input[i] == ans[i] {
			a++
			abMap[i] = struct{}{}
		}
	}
	// B
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(ans); j++ {
			if i == j {
				continue
			} else {
				if input[i] == ans[j] {
					_, ok := abMap[j]
					if !ok {
						b++
						abMap[j] = struct{}{}
						break
					}

				}
			}
		}
	}
	result = fmt.Sprintf("%dA%dB", a, b)
	return
}

// randSeq 產生隨機長度為4字串
func randSeq(n int, repeatable bool) string {
	var letters = []rune("0123456789")
	b := make([]rune, n)
	if repeatable { // 允許重複數字
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
	} else {
		for i := range b {
			letterIndex := rand.Intn(len(letters))
			b[i] = letters[letterIndex]
			letters = append(letters[0:letterIndex], letters[letterIndex+1:]...)
		}
	}
	return string(b)
}
