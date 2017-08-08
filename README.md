# TinyUrl

a url shorten web service written by Golang,Vue and Gin.

## Demo
browse [tinyUrl demo website](http://tinyurl.adolphlwq.xyz) and enjoy yourself.

## Directory structure
```
➜  tinyurl git:(master) ✗ tree
.
├── api.go
├── db.go
├── fe
│   ├── index.html
│   └── index.js
├── main.go
├── Makefile
├── README.md
├── service.go
└── service_test.go
```

`fe` is the front end directory,it contains web pages and extra files needed.

## TODOs
- [X] validate input url format
- [X] improve random generate string algorithm
    - [X] use math/rand.Read instead math/rand.Intn func
- [X] use logrus replace golang log lib
- [X] reserch [wrk](https://github.com/wg/wrk)
- [ ] add QR code support
- [ ] add test case
- [ ] dynamic adjust short path length (default is 4)
- [ ] count each url parse time (high concurrent situation)
- [ ] qrcode support
- [ ] list api

## Reference
- [GitHub/Ourls](https://github.com/takashiki/Ourls)
