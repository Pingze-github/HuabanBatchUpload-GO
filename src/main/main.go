package main

import (
	"fmt"
	"flag"
	"os"
)

var (
	help bool
	acct string
	pass string
	board string
)

func init() {
	flag.BoolVar(&help, "h", false, "帮助")
	flag.StringVar(&acct, "a", "", "账号")
	flag.StringVar(&pass, "p", "", "密码")
	flag.StringVar(&board, "b", "", "画板名")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr,`
使用示例:
    main.exe -a [账号] -p [密码] -b [画板名]
参数:
	-h 帮助
	-a 账号
	-p 密码
	-b 画板名
	`)
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
	}
	fmt.Println(acct, pass, board)
}
