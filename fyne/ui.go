package fyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"net/url"
)

func (f *Forever) makeMainUI() {
	addBtn := widget.NewButtonWithIcon("", theme.ContentAddIcon(), f.addBtnFunc())
	f.refreshList()

	f.ListWidget = widget.NewList(
		func() int {
			return len(f.AllData)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(
				widget.NewLabel(""),
				widget.NewLabel(""),
			)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			item := f.AllData[i]
			hBox := o.(*fyne.Container)

			// 第一个元素是标题
			titleLabel := hBox.Objects[0].(*widget.Label)
			titleLabel.SetText(item.Title)

			// 第二个元素是时间
			timeLabel := hBox.Objects[1].(*widget.Label)
			timeLabel.SetText(item.UpdatedAt.Format("2006-01-02 15:04"))
		})

	f.ListWidget.OnSelected = func(id widget.ListItemID) {
		if id >= 0 && id < len(f.AllData) {
			if f.EditWindow != nil {
				f.EditWindow.Hide()
				f.EditWindow = nil
			}
			f.CurrentID = id
			f.showMemoDetails()
		}
		f.ListWidget.UnselectAll()
	}

	link := widget.NewHyperlink(
		"关于我",
		&url.URL{
			Scheme: "https",
			Host:   "github.com",
			Path:   "/HCH1212",
		},
	)
	link.Alignment = fyne.TextAlignCenter

	f.MainWindow.SetContent(container.NewBorder(addBtn, link, nil, nil, f.ListWidget))
}

func (f *Forever) makeEditUI() {
	f.TitleEntry = widget.NewEntry()
	f.TitleEntry.SetPlaceHolder("输入标题...")

	f.ContentEntry = widget.NewMultiLineEntry()
	f.ContentEntry.SetPlaceHolder("输入内容...")

	saveBtn := widget.NewButton("保存", func() {
		f.saveMemo()
		f.closeEditWindow()
	})

	cancelBtn := widget.NewButton("取消", func() {
		f.closeEditWindow()
	})

	deleteBtn := widget.NewButton("删除", func() {
		f.deleteMemo()
		f.closeEditWindow()
	})

	// 根据是否是新建备忘录来设置删除按钮状态
	if f.CurrentID == -1 {
		deleteBtn.Disable()
	} else {
		deleteBtn.Enable()
	}

	topBtn := container.NewGridWithColumns(3, saveBtn, cancelBtn, deleteBtn)

	f.EditWindow.SetContent(container.NewBorder(
		f.TitleEntry,
		topBtn,
		nil,
		nil,
		f.ContentEntry,
	))
}
