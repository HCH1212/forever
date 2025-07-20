package fyne

import (
	"forever/model"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

type Forever struct {
	App          fyne.App      // 应用实例
	AllData      []model.Data  // 所有备忘录
	MainWindow   fyne.Window   // 主窗口
	EditWindow   fyne.Window   // 编辑窗口
	TitleEntry   *widget.Entry // 标题输入框
	ContentEntry *widget.Entry // 内容输入框
	CurrentID    int           // AllData 索引
	ListWidget   *widget.List  // 列表组件
}

var forever Forever

func Init() {
	forever.CurrentID = -1
	forever.App = app.NewWithID("forever.app")
	forever.MainWindow = forever.App.NewWindow("forever备忘录")
	forever.MainWindow.SetMaster()
	forever.MainWindow.Resize(fyne.NewSize(600, 800))
	forever.MainWindow.CenterOnScreen()

	forever.makeMainUI()
	forever.MainWindow.ShowAndRun()
}
