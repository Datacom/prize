package reader

import (
  "fmt"
  "io"
  "bufio"
  "strings"
)

type DrawFile struct {
  DrawNumber string
  DrawDate string
  DrawTime string
  WinningNumbers string
}

func ParseDraw(fname string, f io.Reader) (d *DrawFile, e error) {
  fmt.Printf("Processing File: %v\n", fname)
  draw := new(DrawFile)
  var err error

  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    line       := scanner.Text()
    line_args  := strings.Split(line, ",")

    switch line_args[0] {
      case "DrawNumber":      draw.DrawNumber = line_args[1]
      case "DrawDate":        draw.DrawDate = line_args[1]
      case "DrawTime":        draw.DrawTime = line_args[1]
      case "WinningNumbers":  draw.WinningNumbers = line_args[1]
      default: err = fmt.Errorf("[Bad Draw Error]: Unrecognized symbol %q", line_args[0])
    }
  }

  return draw, err
}