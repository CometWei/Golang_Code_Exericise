package main

import (
	"fmt"
	"os"
)

var (
	input1, input2 int
	input3, input4 string
)

func main() {
	fmt.Printf("輸入需要產生的數字數量（Ex:0~??）：")
	fmt.Scanln(&input1)
	fmt.Printf("輸入需要補幾位數的0（Ex:1以上）：")
	fmt.Scanln(&input2)
	fmt.Printf("輸入前置字串：")
	fmt.Scanln(&input3)
	fmt.Printf("輸入後置字串：")
	fmt.Scanln(&input4)
	wordlist(input1, input2, input3, input4)
}

func wordlist(input1, input2 int, input3, input4 string) int {
	i := 0
	file, err := os.OpenFile("output.txt", os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return 1
	}
	defer file.Close()
	for i <= input1 {
		out := fmt.Sprintf("%s%0*d%s\n", input3, input2, i, input4)
		file, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("open file failed, err:", err)
			return 1
		}
		file.WriteString(out)
		fmt.Printf("%s", out)
		defer file.Close()
		i++
	}
	return 0
}
