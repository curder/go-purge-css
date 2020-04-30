package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

var (
	originFile string // 源文件
	parseLine  int    // 分析出的行
	distFile   string // 目标文件
)

// 正则表达式获取页面class和id的值
const classAndIDRegex = `(id|class)=["'](.*?)["']`

func main() {
	var (
		fileContent []byte   // 文件内容
		result      []string // 正则匹配后的内容
		err         error    // 错误
	)

	// 初始化命令行参数
	initArgs()

	// 获取页面文件内容
	if fileContent, err = readFile(originFile); err != nil {
		goto ERR
	}

	// 正则表达式分析出页面对应class和id的值并写入到对应的临时文件
	result = regexContent(fileContent)

	// 去重
	result = removeRepByMap(result)

	// 输出排序
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	// 写入到目标文件
	if err = writeFile(distFile, result); err != nil {
		goto ERR
	}

	fmt.Printf("已经分析出需要的class或id，总共%d个，保存在%s里", parseLine, distFile)
	return

ERR:
	fmt.Println(err.Error())
}

// 初始化命令行参数
func initArgs() {
	// ./go-purge-css -origin ./code.html -dist ./dist.txt
	flag.StringVar(&originFile, "origin", "./code.html", "请输入要分析的文件路径")
	flag.StringVar(&distFile, "dist", "./dist.txt", "请输入分析后要保存的文件路径")
	flag.Parse()
}

// removeRepByMap: 通过map主键唯一的特性过滤重复元素
func removeRepByMap(slc []string) []string {
	result := []string{}
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		// fmt.Println(e)
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

// readFile 读文件
func readFile(fileName string) (result []byte, err error) {
	if result, err = ioutil.ReadFile(fileName); err != nil {
		return nil, err
	}
	return result, nil
}

// writeFile 写文件
func writeFile(fileName string, content []string) (err error) {
	var (
		newFile *os.File
		str     string
	)
	if newFile, err = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644); err != nil {
		fmt.Printf("failed to create new file err: %s", err.Error())
		return
	}
	defer newFile.Close()

	for _, str = range content {
		if str == "" { // 去除空格
			continue
		}
		if _, err = newFile.WriteString(str + "\n"); err != nil {
			return
		}
		parseLine++
	}

	return
}

// 正则匹配内容
func regexContent(fileContent []byte) (result []string) {
	var (
		re       *regexp.Regexp
		matches  [][][]byte
		match    [][]byte
		subMatch string
	)
	re = regexp.MustCompile(classAndIDRegex)
	matches = re.FindAllSubmatch(fileContent, -1)
	for _, match = range matches {
		// 按照空格分割class或id字符串
		for _, subMatch = range strings.Split(string(match[2]), " ") {
			result = append(result, subMatch)
		}
	}
	return
}
