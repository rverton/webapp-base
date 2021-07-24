package main

import (
	"embed"
	"io/fs"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

var (
	// filled during compile time
	commitHash string
	commitDate string

	//go:embed public
	embededFiles embed.FS
)

func main() {
	debugMode := os.Getenv("DEBUG") != ""

	e := echo.New()
	setupLogger(e, debugMode)

	// options
	e.HideBanner = true
	e.Renderer = NewRenderer(debugMode)
	e.Debug = debugMode

	// middlewares
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} => ${status}\n",
	}))
	e.Use(middleware.Recover())

	// static assets
	assetHandler := http.FileServer(getFileSystem(debugMode, e.Logger))
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", assetHandler)))

	// routes
	routes(e)

	e.Logger.Fatal(e.Start(":8000"))
}

func setupLogger(e *echo.Echo, debugMode bool) {
	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level}")
	}

	if debugMode {
		e.Logger.SetLevel(log.INFO)
		e.Logger.Info("debug mode on, log level set to INFO")

		if commitHash != "" {
			e.Logger.Printf("commit: %v, date: %v", commitHash, commitDate)
		}
	} else {
		e.Logger.Info("debug mode off")
	}
}

func getFileSystem(useOS bool, logger echo.Logger) http.FileSystem {
	if useOS {
		logger.Info("using assets from filesystem")
		return http.FS(os.DirFS("public"))
	}

	logger.Info("using assets embedded in binary")
	fsys, err := fs.Sub(embededFiles, "public")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
