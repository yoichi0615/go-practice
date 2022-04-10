package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	arg :=flag.Arg(0) //コマンドライン引数を受け取って新しい変数に入れている
	fmt.Printf("Hello %s\n", arg)
}

