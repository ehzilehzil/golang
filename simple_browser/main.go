package main

import (
	"log"
	"os"

	"github.com/jchv/go-webview2/pkg/edge"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
)

// copilot.microsoft.com 에 질문한 내용과 github.com/pirogom/walkmgr 코드를 참고함
func main() {
	declarative()
}

// declarative 방식으로 구현
func declarative() {
	var main_w *walk.MainWindow
	var top_iv *walk.ImageView
	var bottom_iv *walk.ImageView

	top_wv := edge.NewChromium()
	top_wv.SetPermission(edge.CoreWebView2PermissionKindClipboardRead, edge.CoreWebView2PermissionStateAllow)
	top_wv.DataPath = os.TempDir()
	bottom_wv := edge.NewChromium()
	bottom_wv.SetPermission(edge.CoreWebView2PermissionKindClipboardRead, edge.CoreWebView2PermissionStateAllow)
	bottom_wv.DataPath = os.TempDir()

	err := MainWindow{
		AssignTo: &main_w,
		Title:    "심플 브라우저",
		Size:     Size{Width: 800, Height: 600},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				MinSize: Size{Width: 800, Height: 100},
				MaxSize: Size{Width: 800, Height: 100},
				Layout:  VBox{},
				Children: []Widget{
					ImageView{
						AssignTo: &top_iv,
						Mode:     ImageViewModeStretch,
						OnSizeChanged: func() {
							top_wv.NotifyParentWindowPositionChanged()
							top_wv.Resize()
						},
					},
				},
			},
			Composite{
				Layout: VBox{},
				Children: []Widget{
					ImageView{
						AssignTo: &bottom_iv,
						Mode:     ImageViewModeStretch,
						OnSizeChanged: func() {
							bottom_wv.NotifyParentWindowPositionChanged()
							bottom_wv.Resize()
						},
					},
				},
			},
		},
	}.Create()
	if err != nil {
		log.Fatal("윈도우 생성 에러")
	}

	main_w.Synchronize(func() {
		top_wv.Embed(uintptr(top_iv.Handle()))
		top_wv.Navigate("https://google.com")
		top_wv.Resize()

		bottom_wv.Embed(uintptr(bottom_iv.Handle()))
		bottom_wv.Navigate("https://bing.com")
		bottom_wv.Resize()
	})

	win.ShowWindow(main_w.Handle(), win.SW_MAXIMIZE)
	main_w.Run()
}

// Imperative 방식으로 구현
func imperative() {
	mw, _ := walk.NewMainWindow()

	mw.SetTitle("심플 브라우저")
	mw.SetSize(walk.Size{Width: 800, Height: 600})
	mw.SetLayout(walk.NewVBoxLayout())

	top, _ := walk.NewCompositeWithStyle(mw, win.WS_VISIBLE)
	top.SetMinMaxSize(walk.Size{Width: 800, Height: 100}, walk.Size{Width: 800, Height: 100})
	top.SetLayout(walk.NewVBoxLayout())

	bottom, _ := walk.NewCompositeWithStyle(mw, win.WS_VISIBLE)
	bottom.SetLayout(walk.NewVBoxLayout())

	ivTop, _ := walk.NewImageView(top)
	ivTop.SetMode(walk.ImageViewModeStretch)

	ivBottom, _ := walk.NewImageView(bottom)
	ivBottom.SetMode(walk.ImageViewModeStretch)

	wvTop := edge.NewChromium()
	wvTop.SetPermission(edge.CoreWebView2PermissionKindClipboardRead, edge.CoreWebView2PermissionStateAllow)
	wvTop.DataPath = os.TempDir()

	wvBottom := edge.NewChromium()
	wvBottom.SetPermission(edge.CoreWebView2PermissionKindClipboardRead, edge.CoreWebView2PermissionStateAllow)
	wvBottom.DataPath = os.TempDir()

	ivTop.SizeChanged().Attach(func() {
		wvTop.NotifyParentWindowPositionChanged()
		wvTop.Resize()
	})

	ivBottom.SizeChanged().Attach(func() {
		wvBottom.NotifyParentWindowPositionChanged()
		wvBottom.Resize()
	})

	mw.Synchronize(func() {
		// win.ShowWindow(mw.Handle(), win.SW_SHOWMAXIMIZED)
		wvTop.Embed(uintptr(ivTop.Handle()))
		wvTop.Navigate("https://google.com")
		wvTop.Resize()

		wvBottom.Embed(uintptr(ivBottom.Handle()))
		wvBottom.Navigate("https://bing.com")
		wvBottom.Resize()
	})

	mw.Show()
	mw.Run()
}
