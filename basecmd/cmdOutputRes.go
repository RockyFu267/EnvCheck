package basecmd

import (
	"bufio"
	_ "fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

//CmdAndChangeDirToResAllInOne 输入命令 支持管道符， 取消line换行符
func CmdAndChangeDirToResAllInOne(dir string, commandName string) ([]string, error) {
	res := []string{}
	cmd := exec.Command("bash", "-c", commandName)
	//log.Println("CmdAndChangeDirToResAllInOne", dir, cmd.Args)
	//StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
	stdout, err := cmd.StdoutPipe()
	if err != nil {

		log.Println("cmd.StdoutPipe: ", err)
		return res, err
	}
	cmd.Stderr = os.Stderr
	cmd.Dir = dir
	err = cmd.Start()
	if err != nil {

		return res, err
	}
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	for {

		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {

			break
		}
		line = strings.Replace(line, "\n", "", -1)

		// // 去除空格的写法
		// str = strings.Replace(str, " ", "", -1)
		res = append(res, line)
	}
	err = cmd.Wait()
	return res, err
}
