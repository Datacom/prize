package writer

import (
  "encoding/xml"
  "github.com/datacom/prize/reader"
)

type XML struct {
  XMLName xml.Name `xml:"Results"`
  DN      string   `xml:"DrawNumber"`
  DD      string   `xml:"DrawDate"`
  DT      string   `xml:"DrawTime"`
  WN      string   `xml:"WinningNumbers>Number"`
  Loc     []string `xml:"WinnerLocation,omitempty"`
}

func BuildXml(d *reader.DrawFile) (xml *XML, e error) {
  xml_output := &XML{DN: d.DrawNumber, DD: d.DrawDate, DT: d.DrawTime, WN: d.WinningNumbers}

  return xml_output, nil
}