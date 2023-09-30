package main

import (
	"fmt"

	"github.com/ppz-pro/ppz.go/ppz_sli"
)

func main() {
	// 一个 slice
	nums := []int{1, 2, 3, 4, 5, 6}

	// 装上 ppz 武器
	nums_arm := ppz_sli.Arm(nums)

	// 取出偶数
	even_nums_arm := nums_arm.Filter(func(num int) bool {
		return num%2 == 0
	})
	fmt.Println(even_nums_arm)

	// 逆序
	r_nums_arm := even_nums_arm.Reverse()
	fmt.Println(r_nums_arm)

	// 全小于 8
	lt_eight := r_nums_arm.Every(func(num int) bool {
		return num < 8
	})
	fmt.Println(lt_eight) // true

	// 有一个大于 6 的
	gt_six := r_nums_arm.Some(func(num int) bool {
		return num > 6
	})
	fmt.Println(gt_six) // false

	// 第一个小于 5 的
	first, index, found := r_nums_arm.Find(func(num int) bool {
		return num < 5
	})
	if found {
		fmt.Printf("已找到小于 5 的元素：%d，下标为 %d\n", first, index)
	} else {
		fmt.Println("没有小于 5 的元素")
	}

	// 平方
	squares := ppz_sli.Map(r_nums_arm, func(num int) string {
		return fmt.Sprintf("%d 的平方是 %d", num, num*num)
	})
	fmt.Println(squares)
}
