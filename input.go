package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kave08/BMI/info"
)

var reader = bufio.NewReader(os.Stdin)

func getUserMetrics() (weight float64, height float64) {

	weight = getUserInput(info.WeightPrompt)
	height = getUserInput(info.HeightPrompt)

	return
}

func getUserInput(promptText string) (value float64) {
	fmt.Print(promptText)
	usertInput, _ := reader.ReadString('\n')
	usertInput = strings.Replace(usertInput, "\n", "", -1)

	value, _ = strconv.ParseFloat(usertInput, 64)
	return
}
