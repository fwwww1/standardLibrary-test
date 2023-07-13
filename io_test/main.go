package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 读文件
func openTest(fileName string) {
	//只读方式打开
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Print("open file failed")
		return
	}
	//关闭文件
	file.Close()
}

// 写文件
func writeTest(fileName string) {
	//新建文件
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Print("create file failed")
		return
	}
	defer file.Close()
	for i := 0; i < 5; i++ {
		file.WriteString("abc\n")  //写入string信息到文件
		file.Write([]byte("de\n")) //写入byte类型的信息到文件
	}
}

// 读文件
func readTest(fileName string) {
	//打开文件
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Print("failed open file")
		return
	}
	//定义接收读取文件的字节数组
	var buf [128]byte
	var content []byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			fmt.Print("completed")
			//读取结束
			break
		}
		if err != nil {
			fmt.Print("failed")
			return
		}
		content = append(content, buf[:n]...) //加上'...'就可以追加一个数组， 否则只能追加一个元素
	}
	fmt.Println(content)
}

// 拷贝文件
func copyTest(fileName string, newFileName string) {
	//打开源文件
	file, err1 := os.Open(fileName)
	if err1 != nil {
		fmt.Println("failed open")
		return
	}
	//创建新文件
	newFile, err2 := os.Create(newFileName)
	if err2 != nil {
		fmt.Println("failed create")
		return
	}
	//缓冲读取
	buf := make([]byte, 1024)
	//读取源文件
	for {
		n, err3 := file.Read(buf)
		if err3 == io.EOF {
			break //读取结束
		}
		if err3 != nil {
			fmt.Println("failed read")
		}
		//写到新文件中
		newFile.Write(buf[:n])
	}
	defer file.Close()
	defer newFile.Close()

}

// bufio读数据
func wr() {
	//以只写方式打开
	file, err := os.OpenFile("file_test.txt", os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	//获取writer对象
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString("hello\n")
	}
	//刷新缓冲区，强制写出
	writer.Flush()
}
func re() {
	wr()
	//以只读的方式打开文件
	file, err := os.Open("file_test.txt")
	if err != nil {
		return
	}
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		fmt.Println(string(line))
	}
}
func main() {
	//fileName := "file_test.txt"
	//newFileName := "new_file_test.txt"
	//openTest(fileName)
	//writeTest(fileName)
	//readTest(fileName)
	//copyTest(fileName, newFileName)
	//re()
}
