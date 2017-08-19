# TinyURL

[![Build Status](https://travis-ci.org/adolphlwq/tinyurl.svg?branch=master)](https://travis-ci.org/adolphlwq/tinyurl)  [![Go Report Card](https://goreportcard.com/badge/github.com/adolphlwq/tinyurl)](https://goreportcard.com/report/github.com/adolphlwq/tinyurl)

a url shorten web service written by Golang, Vue and Gin.

## Demo
browse [tinyUrl demo website](http://tinyurl.adolphlwq.xyz) and enjoy yourself.

## Requisites
- Golang(1.8+)
- [Govendor](https://github.com/kardianos/govendor)
- MySQL
- make

## Quick Start
1. clone project to **GOPATH**
```
git clone https://github.com/adolphlwq/tinyurl.git $GOPATH/src/tinyurl
```
2. sync golang packages
```
govendor sync
```
3. build binary
```
make
```
4. run binary
```
./tinyurl -dbname tinyurl -user user -pass pass -dbport 2306
```

## TODOs
- [X] validate input url format
- [X] improve random generate string algorithm
    - [X] use math/rand.Read instead math/rand.Intn func
- [X] use logrus replace golang log lib
- [X] reserch [wrk](https://github.com/wg/wrk)
- [X] add test case
- [ ] dynamic adjust short path length (default is 4)
- [ ] count each url parse time (high concurrent situation)
- [ ] qrcode support
- [ ] list api

## Reference
- [GitHub/Ourls](https://github.com/takashiki/Ourls)
- [GitHub/uriuni](https://github.com/dchest/uniuri)
