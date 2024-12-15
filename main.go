package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
)

var (
	lineCounter int
	mutex       sync.Mutex
)

func handleStr(threadNum int, in <-chan string, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case str, ok := <-in:
			if !ok {
				return
			}

			runes := []rune(str)
			var sb strings.Builder
			for i := len(runes) - 1; i >= 0; i-- {
				sb.WriteRune(runes[i])
			}

			mutex.Lock()
			lineCounter++
			fmt.Printf("line %d, thread %d: \"%s\"\n", lineCounter, threadNum, sb.String())
			mutex.Unlock()
		}
	}
}

func main() {
	// if len(os.Args) < 2 {
	//	fmt.Println("Usage: go run prog.go <string1> <string2> ...")
	//	return
	// }
	// stringsArray := os.Args[1:] 
	// stringsCh := make(chan string)
	
	stringsArray := []string{"Hello", "qwerty", "Golang", "platypus", "тест", "level", "generics"}
	stringsCh := make(chan string)

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func(threadNum int) {
			defer wg.Done()
			handleStr(threadNum, stringsCh, ctx)
		}(i + 1)
	}

	for _, str := range stringsArray {
		stringsCh <- str
	}
	close(stringsCh)
	cancel()

	wg.Wait()

	fmt.Println("done")
}
