package main

import (
	"os"

	"github.com/jchv/go-webview2/pkg/edge"
	. "github.com/lxn/walk/declarative"
)

func main() {
	_main_view := initWebview()

	MainWindow{
		Title:  "심플 브라우저",
		Layout: VBox{},
		Children: []Widget{
			ImageView{
				Layout: VBox{}.
				
			},
			Composite{
				Layout: VBox{},
				// AssignTo: &main_view,
			},
		},
	}.Run()
}

func initWebview() *edge.Chromium {
	// https://github.com/pirogom/walkmgr/blob/main/webview2.go 내용 참고

	wv := edge.NewChromium()
	wv.SetPermission(edge.CoreWebView2PermissionKindClipboardRead, edge.CoreWebView2PermissionStateAllow)
	wv.DataPath = os.TempDir()

	return wv
}
