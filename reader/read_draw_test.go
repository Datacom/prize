package reader

import (
  "strings"
  "testing"
)

func TestReader(t *testing.T) {

  drawData := `DrawNumber,1234
WinningNumbers,123`

  sreader := strings.NewReader(drawData)

  draw, err := ParseDraw(sreader)

  if err != nil {
    t.Error(err)
  }

  if !(draw.DrawNumber == "1234") {
    t.Error("expected 1234")
  }
}
