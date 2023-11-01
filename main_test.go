package main

import (
	"kgithub.com/therecipe/qt/core"
	"kgithub.com/therecipe/qt/gui"
	"kgithub.com/therecipe/qt/widgets"
	"os"
)

func main() {
	// 初始化Qt应用程序
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// 创建主窗口
	mainWindow := widgets.NewQMainWindow(nil, 0)
	mainWindow.SetWindowTitle("Hello Qt")

	// 创建标签
	label := widgets.NewQLabel2("Hello, Qt!", nil, 0)
	label.SetAlignment(core.Qt__AlignCenter)

	// 设置标签字体样式
	font := gui.NewQFont2("Arial", 16, 1, false)
	label.SetFont(font)

	// 添加标签到主窗口的中心
	mainWindow.SetCentralWidget(label)

	// 显示主窗口
	mainWindow.Show()

	// 运行Qt应用程序
	app.Exec()
}
