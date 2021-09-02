package style_test

import (
	"github.com/ckall/gomarkdown/style"
	"testing"
)

//红色字体
func TestAddRed(t *testing.T) {
	color := style.AddRed("ceshi")
	t.Log(color)
}

//蓝色
func TestAddBlue(t *testing.T) {
	color := style.AddBlue("ceshi")
	t.Log(color)
}

//绿色
func TestAddGreen(t *testing.T) {
	color := style.AddGreen("ceshi")
	t.Log(color)
}

//黄金
func TestAddGold(t *testing.T) {
	color := style.AddGold("ceshi")
	t.Log(color)
}
