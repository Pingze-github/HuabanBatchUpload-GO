package main

import (
	"fmt"
	"flag"
	"os"
	"time"
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

func goTasks(task func(param Param), paramArr ParamArr, limit int) {
	// goroutine 协程
	// channel 信道
	// 死锁 deadlock

	fcount := 0
	limit = 3
	channel := make(chan int, limit)
	quit := make(chan int)
	taskWrapper := func(param Param) {
		channel <- 1
		task(param)
		<-channel
		fcount ++
		if fcount == paramArr.Len() {
			<-quit
		}
	}
	var i int
	for i = 0; i < paramArr.Len(); i++ {
		go taskWrapper(paramArr.Get(i))
	}
	quit <- 1
}

func foo(param Param) {
	fmt.Println(param.url)
	time.Sleep(1000000000) //1s
}

type ParamArr []Param
func (arr ParamArr) Len() int {
	return len(arr)
}
func (arr ParamArr) Get(index int) Param {
	return arr[index]
}

type Param struct {
	url string
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
	}
	fmt.Println(acct, pass, board)

	// TODO 建立一个有效的并发控制工具
	// 1、可以定制要执行的任务
	// 2、可以定制要执行的数据
	// 3、可以定制并发数
	// 4、可以捕获到单个任务完成，并获得返回值
	// 5、可以在全部完成时获得返回值(选择性)
	//data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}

	var paramArr ParamArr
	paramArr = append(paramArr, Param{url: "http://www.baidu.com"})
	paramArr = append(paramArr, Param{url: "http://www.baidu.com"})
	paramArr = append(paramArr, Param{url: "http://www.baidu.com"})
	paramArr = append(paramArr, Param{url: "http://www.baidu.com"})
	paramArr = append(paramArr, Param{url: "http://www.baidu.com"})
	paramArr = append(paramArr, Param{url: "http://www.baidu.com"})

	// FIXME 使用此函数进行并发控制，必须让任务函数参数、数组类型都为interface{}。这使类型检查无效，不该这样
	goTasks(foo, paramArr, 3)
}
