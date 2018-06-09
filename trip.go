package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var stupidBig int = 1024 * 1024
	reader := bufio.NewReaderSize(os.Stdin, stupidBig)
	var aScore int32 = 0
	var bScore int32 = 0
	var aLineBytes, _, err1 = reader.ReadLine()
	if err1 != nil {
		panic(err1)
	}
	var aLine = string(aLineBytes)
	aSplit := strings.Split(aLine, " ")
	var bLineBytes, _, err2 = reader.ReadLine()
	if err2 != nil {
		panic(err2)
	}
	var bLine = string(bLineBytes)
	bSplit := strings.Split(bLine, " ")

	for i := 0; i < 3; i++ {
		var aVal, _ = strconv.Atoi(aSplit[i])
		var bVal, _ = strconv.Atoi(bSplit[i])
		if aVal > bVal {
			aScore += int32(1)
		} else if bVal > aVal {
			bScore += int32(1)
		}
	}
	fmt.Printf("%d ", aScore)
	fmt.Printf("%d", bScore)
}
