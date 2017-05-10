# TinyUrl

a url shorten web service written by Golang,Vue and Gin.

## Demo
browse [tinyUrl demo website](http://tinyurl.adolphlwq.xyz) and enjoy yourself.

## Directory structure
```
.
├── be
│   ├── bin
│   │   └── tinyurl
│   ├── Makefile
│   ├── pkg
│   │   └── linux-amd64
│   ├── src
│   │   └── tinyurl
│   └── vendor
│       ├── manifest
│       └── src
├── fe
│   ├── index.html
│   └── index.js
└── README.md
```
- `be` is the backend directory,it contains the Golang backend project.I use [gb](https://github.com/constabulary/gb) to manage Golang project.
- `fe` is the front end directory,it contains web pages and extra files needed.

## TODOs
- [ ] validate input url format
- [ ] improve random generate string algorithm
- [ ] reserch [wrk](https://github.com/wg/wrk)