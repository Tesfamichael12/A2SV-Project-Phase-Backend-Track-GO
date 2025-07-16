# Step-by-Step Guide: Adding Swagger Documentation to a Go Gin API

This is a simple, step-by-step guide to help set up Swagger API documentation for a Go Gin project.

---

## 1. Install Required Packages

First, install the Swagger tools for Gin:

```bash
go get -u github.com/swaggo/gin-swagger
```

Then, install the static files package:

```bash
go get -u github.com/swaggo/files
```

And finally, install the CLI tool for generating docs:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

To verify the installation, run:

```bash
swag -h
```

You should see output similar to:

```
NAME:
    swag - Automatically generate RESTful API documentation with Swagger 2.0 for Go.

USAGE:
    swag [global options] command [command options] [arguments]

VERSION:
    v1.8.1

COMMANDS:
    init, i  Create docs.go
    fmt, f   format swag comments
    help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
    --help, -h     show help (default: false)
    --version, -v  print the version (default: false)
```

Make sure `$GOPATH/bin` is in your PATH so you can run `swag` from the terminal.

---

## 2. Add Swagger Annotations to Your Code

Swagger uses special comments above your handler functions to generate documentation. Here’s how to annotate your endpoints:

```go
// album represents a music album.
// @Description Album object
// @name album
// @property ID string "Album ID" example:"1"
// @property Title string "Album title" example:"Blue Train"
// @property Artist string "Artist name" example:"John Coltrane"
// @property Price float64 "Album price" example:"56.99"
type album struct {
    ID     string  `json:"id" example:"1"`
    Title  string  `json:"title" example:"Blue Train"`
    Artist string  `json:"artist" example:"John Coltrane"`
    Price  float64 `json:"price" example:"56.99"`
}
```

Annotate your handler functions like this:

```go
// @Summary List albums
// @Description Get all albums in the store
// @Tags albums
// @Produce json
// @Success 200 {array} album "List of albums"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /albums [get]
func getAlbums(cxt *gin.Context) {
    cxt.IndentedJSON(http.StatusOK, albums)
}

// @Summary Add a new album
// @Description Add a new album to the store
// @Tags albums
// @Accept json
// @Produce json
// @Param album body album true "Album to add"
// @Success 201 {object} album "Created album"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /albums [post]
func postAlbums(cxt *gin.Context) {
    // existing code
}

// @Summary Get album by ID
// @Description Get details of an album by its ID
// @Tags albums
// @Produce json
// @Param id path string true "Album ID"
// @Success 200 {object} album "Album details"
// @Failure 404 {object} map[string]string "Album not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /albums/{id} [get]
func getAlbumByID(cxt *gin.Context) {
    // existing code
}
```

---

## 3. Generate Swagger Documentation

Run the following command in your project directory:

```bash
swag init
```

This will create a `docs` folder with Swagger files (`docs.go`, `swagger.json`, `swagger.yaml`).

---

## 4. Import the Docs Package

In your `main.go`, import the generated docs package. If your module name is `example/web-service-gin` check your `go.mod` file for your module name value, add:

```go
import _ "example/web-service-gin/docs"
```

This ensures the docs are included in your build.

---

## 5. Serve Swagger UI in Gin

Add the following to your `main.go`:

```go
import (
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    // other imports
)

func main() {
    router := gin.Default()
    // other routes

    url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

    router.Run("localhost:8080")
}
```

---

## 6. View Your API Docs

Start your server:

```bash
go run .
```

Open your browser and go to:

```
http://localhost:8080/swagger/index.html
```

You should see the full Swagger UI with all your endpoints, models, and documentation.

---

## 7. Update Docs When You Change Endpoints

Whenever you add or change annotations, run:

```bash
swag init
```

This keeps your documentation up to date.

---

## 8. Tips for Advanced Usage

- **Authentication:** Add security annotations for JWT or API keys as needed.
- **Response Examples:** Use `example` tags in your struct fields for better UI.
- **Error Documentation:** Document error responses with `@Failure` tags.
- **Grouping:** Use `@Tags` to organize endpoints.

---

## 9. Troubleshooting

- If you see only JSON, make sure you visit `/swagger/` not `/swagger/doc.json`.
- If endpoints don’t show, check your annotations and run `swag init` again.
- If you get a 500 error, make sure the docs package is imported and the `docs` folder is present.

---

## Conclusion

Swagger makes your Go Gin API easy to explore, test, and share.

Happy coding!
