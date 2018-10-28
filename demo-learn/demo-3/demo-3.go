package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

//定义结构体,将进程的错误和输出赋值给结构体
type result struct {
	err    error
	output []byte
}

//超时任务主动杀死，实现
func main() {
	//执行一个CMD，让它在一个携程中去执行2秒输出，nihao,到达一秒的时候杀死它

	var (
		//定义安全上下文
		ctx        context.Context
		cancelFunc context.CancelFunc
		cmd        *exec.Cmd
		resultChan chan *result
		res        *result
	)

	// 创建一个结果队列
	resultChan = make(chan *result, 1000)
	// context -> chan byte
	// cancelFunc -> close(chan byte)
	ctx, cancelFunc = context.WithCancel(context.TODO())

	// select {case <- ctx.Done()监听程序，感知状态} -< kill Pid,杀死进程
	go func() {
		var (
			output []byte
			err    error
		)
		cmd = exec.CommandContext(ctx, "C:\\cygwin64\\bin\\bash.exe", "-c", "sleep 1;echo nihao;")
		//子协程中执行，将输出传递给main主进程
		output, err = cmd.CombinedOutput()
		//将任务输出结果给main协程，结构体
		resultChan <- &result{
			err:    err,
			output: output,
		}
	}()

	//传入间隔时间
	time.Sleep(1 * time.Second)

	//取消上下文
	cancelFunc()

	//在main协程里，等待子协程的退出，并且打印任务执行的结果
	res = <-resultChan

	//打印输出的所有结果并转换字符串
	fmt.Println(res.err, string(res.output))

}
