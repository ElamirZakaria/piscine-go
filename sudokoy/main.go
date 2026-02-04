package main

import (
	"fmt"
	"os"
)

func SplitWhiteSpaces(s string) []string {
	var answer []string
	str := ""
	flag := 0
	for _, a := range s {
		if ((a == ' ') || (a == '\n') || (a == '\t')) && (flag == 1) {
			answer = append(answer, str)
			str = ""
			flag = 0

		} else {
			if a != ' ' {
				str = str + string(a)

				flag = 1
			}
		}
	}
	if str != "" {
		answer = append(answer, str)
	}
	return answer
}

func IsInputValid() bool {
	NumberOfInputs := len(os.Args)

	if NumberOfInputs != 10 {
		return false
	}
	for i := 1; i < 10; i++ {
		row := os.Args[i]
		if len(row) != 9 {
			return false
		}
		for _, a := range row {
			if (a == '.') || a >= '1' && a <= '9' {
			} else {
				return false
			}
		}

	}

	return true
}

// ده جدول السودوكو 9 صفوف × 9 أعمدة
// كل خانة فيها رقم من 1 لـ 9
// ولو الخانة فاضية بنحط 0
var grid [9][9]int

// دالة بتطبع الجدول كله
func printGrid() {
	// لف على الصفوف
	for i := 0; i < 9; i++ {
		// لف على الأعمدة جوه كل صف
		for j := 0; j < 9; j++ {
			// اطبع الرقم اللي في الخانة
			fmt.Print(grid[i][j], " ")
		}
		// انزل سطر بعد ما تخلص الصف
		fmt.Println()
	}
}

func isValid(row, col, num int) bool {
	// أول حاجة: نراجع الصف والعمود
	for i := 0; i < 9; i++ {
		// لو الرقم موجود في نفس الصف
		// أو موجود في نفس العمود
		// يبقى مينفعش
		if grid[row][i] == num || grid[i][col] == num {
			return false
		}
	}

	// نحدد أول صف وأول عمود في الـ box (3×3)
	// مثال: لو row = 5 يبقى 5/3 = 1 → 1*3 = 3
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3

	// نلف جوه البوكس 3×3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			// لو الرقم موجود في البوكس
			// يبقى مينفعش
			if grid[i][j] == num {
				return false
			}
		}
	}

	// لو عدّى من كل الفحوصات يبقى ينفع
	return true
}

// دالة الحل الأساسية (Backtracking)
func solve() bool {
	// لف على كل خانات الجدول
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// لو الخانة فاضية
			if grid[i][j] == 0 {

				// جرّب تحط رقم من 1 لـ 9
				for num := 1; num <= 9; num++ {
					// لو الرقم ينفع في المكان ده
					if isValid(i, j, num) {

						// حطه مؤقتًا
						grid[i][j] = num

						// حاول تكمل حل باقي الجدول
						if solve() {
							return true
						}

						// لو فشلنا نرجع الخانة فاضية تاني
						grid[i][j] = 0
					}
				}

				// لو ولا رقم نفع في الخانة دي
				// يبقى الحل اللي ماشيين فيه غلط
				return false
			}
		}
	}

	// لو مفيش خانات فاضية يبقى اتحل
	return true
}

func main() {
	// ده الإدخال اللي جاي من الكوماند لاين
	// النقطة معناها خانة فاضية
	input := os.Args[1:]
	err := IsInputValid()
	if err != true {
		fmt.Println("Error")
	} else {
		// نحول الـ input ده لجدول أرقام
		for i, line := range input {
			for j, ch := range line {
				// لو لقيت نقطة يبقى الخانة فاضية
				if ch == '.' {
					grid[i][j] = 0
				} else {
					// لو رقم نحوله من char لـ int
					grid[i][j] = int(ch - '0')
				}
			}
		}

		// حاول تحل السودوكو
		if solve() {
			// لو اتحل اطبع الجدول
			printGrid()
			fmt.Println()
		} else {
			// لو ملوش حل
			fmt.Println("Error")
		}
	}
}
