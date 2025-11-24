package main

import (
	"fmt"
)

// Matrix представляет матрицу
type Matrix struct {
	rows, cols int
	data       [][]float64
}

// NewMatrix создает новую матрицу
func NewMatrix(rows, cols int) *Matrix {
	data := make([][]float64, rows)
	for i := range data {
		data[i] = make([]float64, cols)
	}
	return &Matrix{rows: rows, cols: cols, data: data}
}

// Print выводит матрицу на экран
func (m *Matrix) Print() {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			fmt.Printf("%8.2f", m.data[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

// Set устанавливает значение элемента матрицы
func (m *Matrix) Set(row, col int, value float64) {
	if row >= 0 && row < m.rows && col >= 0 && col < m.cols {
		m.data[row][col] = value
	}
}

// Add складывает две матрицы
func (m *Matrix) Add(other *Matrix) *Matrix {
	if m.rows != other.rows || m.cols != other.cols {
		fmt.Println("Ошибка: матрицы должны быть одного размера")
		return nil
	}

	result := NewMatrix(m.rows, m.cols)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			result.data[i][j] = m.data[i][j] + other.data[i][j]
		}
	}
	return result
}

// Subtract вычитает две матрицы
func (m *Matrix) Subtract(other *Matrix) *Matrix {
	if m.rows != other.rows || m.cols != other.cols {
		fmt.Println("Ошибка: матрицы должны быть одного размера")
		return nil
	}

	result := NewMatrix(m.rows, m.cols)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			result.data[i][j] = m.data[i][j] - other.data[i][j]
		}
	}
	return result
}

// Multiply умножает две матрицы
func (m *Matrix) Multiply(other *Matrix) *Matrix {
	if m.cols != other.rows {
		fmt.Println("Ошибка: количество столбцов первой матрицы должно равняться количеству строк второй")
		return nil
	}

	result := NewMatrix(m.rows, other.cols)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < other.cols; j++ {
			for k := 0; k < m.cols; k++ {
				result.data[i][j] += m.data[i][k] * other.data[k][j]
			}
		}
	}
	return result
}

// ScalarMultiply умножает матрицу на скаляр
func (m *Matrix) ScalarMultiply(scalar float64) *Matrix {
	result := NewMatrix(m.rows, m.cols)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			result.data[i][j] = m.data[i][j] * scalar
		}
	}
	return result
}

// Transpose возвращает транспонированную матрицу
func (m *Matrix) Transpose() *Matrix {
	result := NewMatrix(m.cols, m.rows)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			result.data[j][i] = m.data[i][j]
		}
	}
	return result
}

// Determinant вычисляет определитель матрицы (только для 2x2 и 3x3)
func (m *Matrix) Determinant() float64 {
	if m.rows != m.cols {
		fmt.Println("Ошибка: матрица должна быть квадратной")
		return 0
	}

	if m.rows == 1 {
		return m.data[0][0]
	}

	if m.rows == 2 {
		return m.data[0][0]*m.data[1][1] - m.data[0][1]*m.data[1][0]
	}

	if m.rows == 3 {
		return m.data[0][0]*(m.data[1][1]*m.data[2][2]-m.data[1][2]*m.data[2][1]) -
			m.data[0][1]*(m.data[1][0]*m.data[2][2]-m.data[1][2]*m.data[2][0]) +
			m.data[0][2]*(m.data[1][0]*m.data[2][1]-m.data[1][1]*m.data[2][0])
	}

	fmt.Println("Ошибка: определитель можно вычислить только для матриц 1x1, 2x2 и 3x3")
	return 0
}

// CreateMatrixFromInput создает матрицу из пользовательского ввода
func CreateMatrixFromInput() *Matrix {
	var rows, cols int
	fmt.Print("Введите количество строк: ")
	fmt.Scan(&rows)
	fmt.Print("Введите количество столбцов: ")
	fmt.Scan(&cols)

	matrix := NewMatrix(rows, cols)

	fmt.Printf("Введите элементы матрицы %dx%d:\n", rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("Элемент [%d][%d]: ", i, j)
			var value float64
			fmt.Scan(&value)
			matrix.Set(i, j, value)
		}
	}
	return matrix
}

func main() {
	fmt.Println("=== КАЛЬКУЛЯТОР МАТРИЦ ===")
	fmt.Println("1. Сложение матриц")
	fmt.Println("2. Вычитание матриц")
	fmt.Println("3. Умножение матриц")
	fmt.Println("4. Умножение на скаляр")
	fmt.Println("5. Транспонирование")
	fmt.Println("6. Определитель (до 3x3)")
	fmt.Println("0. Выход")

	for {
		var choice int
		fmt.Print("\nВыберите операцию: ")
		fmt.Scan(&choice)

		switch choice {
		case 0:
			fmt.Println("Выход из программы")
			return

		case 1: // Сложение
			fmt.Println("\n--- Сложение матриц ---")
			fmt.Println("Первая матрица:")
			m1 := CreateMatrixFromInput()
			fmt.Println("Вторая матрица:")
			m2 := CreateMatrixFromInput()
			
			fmt.Println("Первая матрица:")
			m1.Print()
			fmt.Println("Вторая матрица:")
			m2.Print()
			
			result := m1.Add(m2)
			if result != nil {
				fmt.Println("Результат сложения:")
				result.Print()
			}

		case 2: // Вычитание
			fmt.Println("\n--- Вычитание матриц ---")
			fmt.Println("Первая матрица:")
			m1 := CreateMatrixFromInput()
			fmt.Println("Вторая матрица:")
			m2 := CreateMatrixFromInput()
			
			fmt.Println("Первая матрица:")
			m1.Print()
			fmt.Println("Вторая матрица:")
			m2.Print()
			
			result := m1.Subtract(m2)
			if result != nil {
				fmt.Println("Результат вычитания:")
				result.Print()
			}

		case 3: // Умножение матриц
			fmt.Println("\n--- Умножение матриц ---")
			fmt.Println("Первая матрица:")
			m1 := CreateMatrixFromInput()
			fmt.Println("Вторая матрица:")
			m2 := CreateMatrixFromInput()
			
			fmt.Println("Первая матрица:")
			m1.Print()
			fmt.Println("Вторая матрица:")
			m2.Print()
			
			result := m1.Multiply(m2)
			if result != nil {
				fmt.Println("Результат умножения:")
				result.Print()
			}

		case 4: // Умножение на скаляр
			fmt.Println("\n--- Умножение на скаляр ---")
			m := CreateMatrixFromInput()
			var scalar float64
			fmt.Print("Введите скаляр: ")
			fmt.Scan(&scalar)
			
			fmt.Println("Исходная матрица:")
			m.Print()
			fmt.Printf("Скаляр: %.2f\n", scalar)
			
			result := m.ScalarMultiply(scalar)
			fmt.Println("Результат:")
			result.Print()

		case 5: // Транспонирование
			fmt.Println("\n--- Транспонирование ---")
			m := CreateMatrixFromInput()
			
			fmt.Println("Исходная матрица:")
			m.Print()
			
			result := m.Transpose()
			fmt.Println("Транспонированная матрица:")
			result.Print()

		case 6: // Определитель
			fmt.Println("\n--- Определитель ---")
			m := CreateMatrixFromInput()
			
			fmt.Println("Матрица:")
			m.Print()
			
			det := m.Determinant()
			fmt.Printf("Определитель: %.2f\n", det)

		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}