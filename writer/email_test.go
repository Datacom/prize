package writer

import (
  "testing"
  "github.com/datacom/prize/reader"
)

func TestEmailWriter(t *testing.T) {

  draw := &reader.DrawFile{DrawNumber: "1234", DrawDate: "16/03/2014", DrawTime: "18:05", WinningNumbers: "123"}

  xml, err := BuildEmail(draw)

  if err != nil {
    t.Error(err)
  }

  if !(xml == "sample email") {
    t.Error("expected result: sample email")
  }
}