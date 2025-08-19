package main

import (
	"os"

	"github.com/jchv/go-webview2/pkg/edge"
	"github.com/lxn/walk"
	"github.com/lxn/win"
)

// copilot.microsoft.com 에 질문한 내용과 github.com/pirogom/walkmgr 코드를 참고함
func main() {
	_main()
}
func _main() {
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
	ivBottom.SetMode(walk.ImageViewModeCenter)

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
