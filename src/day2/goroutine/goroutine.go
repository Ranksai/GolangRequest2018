package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func input(r io.Reader) <-chan string {
	// チャネルを作る
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			// チャネルに読み込んだ文字列を送る
			ch <- s.Text()
		}
		// チャネルを閉じる
		close(ch)
	}()
	// チャネルを返す
	return ch
}

func main() {
	ch := input(os.Stdin)
	for {
		fmt.Print(">")
		fmt.Println(<-ch)
	}
}
