package simple

import (
	"fmt"
	"testing"
)

func TestMarkdown(t *testing.T) {
	mr := NewMd(MdWithTOC()).Run(`
# 一级目录
## 本次更新内容
## 功能预览
### 三级目录
## 相关链接
## 目录3
`)

	fmt.Println(mr.TocHtml)
	fmt.Println("---------------------------------------------------------------")
	fmt.Println(mr.ContentHtml)
}
