package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容
func FprintTest() {
	//向标准输出写入内容
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("failed open:err", err)
		return
	}
	name := "zhangsan"
	//向打开的文件句柄中写入内容
	fmt.Fprintf(file, "往文件中写入信息：%s", name)
}

// Sprint 系列函数会把传入数据生成并返回字符串
func SprintTest() {
	s1 := fmt.Sprint("zhangsan")
	name := "zhangsan"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("zhangsan")
	fmt.Println(s1, s2, s3)
}

// Errorf函数根据fomat参数生成格式化字符串并返回一个包含该字符串的错误
func ErrorfTest() {
	//通常使用这种方式来自定义错误类型
	//error := "错误"
	err := fmt.Errorf("这是一个错误")
	fmt.Println(err)
}

// Scan从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符。
// 本函数返回成功扫描的数据个数和遇到的任何错误。如果读取的数据个数比提供的参数少，会返回一个错误报告原因。
func ScanTest() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scan(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t\n", name, age, married)

}

// fmt.Scanf不同于fmt.Scan简单的以空格作为输入数据的分隔符，fmt.Scanf为输入数据指定了具体的输入内容格式，只有按照格式输入数据才会被扫描并存入对应变量
func ScanfTest() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scanf("1:%s,2:%d,3:%t", &name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t\n", name, age, married)

}

// Scanln类似Scan，它在遇到换行时才停止扫描。最后一个数据后面必须有换行或者到达结束位置。
// 本函数返回成功扫描的数据个数和遇到的任何错误。
func ScanlnTest() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scanln(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t\n", name, age, married)
}
func bufioNewReaderTest() {
	reader := bufio.NewReader(os.Stdin) //从标准输入生成一个读对象
	fmt.Println("请输入内容：")
	text, _ := reader.ReadString('\n') //读到换行
	text = strings.TrimSpace(text)     //strings.TrimSpace()函数用来去除一个字符串的首尾空格，如果需要去除所有空格，可以使用 strings.ReplaceAll() 函数替换所有空格为空字符串
	fmt.Printf("%#v", text)
}

// Fscan不是从标准输入中读取数据而是从io.Reader中读取数据
// Sscan不是从标准输入中读取数据而是从指定字符串中读取数据
func main() {
	//FprintTest()
	//SprintTest()
	//ErrorfTest()
	//ScanTest()
	//ScanfTest()
	//ScanlnTest()
	bufioNewReaderTest()
}
