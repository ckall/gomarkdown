package style

import "fmt"

/**
 * @auth: kuncheng
 * @Date: 2021/9/2
 */

//字体型号H1
func AddH1(text string) string {
	return fmt.Sprintf(typefacesMap[h1], text)
}

//字体型号H2
func AddH2(text string) string {
	return fmt.Sprintf(typefacesMap[h2], text)
}

//字体型号H3
func AddH3(text string) string {
	return fmt.Sprintf(typefacesMap[h3], text)
}

//字体型号H4
func AddH4(text string) string {
	return fmt.Sprintf(typefacesMap[h4], text)
}

//字体型号H5
func AddH5(text string) string {
	return fmt.Sprintf(typefacesMap[h5], text)
}

//字体型号H6
func AddH6(text string) string {
	return fmt.Sprintf(typefacesMap[h6], text)
}

//添加URL地址
func AddUrl(title string, url string) string {
	return fmt.Sprintf(UrlFormat, title, url)
}
