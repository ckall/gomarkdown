package style_test

import (
	"github.com/ckall/gomarkdown/style"
	"testing"
)

/**
 * @auth: kuncheng
 * @Date: 2021/9/2
 */

func TestAddH1(t *testing.T) {
	typeface := style.AddH1("ceshi")
	t.Log(typeface)
}

func TestAddH2(t *testing.T) {
	typeface := style.AddH2("ceshi")
	t.Log(typeface)
}

func TestAddH3(t *testing.T) {
	typeface := style.AddH3("ceshi")
	t.Log(typeface)
}

func TestAddH4(t *testing.T) {
	typeface := style.AddH4("ceshi")
	t.Log(typeface)
}

func TestAddH5(t *testing.T) {
	typeface := style.AddH5("ceshi")
	t.Log(typeface)
}

func TestAddH6(t *testing.T) {
	typeface := style.AddH6("ceshi")
	t.Log(typeface)
}

func TestAddUrl(t *testing.T) {
	typeface := style.AddUrl("ceshi", "http://baidu.com")
	t.Log(typeface)
}
