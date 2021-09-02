package style

//编码格式
const (
	UrlFormat   = "[%s](%s)"
	ImageFormat = "![screenshot](%s)"
	Separator   = "\n >"
)

//字体标识
const (
	h1 markType = "h1"
	h2 markType = "h2"
	h3 markType = "h3"
	h4 markType = "h4"
	h5 markType = "h5"
	h6 markType = "h6"
)

//字体规则
var typefacesMap = map[markType]string{
	h1: "# %s",
	h2: "## %s",
	h3: "### %s",
	h4: "#### %s",
	h5: "##### %s",
	h6: "###### %s",
}

//颜色标识
const (
	red   markType = "red"
	blue  markType = "blue"
	green markType = "green"
	gold  markType = "gold"
)

//颜色规则
var colorsMap = map[markType]string{
	red:   "<font color=#ff0000>%s</font>",
	blue:  "<font color=#1E90FF>%s</font>",
	green: "<font color=#00CD66>%s</font>",
	gold:  "<font color=#FFD700>%s</font>",
}
