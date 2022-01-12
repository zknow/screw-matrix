package screwmatrix

import "fmt"

type ClockWise int

const (
	RightClockWise ClockWise = iota
	LeftClockWise
)

type SMatrix struct {
	Matrix   [][]int
	CWise    ClockWise
	Capacity int
	endVal   int
}

func Execute(clockWise ClockWise, cap int) {
	s := &SMatrix{
		CWise:    clockWise,
		Capacity: cap,
		endVal:   cap * cap,
	}
	s.initCapacity()
	s.initMoveWays()
	s.move()
	s.printMatrix()
}

func (s *SMatrix) initCapacity() {
	matrix := make([][]int, s.Capacity)
	for i := 0; i < s.Capacity; i++ {
		matrix[i] = make([]int, s.Capacity)
	}
	s.Matrix = matrix
}

func (s *SMatrix) printMatrix() {
	for _, row := range s.Matrix {
		for _, val := range row {
			fmt.Print(val, "\t")
		}
		fmt.Println()
	}
	fmt.Println()
}
