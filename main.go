package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func isDoubleClickMode() bool {
	// 如果有命令行参数（除了程序名本身），认为是命令行模式
	if len(os.Args) > 1 {
		return false
	}

	// Windows 下双击运行时通常没有控制台
	// macOS/Linux 下双击 .app 或可执行文件时一般也没有附加参数
	// 最简单可靠的方式：看是否有 --command 参数
	return true
}

func main() {
	// 定义命令行参数
	cmdMode := flag.Bool("command", false, "以命令行模式运行，只打印 hello world")
	flag.Parse()

	// 如果明确带了 --command 参数 → 命令行模式
	if *cmdMode {
		fmt.Println("hello world")
		os.Exit(0)
	}

	// 没有参数，且不是明确命令行调用 → 认为是双击打开 GUI
	// （Windows 下双击 .exe、macOS 双击 .app、Linux 双击可执行文件通常走这条路）

	guiMode()
}

func guiMode() {
	// 根据系统设置合适的应用程序 ID（主要是 Windows 和 macOS）
	var appID string
	switch runtime.GOOS {
	case "windows":
		appID = "com.example.helloworld"
	case "darwin":
		appID = "com.example.helloworld"
	default:
		appID = "helloworld"
	}

	myApp := app.NewWithID(appID)
	myWindow := myApp.NewWindow("Hello World")

	// 状态文本
	label := widget.NewLabel("hello world")
	label.Alignment = fyne.TextAlignCenter

	// 关闭按钮
	closeBtn := widget.NewButton("关闭", func() {
		myWindow.Close()
		myApp.Quit()
	})

	// 确认按钮
	confirmBtn := widget.NewButton("确认", func() {
		label.SetText("软件运行中")
	})

	// 布局：垂直排列
	content := container.NewVBox(
		widget.NewLabel(""), // 顶部留空
		label,
		widget.NewLabel(""),
		container.NewHBox(
			closeBtn,
			confirmBtn,
		),
	)

	myWindow.SetContent(container.NewCenter(content))
	myWindow.Resize(fyne.NewSize(320, 180))
	myWindow.CenterOnScreen()

	myWindow.ShowAndRun()
}