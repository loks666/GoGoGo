package main

import (
	"bytes"
	"fmt"
	"strings"
)

//定义Buffer类型
var bt bytes.Buffer
var build strings.Builder

func main() {
	build.WriteString("1111111")
	println(build.String())
	//需要先导入bytes包
	s1 := "字符串"
	s2 := "拼接"
	bt.WriteString(s1)
	bt.WriteString(s2)
	//获得拼接后的字符串
	s3 := bt.String()

	println(s3)
	// %d 表示整型数字，%s 表示字符串
	var stockade = 123
	var enodata = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var targetUrl = fmt.Sprintf(url, stockade, enodata)
	fmt.Println(targetUrl)

	if Sub(4, 3) > 0 {
		println("大于0测试一下")
	}
	//println("结果值：" + strconv.Itoa())
}

func Sub(x, y int) int {
	return x - y
}
