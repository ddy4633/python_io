package main

import "os/exec"

//超时任务主动杀死，实现
func main() {
	//执行一个CMD，让它在一个携程中去执行2秒输出，nihao,到达一秒的时候杀死它

	cts,cn
	go func() {
		exec.CommandContext(/*Context*/,"C:\\cygwin64\\bin\\bash.exe","-c","sleep 5;echo nihao;")
	}()

}
