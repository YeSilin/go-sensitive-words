package main

import (
	"fmt"
	"github.com/YeSilin/go-sensitive-words/tools"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	// 读取数据
	data, err := ioutil.ReadFile("examples/original.txt")
	if err != nil {
		fmt.Println("ioutil.ReadFile err:", err)
		return
	}

	// 有的是用 or 分割，先全部转成换行
	str := strings.ReplaceAll(string(data), "|", "\n")

	// 有的用#号分割的，也全部转成换行
	str = strings.ReplaceAll(str, "#", "\n")

	// 统一换行符
	str = strings.ReplaceAll(str, "\r\n", "\n")

	// 再根据换行符转成字符串切片
	strSlice := strings.Split(str, "\n")

	// 去重
	strSlice = tools.RemoveRepeatedElement(strSlice)

	// 删除切片空元素
	strSlice = tools.DeleteSlice(strSlice)

	// 按拼音排序
	sort.Sort(tools.ByPinyin(strSlice))

	// 保存两份不同的数据
	tools.SaveFile(strSlice, "listLn.txt", "\n")
	tools.SaveFile(strSlice, "listOr.txt", "|")
	tools.SaveFile(strSlice, "listWell.txt", "#")
}
