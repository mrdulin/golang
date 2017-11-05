# concurrency

`sequence` 例子，每个http请求都是阻塞的，和Node.js中每个请求前使用 `await` 效果一样，三个请求是顺序执行的，结果是

```bash
☁  client [master] ⚡  go run client.go
2019/06/28 17:10:39 http://localhost:9090?echo=aaa
2019/06/28 17:10:41 rval1:  aaa
2019/06/28 17:10:41 http://localhost:9090?echo=bbb
2019/06/28 17:10:43 rval2:  bbb
2019/06/28 17:10:43 http://localhost:9090?echo=ccc
2019/06/28 17:10:45 rval3:  ccc
```

`concurrency` 例子，演示三个请求并发，结果是

```bash
☁  client [master] ⚡  go run client.go
2019/06/28 17:20:54 http://localhost:9090?echo=aaa
2019/06/28 17:20:54 http://localhost:9090?echo=bbb
2019/06/28 17:20:54 http://localhost:9090?echo=ccc
2019/06/28 17:20:56 rval1: aaa, rval2: bbb, rval3: ccc
```
