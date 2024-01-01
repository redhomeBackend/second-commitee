package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(100)
	fmt.Println(number)
	fmt.Println("欢迎来到猜数字游戏,在1到100之间,你只有五次机会哦!")
	var guess int
	var times int
	var choice string
	for {
		fmt.Println("请输入你的猜测")
		fmt.Scanln(&guess)
		if guess < 1 || guess > 100 {
			fmt.Println("请输入符合条件的数字")
			continue
		}
		times++
		if guess < number {
			fmt.Println("你猜测的数字太小了")
		} else if guess > number {
			fmt.Println("你猜测的数字太大了")
		} else {
			fmt.Printf("恭喜你猜对了,共用了%d次机会", times)
			break
		}
		if times == 5 {
			fmt.Println("你已经用了5次机会，游戏结束。")
			fmt.Println("如果你想继续游戏,请输入1,否则0")
			fmt.Println("请输入")
			fmt.Scanln(&choice)
			if choice == "1" {
				continue
			} else {
				break
			}
		}

	}
}
