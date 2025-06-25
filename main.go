package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/mozillazg/go-pinyin"
)

// 处理结果结构
type result struct {
	line string
}

// worker函数：处理输入行
func worker(id int, input <-chan string, output chan<- result, wg *sync.WaitGroup, outputOnly bool) {
	defer wg.Done()

	// 创建拼音转换器（每个worker一个实例）
	p := pinyin.NewArgs()
	p.Style = pinyin.Normal
	p.Separator = ""

	for line := range input {
		if line == "" {
			continue
		}

		// 如果启用了 -o 选项
		if outputOnly {
			if idx := strings.Index(line, "|"); idx != -1 {
				output <- result{line: line[:idx]}
			} else {
				output <- result{line: line}
			}
			continue
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
		output <- result{line: fmt.Sprintf("%s|%s|%s", line, initials, fullPinyinStr)}
	}
}

func main() {
	// 定义命令行参数
	outputOnly := flag.Bool("o", false, "只输出第一列内容")
	workerCount := flag.Int("w", runtime.NumCPU(), "worker协程数量")
	flag.Parse()

	// 创建channel
	inputChan := make(chan string, 100000)   // 输入channel，带缓冲
	outputChan := make(chan result, 1000000) // 输出channel，带缓冲

	// 创建WaitGroup来等待所有worker完成
	var wg sync.WaitGroup

	// 启动worker协程
	for i := 0; i < *workerCount; i++ {
		wg.Add(1)
		go worker(i, inputChan, outputChan, &wg, *outputOnly)
	}

	// 启动输出协程
	go func() {
		wg.Wait()
		close(outputChan)
	}()

	// 启动输入协程
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			inputChan <- scanner.Text()
		}
		close(inputChan)

		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			os.Exit(1)
		}
	}()

	// 主协程负责输出结果
	for result := range outputChan {
		fmt.Println(result.line)
	}
}
