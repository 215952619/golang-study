package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type jjj struct {
	User string `json:"user"`
	Status string `json:"status"`
}

func main() {
	var j jjj
	filename := "test.json"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("load err: ", err)
	}
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		fmt.Println("load stat err: ", err)
	}

	data := make([] byte, stats.Size())
	count, err := file.Read(data)
	if err != nil {
		fmt.Println("read err: ", err)
	}
	fmt.Printf("read file %s success, len: %d, content: %s\n", filename, count, data)
	err2 := json.Unmarshal(data, &j)
	if err != nil {
		fmt.Println("json parse err: ", err2)
	}
	fmt.Println(j)
	fmt.Printf("user: %s, status: %s", j.User, j.Status)
}
