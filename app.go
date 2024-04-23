package main

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func isImageFile(mimetype string) bool {
	switch mimetype {
	case "image/jpeg":
	case "image/png":
	case "image/webp":
	case "image/gif":
	default:
		return false
	}

	return true
}

func walkFilePath(location string) []string {
	filePaths := []string{}

	e := filepath.WalkDir(location, func(path string, d fs.DirEntry, e error) error {
		if e != nil {
			log.Fatal(e)
		}

		if !d.IsDir() {
			filePaths = append(filePaths, path)
		}

		return nil
	})

	if e != nil {
		fmt.Println(e)
	}

	return filePaths
}

func (a *App) GetImageFilePaths() []string {
	options := runtime.OpenDialogOptions{
		DefaultDirectory: "/run/media/sabix/T7/Art/Goal Super",
		Title:            "title",
	}

	path, e := runtime.OpenDirectoryDialog(a.ctx, options)
	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(path)
	filePaths := walkFilePath(path)

	fmt.Println(&filePaths)

	imagePaths := []string{}

	for _, path := range filePaths {
		file, e := os.ReadFile(path)
		if e != nil {
			log.Fatal(e)
		}

		mimetype := http.DetectContentType(file)

		if isImageFile(mimetype) {
			imagePaths = append(imagePaths, path)
		}
	}

	return imagePaths
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
