package main

import (
	"fmt"
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

func main() {
	mat1 := matrix{100, 100, populate(100, 100, 1)}
	mat2 := matrix{100, 100, populate(100, 100, 1)}
	answ := matrix{100, 100, populate(100, 100, 0)}

	//mat1.print()
	//mat2.print()

	fmt.Println("NORMAL MULTIPLY")
	start := time.Now()
	answ.mat = mat1.multiply(mat2)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Time Elapsed Normal: %v\n", elapsed)
	//answ.print()

	fmt.Println("PARALLEL MULTIPLY")
	start = time.Now()
	answ.mat = mat1.pmultiply(mat2)
	t = time.Now()
	elapsed = t.Sub(start)
	fmt.Printf("Time Elapsed Normal: %v\n", elapsed)
	//answ.print()
}
