package reader

import (
  // "bufio"
  // "io"
  // "net"
  // "net/http"
  // "net/http/httptest"
  "strings"
  "testing"
  "github.com/datacom/prize/reader"
)

func Test_Reader(t *testing.T) {

  drawData := `
    DrawNumber,1234
    WinningNumbers, 1 2 3`

  sreader := strings.NewReader(drawData)

  draw, err := reader.ParseDraw(sreader)  

  if err != nil {
    t.Error(err)
  }

  if !(draw.DrawNumber == 1234) {
    t.Error("expected 1234")
  }
}
