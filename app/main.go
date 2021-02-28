package main

import (
	"fmt"
	"log"

	"github.com/dzharii/mindb/webapi"
	"github.com/webview/webview"
)

func main() {
	go func() {
		err := webapi.StartServer("localhost:8081")
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("API Server has been started")
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintNone)

	// w.Navigate("http://localhost:3333/")
	w.Navigate(`data:text/html,
  <!doctype html>
  <html>
   <head><title>Hello</title></head>
   <body><h1>Hello, world!</h1>
   <div>
   <a href="http://localhost:3333">Click me!</a></body>
   </div>
   <div>
      <button onclick="document.location.href='http://localhost:3333'">Not a butt</button>
   </div>
   <div>
   <a href="https://google.com">Click me!</a></body>
   </div>
   <script> </script>
  </html>`)
	// w.Eval(`alert(document.location.href);`)
	//w.Navigate("http://example.com/")
	w.Run()

}

/*

# Linux
$ go build -o webview-example && ./webview-example

# MacOS uses app bundles for GUI apps
$ mkdir -p example.app/Contents/MacOS
$ go build -o example.app/Contents/MacOS/example
$ open example.app # Or click on the app in Finder

# Windows requires special linker flags for GUI apps.
# It's also recommended to use TDM-GCC-64 compiler for CGo.
# http://tdm-gcc.tdragon.net/download
$ go build -ldflags="-H windowsgui" -o webview-example.exe

*/
