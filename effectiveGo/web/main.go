/*
此程序为一种数据格式提供了更好的的接口：
给定一小段文本，它将调用图表服务器来生成二维码（QR 码），这是一种编码文本的点格矩阵。
该图像可被你的手机摄像头捕获，并解释为一个字符串，比如 URL，
这样就免去了你在狭小的手机键盘上键入 URL 的麻烦。
*/
package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")
var templ = template.Must(template.New("qr").Parse(templateStr))

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func QR(w http.ResponseWriter, req *http.Request) {
	templ.Execute(w, req.FormValue("s"))
}

// 在模板文本（templateStr）中，双大括号界定的文本表示模板的动作。
//从 {{if .}} 到 {{end}} 的代码段仅在当前数据项（这里是点 .）的值非空时才会执行。
//也就是说，当字符串为空时，此部分模板段会被忽略。
const templateStr = `
<html>
<head>
	<title>QR Link Generator</title>
</head>
<body>
	{{if .}}
		<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
		<br>
		{{.}}
		<br>
		<br>
	{{end}}
	<form action="/" name=f method="GET">
		<input maxLength=1024 size=70 name=s value="" title="Text to QR Encode">
		<input type=submit value="Show QR" name=qr>
	</form>
</body>
</html>
`
