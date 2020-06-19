package main

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

type matrix struct {
	row, col int
	mat      [][]float64
}

// method pointer receiver
func (m *matrix) newMatrix(r int, c int) {
	m.row = r
	m.col = c
	matemp := [][]float64{}

	for i := 0; i < m.row; i++ {
		tmp := []float64{}
		for j := 0; j < m.col; j++ {
			tmp = append(tmp, 1.11)
		}
		matemp = append(matemp, tmp)
	}

	m.mat = matemp
}

//method value receiver
func (m matrix) print() {

	for i := range m.mat {
		fmt.Println(m.mat[i])
	}
}

//normal return
func (m matrix) multiply(other matrix) [][]float64 {

	answer := [][]float64{}
	answer = populate(m.row, other.col, 0)

	for i := 0; i < m.row; i++ {
		for j := 0; j < other.col; j++ {
			for k := 0; k < m.col; k++ {
				answer[i][j] += m.mat[i][k] * other.mat[k][j]
			}
		}
	}
	return answer
}

//range, gorouitne
func (m matrix) pmultiply(other matrix) [][]float64 {

	var wg sync.WaitGroup
	result := [][]float64{}
	result = populate(m.row, other.col, 0)

	for i := range m.mat {
		wg.Add(1)
		go func(i int) {
			for j := range other.mat[0] {
				result[i][j] = 0
				for k := range other.mat {
					result[i][j] += m.mat[i][k] * other.mat[k][j]
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return result
}

//return using named values
func populate(r int, c int, p float64) (m [][]float64) {

	for i := 0; i < r; i++ {
		tmp := []float64{}
		for j := 0; j < c; j++ {
			tmp = append(tmp, p)
		}
		m = append(m, tmp)
	}
	return
}

func maker(row, col, row2, col2 int, p float64) (matrix, matrix, matrix, error) {

	mat := matrix{1, 1, populate(1, 1, 1)}

	if row == 0 || col == 0 || row2 == 0 || col2 == 0 {
		return mat, mat, mat, errors.New("0 for size not allowed")
	}
	if col != row2 {
		return mat, mat, mat, errors.New("Column of Matrix 1 and Row of Matrix 2 must be Equal")
	}
	mat1 := matrix{row, col, populate(row, col, p)}
	mat2 := matrix{row2, col2, populate(row2, col2, p)}
	answ := matrix{row, col2, populate(row, col2, 0)}

	return mat1, mat2, answ, nil
}

func main() {

	var row, col, row2, col2 int
	var p float64
	var inp string

	for true {

		fmt.Print("Enter First Matrix Row: ")
		fmt.Scanln(&row)
		fmt.Print("Enter First Matrix Column: ")
		fmt.Scanln(&col)
		fmt.Print("Enter Second Matrix Row: ")
		fmt.Scanln(&row2)
		fmt.Print("Enter Second Matrix Column: ")
		fmt.Scanln(&col2)
		fmt.Print("Enter Number to populate with: ")
		fmt.Scanln(&p)

		mat1, mat2, answ, err := maker(row, col, row2, col2, p)
		if err != nil {
			fmt.Println("Error")
			os.Exit(3)
		}
		fmt.Println("NORMAL MULTIPLY")
		start := time.Now()
		answ.mat = mat1.multiply(mat2)
		t := time.Now()
		elapsed := t.Sub(start)
		fmt.Printf("Time Elapsed Normal: %v\n", elapsed)

		if row < 10 && col < 10 {
			answ.print()
		}

		fmt.Println("PARALLEL MULTIPLY")
		start = time.Now()
		answ.mat = mat1.pmultiply(mat2)
		t = time.Now()
		elapsed = t.Sub(start)
		fmt.Printf("Time Elapsed Normal: %v\n", elapsed)

		if row < 10 && col < 10 {
			answ.print()
		}

		fmt.Printf("Enter New Matrix? Y/n: ")
		fmt.Scan(&inp)
		if inp == "n" {
			os.Exit(1)
		}

	}

}
