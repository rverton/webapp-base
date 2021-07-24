# Skeleton for a go web app

This is a very simple and minimal skeleton for a web app.

used libraries:
* [echo](https://github.com/labstack/echo) as a minimal web handler
* [pongo2](https://github.com/flosch/pongo2) for template renderering (django2 style)
* [tailwindcss](https://tailwindcss.com) for styling

features:
* assets and templates are embedded when building (via go1.16+ embed module)
* base layout for extending
* macros for reusable components
* git commit hash and date injected during build for `-version` flag
* `make` for development and building

## Usage

### Development 

* Use `make tailwind-dev` to watch for templates changes and recompile tailwind styles during development.
* Use `DEBUG=1` env variabel to enable debug logging and to deactivate file embedding

```
$ export DEBUG=1
$ make tailwind-dev
$ go build && ./webapp
```

### Building for production

This will build the app binary, compile tailwind styles and embed all static
files from `./public`.

```
$ make
```

The resulting binary contains all templates and assets.
