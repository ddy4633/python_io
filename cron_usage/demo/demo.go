package main

import (
	"fmt"

	"github.com/gorhill/cronexpr"
)

func main() {
	var (
		expr *cronexpr.Expression
		err  error
	)

	//分钟(0-59),小时(0-23),天(1-31),月份(1-12),星期(0-6)

	//每分钟执行一次
	if expr, err = cronexpr.Parse("* * * * *"); err != nil {
		fmt.Println(err)
		return
	}
	expr = expr
}
