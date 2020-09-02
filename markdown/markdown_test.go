package markdown

import (
	"fmt"
	"testing"
)

func TestMarkdown(t *testing.T) {
	htmlStr, summary := New(EnableTOC()).Run(`
## 本次更新内容
## 功能预览
### 三级目录
## 相关链接
## 目录3
## 目录2
`)

	fmt.Println(htmlStr)
	fmt.Println("---------------------------------------------------------------")
	fmt.Println(summary)
}
