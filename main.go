package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mozillazg/go-pinyin"
)

func main() {
	// 定义命令行参数
	outputOnly := flag.Bool("o", false, "只输出第一列内容")
	flag.Parse()

	// 创建拼音转换器
	p := pinyin.NewArgs()
	p.Style = pinyin.Normal // 使用普通风格（不带声调）
	p.Separator = ""        // 不使用分隔符

	// 创建带缓冲的读取器
	scanner := bufio.NewScanner(os.Stdin)

	// 逐行处理输入
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// 如果启用了 -o 选项，检查是否已经是三列格式
		if *outputOnly {
			parts := strings.Split(line, "|")
			if len(parts) >= 1 {
				fmt.Println(parts[0])
				continue
			}
		}

		// 获取全拼
		fullPinyin := pinyin.Pinyin(line, p)
		fullPinyinStr := ""
		for _, py := range fullPinyin {
			fullPinyinStr += py[0]
		}

		// 获取首字母
		initials := ""
		for _, py := range fullPinyin {
			if len(py[0]) > 0 {
				initials += strings.ToLower(py[0][:1])
			}
		}

		// 输出格式：原文|首字母|全拼
		fmt.Printf("%s|%s|%s\n", line, initials, fullPinyinStr)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
