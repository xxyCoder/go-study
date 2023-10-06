package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var firstname, lastname string
	// 标准输入流读写
	fmt.Println("Please enter your full name:")
	// 回车表示输入完毕
	fmt.Scanln(&firstname, &lastname)
	// fmt.Scanf("%s %s", &firstname, &lastname)
	fmt.Println(firstname, lastname)

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter some input:")
	input, err := inputReader.ReadString(' ')
	if err == nil {
		fmt.Println(input)
	}

	// 文件读写
	inputFile, inputErr := os.Open("readme.md")
	if inputErr == nil {
		fmt.Println(inputFile)
		inputReader = bufio.NewReader(inputFile)
		for {
			inputString, readError := inputReader.ReadString('\n')
			fmt.Println(inputString)
			if readError == io.EOF {
				break
			}
		}
	}
	defer inputFile.Close()
}
