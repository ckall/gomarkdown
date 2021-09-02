package gomarkdown

import (
	"encoding/json"
	"fmt"
	"github.com/ckall/gomarkdown/style"
	"strings"
)

type Format interface {
	//添加一行
	AddText(text string, color ...interface{})
	//添加图片
	AddImage(url ...string)
	//获取处理好的数据格式
	GetContext() string
	//以key：value的数据格式
	AddKeyValue(title string, value interface{})
}

/**
 * @auth: kuncheng
 * @Date: 2021/9/1
 */
type format struct {
	text        []string
	imageFormat string
	separator   string
}

func NewConText() Format {
	return &format{
		imageFormat: style.ImageFormat,
		separator:   style.Separator,
	}
}

//添加文字
func (f *format) AddText(text string, color ...interface{}) {
	f.text = append(f.text, fmt.Sprintf(text, color...))
}

//添加图片
func (f *format) AddImage(urls ...string) {
	for _, url := range urls {
		f.text = append(f.text, fmt.Sprintf(f.imageFormat, url))
	}
}

//一次性添加多个图片
func (f *format) AddImages(urls []string) {
	for i := 0; i < len(urls); i++ {
		f.AddImage(urls[i])
	}
}

//输出原生语句可以自己研究
func (f *format) GetContext() string {
	return strings.Join(f.text, f.separator)
}

//以 【这样的】:形式输出
func (f *format) AddKeyValue(title string, value interface{}) {
	var str string
	switch value.(type) {
	case string:
		str = value.(string)
	default:
		b, _ := json.Marshal(value)
		str = string(b)
	}
	f.text = append(f.text, title+str)
}
