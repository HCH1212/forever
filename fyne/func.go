package fyne

import (
	"forever/dao"
	"forever/model"
	"fyne.io/fyne/v2"
)

func (f *Forever) addBtnFunc() func() {
	return func() {
		// 如果已有编辑窗口打开，先关闭它
		if f.EditWindow != nil {
			f.EditWindow.Hide()
			f.EditWindow = nil
		}

		f.CurrentID = -1
		f.openEditWindow()
	}
}

func (f *Forever) openEditWindow() {
	f.EditWindow = f.App.NewWindow("备忘录")
	f.EditWindow.Resize(fyne.NewSize(500, 400))
	f.EditWindow.SetCloseIntercept(func() {
		f.closeEditWindow()
	})
	f.makeEditUI()

	if f.CurrentID != -1 {
		f.TitleEntry.SetText(f.AllData[f.CurrentID].Title)
		f.ContentEntry.SetText(f.AllData[f.CurrentID].Content)
	}

	f.EditWindow.Show()
}

func (f *Forever) closeEditWindow() {
	if f.EditWindow != nil {
		f.EditWindow.Hide()
		f.EditWindow = nil
	}
	f.CurrentID = -1
}

func (f *Forever) showMemoDetails() {
	if f.CurrentID < 0 || f.CurrentID >= len(f.AllData) {
		return
	}
	f.openEditWindow()
}

func (f *Forever) saveMemo() {
	title := f.TitleEntry.Text
	content := f.ContentEntry.Text
	if title == "" {
		return
	}

	if f.CurrentID == -1 {
		// 新建
		err := model.CreateData(dao.DB, title, content)
		if err != nil {
			return
		}
	} else {
		// 修改
		err := model.UpdateDataByID(dao.DB, f.AllData[f.CurrentID].ID, title, content)
		if err != nil {
			return
		}
	}

	f.refreshList()
	f.closeEditWindow()
}

func (f *Forever) deleteMemo() {
	if f.CurrentID == -1 {
		return
	}
	err := model.DeleteDataByID(dao.DB, f.AllData[f.CurrentID].ID)
	if err != nil {
		return
	}
	f.refreshList()
	f.closeEditWindow()
}

func (f *Forever) refreshList() {
	datas, err := model.ListAllData(dao.DB)
	if err != nil {
		return
	}
	f.AllData = datas
	if f.ListWidget != nil {
		f.ListWidget.Refresh()
	}
}
