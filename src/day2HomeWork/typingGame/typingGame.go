package typingGame

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

const TIMEOUT = 10

var words = []string{"dog", "cat"}

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}

func timeout() <-chan string {
	ch := make(chan string)
	go func() {
		time.Sleep(TIMEOUT * time.Second)
		ch <- "timeout"
	}()
	return ch
}

func printQuiz() string {
	fmt.Printf("Word: %s\n", words[0])
	return words[0]
}

func TypingGame() {

	var score int
	ch := input(os.Stdin)
	chTime := timeout()
	for {
		quiz := printQuiz()
		fmt.Print(">")
		select {
		// receive typed string
		case receive := <-ch:
			if quiz == receive {
				fmt.Println("correct")
				score++
			} else {
				fmt.Println("not correct")
			}
			// timeout typing
		case <-time.After(10 * time.Second):
			fmt.Println("not correct")
			// timeout game
		case timeout := <-chTime:
			fmt.Printf("\n%s\n", timeout)
			fmt.Printf("your score is %d !!!\n", score)
			return
		}
	}
}
