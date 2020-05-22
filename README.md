# TinyURL - Make Personal URL Shortening Service Easy

[![Build Status](https://travis-ci.com/tinyurl/tinyurl.svg?branch=master)](https://travis-ci.com/tinyurl/tinyurl)  [![Go Report Card](https://goreportcard.com/badge/github.com/tinyurl/tinyurl)](https://goreportcard.com/report/github.com/tinyurl/tinyurl)  [![GoDoc](https://godoc.org/github.com/tinyurl/tinyurl?status.svg)](https://godoc.org/github.com/tinyurl/tinyurl)

<p align="center">
  <a href="https://url.algcs.xyz" target="_blank">
    <img src="assets/tinyurl.gif" width="700px">
    <br>
    Live Demo
  </a>
</p>

<p align="center">a url shorten web service written by Golang, Vue and Gin.</p>

## Requisites
- Golang(1.8+)
- [Govendor](https://github.com/kardianos/govendor)
- MySQL/Sqlite3
- make

## Quick Start
1. clone project to **GOPATH**
```
git clone https://github.com/tinyurl/tinyurl.git $GOPATH/src/github.com/tinyurl/tinyurl
```
2. sync golang packages
```
govendor sync
```
3. build binary
```
make
```
4. change config in default.properties
```
# tinyurl dev config
#app
app.host=0.0.0.0
app.port=8877
app.domain=0.0.0.0:8877

# db config
db.type=sqlite3
db.path=.
db.host=
db.port=
db.name=tinyurldb
db.user=
db.password=
```
5. run binary
```
./tinyurl -config default.properties
```
6. open index.html in `frontend/` with broswer

## TODOs
- [X] validate input url format
- [X] improve random generate string algorithm
    - [X] use math/rand.Read instead math/rand.Intn func
- [X] use logrus replace golang log lib
- [X] reserch [wrk](https://github.com/wg/wrk)
- [X] add test case
- [ ] Backend
  - [X] data type support multi database(index in sender)
  - [X] adjust short path generating algorithm
  - [ ] custom short url
  - [ ] API rate
  - [ ] Swagger for api management
  - [ ] Admin account?
  - [ ] count each url parse time (high concurrent situation)
- [ ] Frontend
  - [ ] qrcode support

## Reference
- [GitHub-YOURLS](https://github.com/YOURLS/YOURLS)
- [知乎-短 URL 系统是怎么设计的？](https://www.zhihu.com/question/29270034)
- [GitHub/Ourls](https://github.com/takashiki/Ourls)
- [GitHub/uriuni](https://github.com/dchest/uniuri)