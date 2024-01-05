package main

import (
	"fmt"
	"math/rand"
	"time"
)

var arr [4][4]int

func Creatnumber() {
	rand.Seed(time.Now().UnixNano())
	for {
		x := rand.Intn(4)
		y := rand.Intn(4)
		if arr[x][y] == 0 {
			arr[x][y] = TwoorFour(1)
			break
		} else {
			continue
		}
	}
	return
}
func TwoorFour(a int) int {
	rand.Seed(time.Now().UnixNano())
	a = rand.Intn(100)
	if a%10 == 0 {
		a = 4
	} else {
		a = 2
	}
	return a
}
func main() {
	rand.Seed(time.Now().UnixNano())
	k := 0
	for k < 2 {
		x := rand.Intn(4)
		y := rand.Intn(4)
		if arr[x][y] == 0 {
			arr[x][y] = TwoorFour(1)
		}
		k++
	}
	fmt.Println("---------------------------")
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			fmt.Print(arr[i][j], "\t")
		}
		fmt.Println()
	}
	fmt.Println("---------------------------")
	found := false
	for {
		Action()

		Creatnumber()
		Print()

		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr); j++ {
				if arr[i][j] == 2048 {
					found = true
					break
				}
			}
		}
		if found {
			fmt.Println("恭喜你,游戏胜利!")
			break
		}

	}
}
func Print() {
	fmt.Println("---------------------------")
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			fmt.Print(arr[i][j], "\t")
		}
		fmt.Println()
	}
	fmt.Println("---------------------------")
}
func Left90() {
	for i, line := range arr {
		for j, number := range line {
			arr[len(line)-j-1][i] = number
		}
	}
}
func Right90() {
	for i, line := range arr {
		for j, number := range line {
			arr[j][len(line)-i-1] = number
		}
	}
}
func Up() {
	for m := 0; m < 3; m++ {
		for j := 0; j < 4; j++ {
			for i := 0; i < 3; i++ {
				if arr[i][j] == 0 {
					arr[i][j] = arr[i+1][j]
					arr[i+1][j] = 0
				}
			}
		}
	}
}
func merge() {
	for j := 0; j < 4; j++ {
		for i := 0; i < 3; i++ {
			if arr[i][j] == 0 {
				arr[i][j] = arr[i+1][j]
				arr[i+1][j] = 0
			} else if arr[i][j] == arr[i+1][j] {
				arr[i][j] = arr[i][j] + arr[i+1][j]
				arr[i+1][j] = 0
			}
		}
	}
}
func mergeRight() {
	Left90()
	Up()
	merge()
	Right90()
}
func mergeLeft() {
	Right90()
	Up()
	merge()
	Left90()
}
func mergeDown() {
	Right90()
	Right90()
	Up()
	merge()
	Left90()
	Left90()
}
func mergeUp() {
	Up()
	merge()
}
func Action() {
	for {
		var action byte
		fmt.Println("请输入你接下来的操作(wsad)")
		fmt.Scanf("%c\n", &action)
		switch action {
		case 'w':
			if canPerformUp() {
				mergeUp()
				return
			} else {
				fmt.Println("无法执行该操作")
			}
		case 's':
			if canPerformDown() {
				mergeDown()
				return
			} else {
				fmt.Println("无法执行该操作")
			}
		case 'a':
			if canPerformLeft() {
				mergeLeft()
				return
			} else {
				fmt.Println("无法执行该操作")
			}

		case 'd':
			if canPerformRight() {
				mergeRight()
				return
			} else {
				fmt.Println("无法执行该操作")
			}
		default:
			fmt.Println("请输入有效操作")
		}
		if !check(arr) {
			fmt.Println("很遗憾你输了,游戏结束!")
			break
		}
	}

}
func check(arr [4][4]int) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if arr[i][j] == 0 {
				return true
			}
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if j < 3 && arr[i][j] == arr[i][j+1] {
				return true
			}
			if i < 3 && arr[i][j] == arr[i+1][j] {
				return true
			}
		}
	}
	for i := 0; i < 3; i++ {
		if arr[i][3] == arr[i+1][3] {
			return true
		}
	}
	for j := 0; j < 3; j++ {
		if arr[3][j] == arr[3][j+1] {
			return true
		}

	}
	return false
}
func canPerformUp() bool {
	for j := 0; j < 4; j++ {
		for i := 1; i < 4; i++ {
			if arr[i][j] != 0 {
				if arr[i-1][j] == 0 || arr[i-1][j] == arr[i][j] {
					return true
				}
			} else if arr[i][j] == 0 {
				return true
			}

		}
	}
	return false
}
func canPerformDown() bool {
	for j := 0; j < 4; j++ {
		for i := 0; i < 3; i++ {
			if arr[i][j] != 0 {
				if arr[i+1][j] == 0 || arr[i+1][j] == arr[i][j] {
					return true
				}
			} else if arr[i][j] == 0 {
				return true
			}
		}
	}
	return false
}
func canPerformRight() bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if arr[i][j] != 0 {
				if arr[i][j+1] == 0 || arr[i][j+1] == arr[i][j] {
					return true
				}
			} else if arr[i][j] == 0 {
				return true
			}
		}
	}
	return false
}
func canPerformLeft() bool {
	for i := 0; i < 4; i++ {
		for j := 1; j < 4; j++ {
			if arr[i][j] != 0 {
				if arr[i][j-1] == 0 || arr[i][j-1] == arr[i][j] {
					return true
				}
			} else if arr[i][j] == 0 {
				return true
			}
		}
	}
	return false
}
