package main
import (
	"fmt"
	"os/exec"
)

func main() {
	//头部定义声明变量
	var (
		cmd *exec.Cmd
		output []byte
		err error
	)
	//生产cmd命令
	cmd = exec.Command("C:\\cygwin64\\bin\\bash.exe","-c","ls -l")
	//如果结果不为空则捕获进程的输出，调用CombinedOutput
	if output,err = cmd.CombinedOutput();err != nil {
		fmt.Println(err)
		return
	}
	//打印子进程的输出
	fmt.Println(string(output))
}