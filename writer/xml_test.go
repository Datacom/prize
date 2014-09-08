package writer

import (
  "testing"
  "github.com/datacom/prize/reader"
)

func TestXmlWriter(t *testing.T) {

  draw := &reader.DrawFile{DrawNumber: "1234", DrawDate: "16/03/2014", DrawTime: "18:05", WinningNumbers: "123"}

  xml, err := BuildXml(draw)

  if err != nil {
    t.Error(err)
  }
  
  if !(xml.DN == "1234") {
    t.Error("expected Draw Number to be 1234")
  }
}