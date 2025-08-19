package main

import (
	"github.com/lxn/walk"
)

func test() {
	mw, _ := walk.NewMainWindow()
	mw.SetTitle("위아래 분할 레이아웃")
	mw.SetSize(walk.Size{Width: 800, Height: 600})
	mw.SetLayout(walk.NewVBoxLayout())

	// 위쪽 고정 영역
	top, _ := walk.NewComposite(mw)
	top.SetBounds(walk.Rectangle{X: 0, Y: 0, Width: 800, Height: 100}) // 고정 높이 100px

	label, _ := walk.NewLabel(top)
	label.SetText("이 영역은 고정 크기입니다.")
	label.SetBounds(walk.Rectangle{X: 10, Y: 10, Width: 200, Height: 20})

	// 아래쪽 확장 영역
	bottom, _ := walk.NewComposite(mw)
	bottom.SetBounds(walk.Rectangle{X: 0, Y: 100, Width: 800, Height: 500}) // 남은 영역

	textEdit, _ := walk.NewTextEdit(bottom)
	textEdit.SetBounds(walk.Rectangle{X: 10, Y: 10, Width: 780, Height: 480})
	textEdit.SetText("이 영역은 창 크기에 따라 조정 가능합니다.")

	mw.Show()
	mw.Run()
}
