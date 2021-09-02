package gomarkdown_test

import (
	"github.com/ckall/gomarkdown"
	"github.com/ckall/gomarkdown/style"
	"testing"
)

func TestFormat_AddImage(t *testing.T) {
	format := gomarkdown.NewConText()
	format.AddImage("https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png")
	context := format.GetContext()
	t.Log(context)
}

func TestFormat_AddKeyValue(t *testing.T) {
	format := gomarkdown.NewConText()
	format.AddKeyValue("【标题】:", "标题11")
	context := format.GetContext()
	t.Log(context)
}

func TestFormat_AddText(t *testing.T) {
	format := gomarkdown.NewConText()
	format.AddText("https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png")
	context := format.GetContext()
	t.Log(context)
}

func TestFormat_GetContext(t *testing.T) {
	context := gomarkdown.NewConText()
	context.AddText("# 杭州天气")
	context.AddText(style.AddH3("9度，西北风1级，空气良89，相对温度73%"))
	context.AddKeyValue("#### 【测试】:", map[string]interface{}{"测": "试"})
	context.AddKeyValue(style.AddH4("【测试】:"), map[string]interface{}{"测": "试"})
	context.AddImage("https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png")
	context.AddImage("https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png")
	context.AddText(
		style.AddH4("10点20分发布: %s 和 %s "),
		style.AddUrl(style.AddRed("天气"), "http://www.thinkpage.cn/"),
		style.AddUrl(style.AddBlue("天气11"), "http://www.baidu.com/"),
	)
	context.AddText(style.AddH6("杭 %s 和 %s"), style.AddRed("hiehie"), style.AddBlue("hiehie"))
	context.AddText("杭 %s 和 %s", style.AddGreen("hiehie"), style.AddGold("hiehie"))
	t.Log(context.GetContext())
}
