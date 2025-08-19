package main

import (
	"os"

	"github.com/jchv/go-webview2/pkg/edge"
	"github.com/lxn/walk"
)

// copilot.microsoft.com 에 질문한 내용과 github.com/pirogom/walkmgr 코드를 참고함
func main() {
	mw, _ := walk.NewMainWindow()

	mw.SetTitle("심플 브라우저")
	mw.SetSize(walk.Size{Width: 800, Height: 600})
	mw.SetLayout(walk.NewVBoxLayout())

	iv, _ := walk.NewImageView(mw)
	iv.SetMode(walk.ImageViewModeStretch)

	wv := edge.NewChromium()
	wv.SetPermission(edge.CoreWebView2PermissionKindClipboardRead, edge.CoreWebView2PermissionStateAllow)
	wv.DataPath = os.TempDir()

	iv.SizeChanged().Attach(func() {
		wv.NotifyParentWindowPositionChanged()
		wv.Resize()
	})

	mw.Synchronize(func() {
		wv.Embed(uintptr(iv.Handle()))
		wv.Navigate("https://google.com")
		wv.Resize()
	})

	mw.Show()
	mw.Run()
}
