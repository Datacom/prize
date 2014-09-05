package main

import (
  "fmt"
  "net/http"
  // "encoding/json"
  // "encoding/xml"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  // "strings"
  // "regexp"
  "github.com/datacom/prize/reader"
)

func main() {
  m := martini.Classic()

  // render html templates from templates directory
  m.Use(render.Renderer())
  m.Use(martini.Recovery())
  m.Use(martini.Logger())

  // request using curl:
  // curl -i -F file=@./data/draw1234.FIL http://localhost:3000/draw
  m.Post("/draw", func(params martini.Params, r *http.Request, w http.ResponseWriter) string {
    file, header, err := r.FormFile("file")
    
    if err != nil {
      fmt.Fprintln(w, err)
      return "Error"
    }

    fmt.Printf("File Name: %v\n", header.Filename)
    draw := reader.ParseDraw(file)
    fmt.Printf("Draw Number: %v\n", draw.DrawNumber)

    return "posted draw"
  })

  m.Run()
}

func ListDraws(r render.Render) (string, error) {
  return "List of Draws", nil
  // return r.HTML(200, "hello", "world")
}

func ShowDraw(r *http.Request, params martini.Params) (string, error) {
  return "Show draw with id: " + params["id"], nil
}

// func CreateDraw(r *http.Request, params martini.Params) (string, error) {
//   return "nada", nil
// }

// func UpdateDraw(r *http.Request, params martini.Params) (string, error) {
//   return "nada", nil
// }
// func DeleteDraw(r *http.Request, params martini.Params) (string, error) {
//   return "nada", nil
// }

/////////////////////////////////////////////////////////////////////////

  // m.Group("/draws", func(r martini.Router) {
  //   r.Get("/list", ListDraws)
  //   r.Get("/:id", ShowDraw)
    // r.Post("/new", CreateDraw)
    // r.Put("/update/:id", UpdateDraw)
    // r.Delete("/delete/:id", DeleteDraw)
  // })

  // Setup routes
  // r := martini.NewRouter()
  // m.Get("/draws", ListDraws)
  // m.Get("/draws/:id", ShowDraw)
  // r.Post("/draws", CreateDraw)
  // r.Put("/draws/:id", UpdateDraw)
  // r.Delete("/draws/:id", DeleteDraw)
  // m.Action(r.Handle)

/////////////////////////////////////////////////////////////////////////
  // Handlers - any kind of callable function
  // m.Get("/", func(r render.Render) (int, string) {
  //   return 300, "Hello world!" // HTTP 200 : "hello world"
  // })

  // m.Get("/", func(r render.Render) string {
  //   return "Hello world!"
  //   // return r.HTML(200, "hello", "world")
  // })

  // m.Get("/", func(res http.ResponseWriter, req *http.Request) { // res and req are injected by Martini
  //   res.WriteHeader(200) // HTTP 200
  // })

  // m.Get("/hello/:name", func(params martini.Params) string {
  //   return "Hello " + params["name"]
  // })