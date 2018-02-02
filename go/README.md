## playground-go

http://techblog.raccoon.ne.jp/archives/37168009.html

### init
- `brew install go`
- `go env GOROOT`
- `go env GOPATH` working dir
	- `/Users/bandai/Documents/repos/go_learn`
- http://qiita.com/1000ch/items/e42e7c28cf7a7b798a02

### command
- `go run`
- `go fmt`
- `go get`
- https://speakerdeck.com/mitsuse/golang-texin-wowu-nisitekotowoshu-ku

### writing method
- package
- import
- main function
- sub function
	- Public
	- private

### test
- `mymath.go` <=> `mymath_test.go`
- `go test mymath -v -cover`
	- `-cpuprofile`
	- `-parallel`
	- `go help testflag`

### iota
- http://yosssi.hatenablog.com/entry/2014/01/27/000135
