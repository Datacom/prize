## Prize

Prize is a demo API application that process input draw files and create output in different formats, e.g. XML, Email.

### Setup

* [Install Go](https://golang.org/doc/install).

* [Set up you work environment](https://golang.org/doc/code.html), e.g.

```
$ mkdir $HOME/projects/go
$ export GOPATH=$HOME/projects/go
$ export PATH=$PATH:$GOPATH/bin
```

* Create a directory inside your workspace in which to keep the source code:

```
$ mkdir -p $GOPATH/src/github.com/datacom
```

* Clone the project from github:

```
$ cd $GOPATH/src/github.com/datacom
$ git clone git@github.com:Datacom/prize.git
```

* Build and install the program:

```
$ go install github.com/datacom/prize
```

* Run the program:

```
$ prize
```
Sample output:
```
[martini] listening on :3000 (development)
```

### Sample requests using curl

#### Get draw results in xml format:
Run curl command to send HTTP POST request and upload a Draw File:

```
curl -i -F file=@./data/draw1234.FIL http://localhost:3000/draw/xml
```
Example response:

```
HTTP/1.1 100 Continue

HTTP/1.1 200 OK
Content-Type: text/xml; charset=UTF-8
Date: Tue, 09 Sep 2014 23:19:29 GMT
Content-Length: 176

<Results>
  <DrawNumber>3</DrawNumber>
  <DrawDate>16/03/2014</DrawDate>
  <DrawTime>18:05</DrawTime>
  <WinningNumbers>
    <Number>963</Number>
  </WinningNumbers>
</Results>
```

#### Get draw results in email format:
Run curl command to send HTTP POST request and upload a Draw File:

```
curl -i -F file=@./data/draw1234.FIL http://localhost:3000/draw/email
```
Example response:

```
HTTP/1.1 100 Continue

HTTP/1.1 200 OK
Content-Type: text/html; charset=UTF-8
Date: Tue, 09 Sep 2014 23:18:01 GMT
Content-Length: 418

<html>
  <head></head>
  <body>
    <pre>
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
  </pre>
  </body>
</html>
```