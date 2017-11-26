package main

import "fmt"
import "github.com/mikespokefire/bazel-go-hello-world/utils"

func main() {
	greeting := utils.Hello()
	fmt.Println(greeting)
}
