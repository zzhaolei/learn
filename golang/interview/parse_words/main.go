/*
请完成以下任务：
 1. 请勿将外部IDE中编写好的代码粘贴到此Pad中，需在此Pad中完成代码编写和调试。
 2. 面试结束后，请点击右下角的“End Interview”按钮。
 3. 字符输入规则如下：
    a. 每行表示一条记录，字段之间以逗号（,）分隔。
    b. 如果字段内容包含逗号（,），则需使用双引号（"）包裹。
    c. 如果字段内容包含双引号（"），则需使用两个双引号（""）进行转义并用双引号包裹。
 4. 编写解析程序，将解析后的内容按行输出，字段之间以制表符（\t）分隔。
 5. 以下是一个示例：
    1）输入：Linda,47,"旅游,""攀岩",New Job
    2）输出：Linda 47 旅游,"攀岩 New Job

请根据以上要求编写解析程序，并将解析后的内容按照规则输出。
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	rows := `2,Tina,37,"足球,""篮球",Old Job
3,Alice Job,66,"""看电影"",旅游","上海,上海市"
4,John,44,"洗衣机101,""","LA""CITY"""
5,"Jane,li",55,Hiking,Canada`
	execute(rows)
}

func execute(rows string) {
	for _, str := range strings.Split(rows, "\n") {
		inQuote := false
		runes := []rune(str)
		var builder strings.Builder
		for i := 0; i < len(runes); i++ {
			if runes[i] == '"' {
				if !inQuote {
					inQuote = true
				} else if (i+1) < len(runes) && runes[i+1] == '"' {
					builder.WriteRune('"')
					i++
				} else {
					inQuote = false
				}
				continue
			}

			if inQuote || runes[i] != ',' {
				builder.WriteRune(runes[i])
			} else {
				builder.WriteRune('\t')
			}
		}
		fmt.Println(builder.String())
	}
}
