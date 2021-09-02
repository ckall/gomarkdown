package style

import "fmt"

type markType string

//红色字体
func AddRed(text string) string {
	return fmt.Sprintf(colorsMap[red], text)
}

//蓝色
func AddBlue(text string) string {
	return fmt.Sprintf(colorsMap[blue], text)
}

//绿色
func AddGreen(text string) string {
	return fmt.Sprintf(colorsMap[green], text)
}

//黄金
func AddGold(text string) string {
	return fmt.Sprintf(colorsMap[gold], text)
}
