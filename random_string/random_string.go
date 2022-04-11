package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var qua, input2 int
	var input3 string
	fmt.Printf("輸入產生字元數：")
	fmt.Scanln(&qua)
	fmt.Printf("輸入產生數量：")
	fmt.Scanln(&input2)
	fmt.Printf("輸入前置字串：")
	fmt.Scanln(&input3)
	fmt.Printf("\n")
	for a := 0; a < input2; a++ {
		fmt.Println(input3 + GetRandomString(qua))
	}
}

func GetRandomString(l int) string {
	str := "1234567890abcdf"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
