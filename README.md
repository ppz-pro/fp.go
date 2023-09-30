# PPZ.go
一堆 Go 语言的小工具。

## 安装
``` bash
go get -u github.com/ppz-pro/ppz.go
```

## 举个例子
ppz_sli 包提供了一些操作 slice 的工具。

比如，我们有一个 slice：
``` go
nums := []int{1, 2, 3, 4, 5, 6}
```

装上 ppz 武器：
``` go
nums_arm := ppz_sli.Arm(nums)
```

取出偶数：
``` go
even_nums_arm := nums_arm.Filter(func(num int) bool {
  return num%2 == 0
})
fmt.Println(even_nums_arm) // [2 4 6]
```

逆序：
``` go
r_nums_arm := even_nums_arm.Reverse()
fmt.Println(r_nums_arm) // [6 4 2]
```

全小于 8：
``` go
lt_eight := r_nums_arm.Every(func(num int) bool {
  return num < 8
})
fmt.Println(lt_eight) // true
```

有一个大于 6 的（当然是 false）：
``` go
gt_six := r_nums_arm.Some(func(num int) bool {
  return num > 6
})
fmt.Println(gt_six) // false
```

第一个小于 5 的：
``` go
first, index, found := r_nums_arm.Find(func(num int) bool {
  return num < 5
})
if found {
  fmt.Printf("已找到小于 5 的元素：%d，下标为 %d\n", first, index)
} else {
  fmt.Println("没有小于 5 的元素")
}
```

平方：
``` go
squares := ppz_sli.Map(r_nums_arm, func(num int) string {
  return fmt.Sprintf("%d 的平方是 %d", num, num*num)
})
fmt.Println(squares)
/*
  [
    "6 的平方是 36",
    "4 的平方是 16",
    "2 的平方是 4",
  ]
*/
```
