// 编写一个小程序：
// 给定一个字符串数组
// [“I”,“am”,“stupid”,“and”,“weak”]
// 用 for 循环遍历该数组并修改为
// [“I”,“am”,“smart”,“and”,“strong”]
package main

import "fmt"

func main() {
	array := [5]string{"I", "am", "stupid", "and", "weak"}
	for key, value := range array {
		if value == "stupid" {
			array[key] = "smart"
		} else if value == "weak" {
			array[key] = "strong"
		}
	}
	fmt.Println(array)
}
