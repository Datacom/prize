package writer

import (
  "github.com/datacom/prize/reader"
  "text/template"
  "bytes"
  "log"
)

func BuildEmail(draw *reader.DrawFile) (email string, e error) {
  tmpl, err :=template.ParseFiles("templates/email.tmpl")
  
  if err != nil {
    log.Fatal(err)
  }

  var doc bytes.Buffer
  
  err = tmpl.Execute(&doc, draw)
  
  if err != nil {
    log.Fatal(err)
  }

  return doc.String(), nil
}