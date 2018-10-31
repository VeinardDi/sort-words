package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		return
	}
	defer file.Close()

	size, err := file.Stat()
	if err != nil {
		return
	}

	array := make([]byte, size.Size())
	_, err = file.Read(array)
	if err != nil {
		return
	}

	str := string(array)
	fmt.Println(str)
	re := regexp.MustCompile("[a-zA-Z]+")
	str2 := re.FindAllString(str, -1)
	sort.Strings(str2)
	fmt.Println(str2)

	output, err:= os.Create("outout.txt")
	if err != nil {
		return
	}
	output.WriteString(strings.Join(str2, ", "))
	defer output.Close()
