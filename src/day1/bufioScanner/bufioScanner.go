package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 標準入力から読み込む
	scanner := bufio.NewScanner(os.Stdin)
	// 1行ずつ読み込んで繰り返す
	for scanner.Scan() {
		//1行分を出力する
		fmt.Println(scanner.Text())
	}
	// まとめてエラー処理をする
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みに失敗しました:", err)
	}
}
