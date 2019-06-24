# error stack

执行源码`main.go`

```bash
☁  error-stack [master] ⚡  go run main.go
open test.txt: no such file or directory
open file error
main.loadFile
        /Users/ldu020/workspace/github.com/mrdulin/golang/src/error-handling/error-stack/main.go:24
main.main
        /Users/ldu020/workspace/github.com/mrdulin/golang/src/error-handling/error-stack/main.go:13
runtime.main
        /usr/local/go/src/runtime/proc.go:200
runtime.goexit
        /usr/local/go/src/runtime/asm_amd64.s:1337%
```

执行 build 后的二进制可执行文件`error-stack`

```bash
☁  golang [master] ⚡  ./bin/error-stack
open test.txt: no such file or directory
open file error
main.loadFile
        /Users/ldu020/workspace/github.com/mrdulin/golang/src/error-handling/error-stack/main.go:24
main.main
        /Users/ldu020/workspace/github.com/mrdulin/golang/src/error-handling/error-stack/main.go:13
runtime.main
        /usr/local/go/src/runtime/proc.go:200
runtime.goexit
        /usr/local/go/src/runtime/asm_amd64.s:1337%
```
