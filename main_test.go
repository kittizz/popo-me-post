package main

import (
	"log"
	"os"
	"testing"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
)

func TestAS(t *testing.T) {
	// Initialize astilectron
	var a, _ = astilectron.New(log.New(os.Stderr, "", 0), astilectron.Options{
		// AppName:            "<your app name>",
		// AppIconDefaultPath: "<your .png icon>",  // If path is relative, it must be relative to the data directory
		// AppIconDarwinPath:  "<your .icns icon>", // Same here
		// BaseDirectoryPath:  "<where you want the provisioner to install the dependencies>",
		// VersionAstilectron: "<version of Astilectron to utilize such as `0.33.0`>",
		// VersionElectron:    "<version of Electron to utilize such as `4.0.1` | `6.1.2`>",
	})
	defer a.Close()

	// Start astilectron
	a.Start()
	var w, _ = a.NewWindow("https://google.com", &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(600),
		Width:  astikit.IntPtr(600),
	})
	w.Create()
	// Blocking pattern
	a.Wait()
}
