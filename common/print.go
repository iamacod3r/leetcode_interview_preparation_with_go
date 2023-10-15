package common

import "fmt"

func PrintStr(r, e string) {
	fmt.Printf("Result : %s Expected : %s = %v\n", r, e, (r == e))
}

func PrintInt(r int, e int) {
	fmt.Printf("Result : %d Expected : %d = %v\n", r, e, (r == e))
}

func PrintBool(r bool, e bool) {
	fmt.Printf("Result : %v Expected : %v = %v\n", r, e, (r == e))
}

func PrintIntSlice(r []int, e []int) {
	if len(r) != len(e) {
		fmt.Printf("the sizes are different : %d != %d\n", len(r), len(e))
		return
	}

	size := len(r)
	fmt.Println("Result\t Expected")
	for i := 0; i < size; i++ {
		fmt.Printf("%d\t %d = %v\n", r[i], e[i], (r[i] == e[i]))
	}
}

func PrintStrSlice(r []string, e []string) {
	if len(r) != len(e) {
		fmt.Printf("the sizes are different : %d != %d\n", len(r), len(e))
		return
	}

	size := len(r)
	fmt.Println("Result\t Expected")
	for i := 0; i < size; i++ {
		fmt.Printf("%s\t %s = %v\n", r[i], e[i], (r[i] == e[i]))
	}
}

func Print2DSlice(r [][]int, e [][]int) {
	if len(r) != len(e) || len(r[0]) != len(e[0]) {
		fmt.Printf("the sizes are different : r[%d != %d] - c[%d != %d]\n", len(r), len(e), len(r[0]), len(e[0]))
		return
	}

	for rc := 0; rc < len(r); rc++ {
		for cc := 0; cc < len(r[rc]); cc++ {
			if len(r[rc]) != len(e[rc]) {
				fmt.Printf("%d.row column sizes different : r[%d](%d) != e[%d](%d)\n", rc, rc, len(r[rc]), rc, len(e[rc]))
				return
			}
		}
	}

	trueCount := 0
	falseCount := 0
	rowSize := len(r)
	fmt.Println("Result\t\tExpected")
	for row := 0; row < rowSize; row++ {
		colSize := len(r[row])
		fmt.Print("[")
		for col := 0; col < colSize; col++ {
			fmt.Printf("%d", r[row][col])
			if col < colSize-1 {
				fmt.Print(", ")
			}
		}
		fmt.Print("]\t")
		fmt.Print("[")
		lineTrueCount := 0
		for col := 0; col < colSize; col++ {
			result := r[row][col] == e[row][col]
			if result {
				trueCount++
				lineTrueCount++
			} else {
				falseCount++
			}
			fmt.Printf("%d", e[row][col])
			if col < colSize-1 {
				fmt.Print(", ")
			}
		}
		fmt.Print("]")
		fmt.Printf("\t\t%d.row => %v\n", row+1, lineTrueCount == colSize)
	}

	fmt.Printf("\nFinal Result is %v\n", trueCount > 0 && falseCount == 0)
}

func Print2Dint64(r [][]int64, e [][]int64) {

	big := r
	small := e

	if len(e) > len(r) {
		big = e
		small = r
	}

	rowSize := len(big)
	sr := len(small)
	sc := len(small[0])
	fmt.Println("[][]int64")
	for row := 0; row < rowSize; row++ {
		colSize := len(big[row])
		println("-----")
		if row < sr {
			for col := 0; col < colSize; col++ {
				if col < sc {
					fmt.Printf("%d\t %d = %v\n", big[row][col], small[row][col], (big[row][col] == small[row][col]))
				} else {
					fmt.Printf("%d\n", big[row][col])
				}

			}
		} else {
			for col := 0; col < colSize; col++ {
				fmt.Printf("%d\n", big[row][col])
			}
		}
	}
	fmt.Println("")
}

func PrintlnStr(msg string) {
	fmt.Println(msg)
}
