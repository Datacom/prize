package main

import (
  // "fmt"
  "net/http"
  // "encoding/json"
  // "encoding/xml"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  // "strings"
  // "regexp"
)

type DrawFile struct {
  DrawNumber int
  DrawDate string
  DrawTime string
  WinningNumbers string
}

func main() {
  m := martini.Classic()

  // render html templates from templates directory
  m.Use(render.Renderer())
  m.Use(martini.Recovery())
  m.Use(martini.Logger())
  
  m.Get("/", func(r render.Render) {
    r.HTML(200, "hello", "world")
  })

  // Setup routes
  r := martini.NewRouter()
  r.Get("/draws", ListDraws)
  r.Get("/draws/:id", ShowDraw)
  r.Post("/draws", CreateDraw)
  r.Put("/draws/:id", UpdateDraw)
  r.Delete("/draws/:id", DeleteDraw)
  m.Action(r.Handle)

  m.Run()
}

func ListDraws(r render.Render, params martini.Params) {
  return r.HTML(200, "hello", "world")
}
func CreateDraw(r *http.Request, params martini.Params) (string, error) {
  return "nada", nil
}
func ShowDraw(r *http.Request, params martini.Params) (string, error) {
  return "nada", nil
}
func UpdateDraw(r *http.Request, params martini.Params) (string, error) {
  return "nada", nil
}
func DeleteDraw(r *http.Request, params martini.Params) (string, error) {
  return "nada", nil
}

