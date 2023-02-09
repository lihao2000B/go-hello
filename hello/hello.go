package main

import (
    "fmt"

    "github.com/lihao2000B/go-hello.git/greetings"
)

func main() {
    // 获取问候信息并打印出来.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}