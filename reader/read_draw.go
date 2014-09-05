package reader

import (
  "fmt"
  "io"
  "bufio"
)

type DrawFile struct {
  DrawNumber int
  DrawDate string
  DrawTime string
  WinningNumbers string
}

func ParseDraw(f io.Reader) (d DrawFile, err error) {
  fmt.Println("Parsing Draw File")
  fmt.Println("File: ", f) 

  draw := DrawFile{DrawNumber:123}

  var lines []string
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    fmt.Println("Line: ", scanner.Text())
    lines = append(lines, scanner.Text())
  }

  return draw, nil
}