package writer

import (
  "github.com/datacom/prize/reader"
  "strings"
)

func BuildEmail(d *reader.DrawFile) (email string, e error) {
  email_output := `
      Auckland, [draw_date] - Official Results.
      RELEASE FROM LOTTO NZ
      ---------------------
      OFFICIAL RESULT CERTIFICATE
      ---------------------------

      Draw Number : [draw_number]
      Draw Date : [draw_date]
      Winning Numbers: [winning_numbers]

      ENDS...

      Issued on behalf of:

      The Chief Executive
      Lotto NZ
      Auckland
  `
  email_output = strings.Replace(email_output, "[draw_date]", d.DrawDate, -1)
  email_output = strings.Replace(email_output, "[draw_number]", d.DrawNumber, -1)
  email_output = strings.Replace(email_output, "[winning_numbers]", d.WinningNumbers, -1)

  return email_output, nil
}