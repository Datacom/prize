package main

import (
  "fmt"
  "net/http"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  "github.com/datacom/prize/reader"
  "github.com/datacom/prize/writer"
)

func main() {
  m := martini.Classic()

  m.Use(render.Renderer(render.Options{
    IndentXML: true,
    Directory:       "templates",
    Layout:          "layout",
    Extensions:      []string{".tmpl", ".html"},
  }))

  // e.g.: curl -i -F file=@./data/draw1234.FIL http://localhost:3000/draw/email
  m.Get("/", ShowHelp)
  m.Post("/draw/email", RenderEmail)
  m.Post("/draw/xml",   RenderXml)
  m.Run()
}

func ShowHelp(r render.Render, req *http.Request, res http.ResponseWriter){
  r.HTML(200, "help", nil)
}

func RenderEmail(r render.Render, req *http.Request, res http.ResponseWriter) string {
  draw, err := ProcessDraw(req, res)

  if err != nil {
    fmt.Fprintln(res, err)
  }

  email, err := writer.BuildEmail(draw)

  return email
  // r.HTML(200, "email", draw)
}

func RenderXml(r render.Render, req *http.Request, res http.ResponseWriter){
  draw, err := ProcessDraw(req, res)

  if err != nil {
    fmt.Fprintln(res, err)
  }

  xml, err := writer.BuildXml(draw)

  r.XML(200, xml)
}

func ProcessDraw(req *http.Request, res http.ResponseWriter) (draw *reader.DrawFile, err error) {
  // Get an uploaded file (Content-Type multipart/form-data)
  file, header, f_err := req.FormFile("file")
    
  if f_err != nil {
    res.WriteHeader(400)    // Bad Request
    return nil, f_err
  }

  // Parse Draw File and return Draw Struct
  draw, d_err := reader.ParseDraw(header.Filename, file)

  if d_err != nil {
    res.WriteHeader(422)    // Unprocessable Entity
    return nil, d_err
  }

  return draw, err
}