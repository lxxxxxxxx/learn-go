package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//显示定义变量，一般for循环
	var s, seq string
	for i := 0; i < len(os.Args); i++ {
		s += seq + os.Args[i]
		seq = " "
	}
	fmt.Println("第一种实现：" + s)

	s1, seq1 := "", ""
	for _, arg := range os.Args[0:] {
		s1 += seq1 + arg
		seq1 = " "
	}
	fmt.Println("第二种实现：" + s1)

	fmt.Println("第三种实现：" + strings.Join(os.Args[0:], " "))

	fmt.Print("第四种实现：")
	fmt.Print(os.Args)

	fmt.Println("命令名：" + os.Args[0])

	fmt.Println("打印索引和值：")
	for i, arg := range os.Args[:] {
		fmt.Printf("%d  %s \n,type : %T %T", i, arg, i, arg)
	}

}
