package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

var base64EncodeChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func main() {
	defendChange := base64EncodeChars
	for {
		var ques, ntbase64EncodeChars string
		fmt.Printf("请输入自定义码表(按enter为default)：")
		reader0 := bufio.NewReader(os.Stdin)              // 标准输入输出
		ntbase64EncodeChars, _ = reader0.ReadString('\n') // 回车结束
		ntbase64EncodeChars = strings.TrimSpace(ntbase64EncodeChars)
		if len(ntbase64EncodeChars) != 0 {
			base64EncodeChars = ntbase64EncodeChars
		}
		fmt.Printf("请输入要编码的内容：")
		reader := bufio.NewReader(os.Stdin) // 标准输入输出
		ques, _ = reader.ReadString('\n')   // 回车结束
		ques = strings.TrimSpace(ques)      // 去除最后一个空格
		//fmt.Printf(ques)
		answer, err := Conversion(ques)
		if err != nil {
			return
		}
		fmt.Println("编码结果：", answer)
		readers := bufio.NewReader(os.Stdin)
		fmt.Print("按下回车键继续...")
		_, _ = readers.ReadString('\n')
		base64EncodeChars = defendChange
		fmt.Print("\033[H\033[2J")
		time.Sleep(250 * time.Millisecond)
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	return
}
