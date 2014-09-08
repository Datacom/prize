package writer

import (
  "github.com/datacom/prize/reader"
)

func BuildEmail(d *reader.DrawFile) (email string, e error) {
  email_output := `
      Auckland, 16/03/2014 - Official Results.
      RELEASE FROM LOTTO NZ
      ---------------------
      OFFICIAL RESULT CERTIFICATE
      ---------------------------

      Draw Number : 3
      Draw Date : 16/03/2014
      Winning Numbers: 963

      ENDS...

      Issued on behalf of:

      The Chief Executive
      Lotto NZ
      Auckland
  `

  return email_output, nil
}