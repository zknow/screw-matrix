package main

import (
	"fmt"
	"strconv"

	screwmatrix "github.com/zknow/screw-matrix"
)

type Oprion struct {
	Capacity  int
	ClockWise screwmatrix.ClockWise
}

func main() {
	num := GetMatrixNumInput()
	options := GetMatrixOption(num)
	for _, opiton := range options {
		screwmatrix.Execute(opiton.ClockWise, opiton.Capacity)
	}
}

func GetMatrixOption(num int) []Oprion {
	var capStr, clockWiseStr string
	var options []Oprion
LOOP:
	for i := 0; i < num; i++ {
		fmt.Println("請輸入矩陣大小&移動方向(1順時針，2逆時針):")
		fmt.Scanln(&capStr, &clockWiseStr)

		cap, err := strconv.Atoi(capStr)
		if err != nil {
			goto LOOP
		}
		wise, err := strconv.Atoi(clockWiseStr)
		if err != nil {
			goto LOOP
		}

		var clockWise screwmatrix.ClockWise
		if wise == 1 {
			clockWise = screwmatrix.RightClockWise
		} else if wise == 2 {
			clockWise = screwmatrix.LeftClockWise
		} else {
			goto LOOP
		}
		options = append(options, Oprion{Capacity: cap, ClockWise: clockWise})
	}

	return options
}

func GetMatrixNumInput() int {
	var inputNumStr string
	var num int
	for {
		fmt.Println("請輸入Matrix數量:")
		fmt.Scanln(&inputNumStr)
		i, err := strconv.Atoi(inputNumStr)
		if err != nil || i <= 0 {
			fmt.Println("數量輸入錯誤:")
		} else {
			num = i
			break
		}
	}
	return num
}
