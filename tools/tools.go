// 其他常用工具函数

package tools

import (
	"fmt"
	"os"
	"strings"
)

// 通过两重循环过滤重复元素
func RemoveRepeatedElement(arr []string) []string {
	// 存放结果，空初始化比make定义要快
	result := []string{}
	// 外层循环准备添加到结果的切片
	for i := 0; i < len(arr); i++ {
		// 初始定义该元素不存在，很奇怪，初始在循环里面比先声明在外面之后再赋值要快
		exist := false
		// 这里根据当前切片的长度进行循环，直接使用 len 比初始一个 count变量 记数要快
		for j := 0; j < len(result); j++ {
			// 如果遇到重复提前退出
			if result[j] == arr[i] {
				// 并且说明已存在
				exist = true
				break
			}
		}
		// 如果在 result切片都没有遍历到此元素
		if !exist {
			// 那么就追加到 result
			result = append(result, arr[i])
		}
	}
	return result
}

// 创建文件并写入数据
func CreateFile(path, data string) (err error) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}
	defer f.Close()

	// 第一个参数返回数据长度，第二个，错误信息
	_, err = f.WriteString(data)
	if err != nil {
		fmt.Println("f.WriteString err:", err)
		return
	}
	return
}

// 切片删除所有空元素
// 这里利用一个index，记录应该下一个有效元素应该在的位置，遍历所有元素，当遇到有效元素，index加一，否则不加，
// 最终index的位置就是所有有效元素的下一个位置。最后做一个截取就行了。这种方法会对原来的slice进行修改
func DeleteSlice(s []string) []string {
	j := 0
	for _, val := range s {
		if val != "" {
			// 该元素有效，覆盖写入
			s[j] = val
			j++
		}
	}
	return s[:j]
}

// 将切片按指定分割符，保存文件
func SaveFile(slice []string, filename, sep string) {
	// 按指定分割符连接成字符串
	str := strings.Join(slice, sep)

	// 创建文件并写入数据
	err := CreateFile(filename, str)
	if err != nil {
		fmt.Println("CreateFile err:", err)
		return
	}
}