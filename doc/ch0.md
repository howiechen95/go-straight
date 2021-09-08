## Golang 环境配置

- [安装包下载地址1](https://golang.org/dl/)
- [安装包下载地址2](https://golang.google.cn/dl/)

> ### windows
<details>
<summary>展开查看</summary>
<pre><code>

```

```

</code></pre>
</details>

> ### linux
<details>
<summary>展开查看</summary>
<pre><code>

```

```

</code></pre>
</details>

> ### Mac OS X（Darwin）
<details>
<summary>展开查看</summary>
<pre><code>

```

```

</code></pre>
</details>

## go 编译

- CGO_ENABLED=1 启用CGO
- GOOS=linux 指定编译平台为linux，
- GOARCH=amd64 指定架构为 amd64
- CGO_LDFLAGS="-static" static表示静态连接，没有这个选项，linux上运行报：-bash: ./xxx: /lib/ld-musl-x86_64.so.1: bad ELF interpreter: No such file or directory
- CC=x86_64-linux-musl-gcc macOS 上交叉编译指定gcc 
- CXX=x86_64-linux-musl-g++ macOS 上交叉编译指定c++
- -ldflags '-w -s'

在macOS 上启用cgo编译，并交叉编译linux平台的二进制文件
```
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CGO_LDFLAGS="-static" CC=x86_64-linux-musl-gcc  CXX=x86_64-linux-musl-g++ go build -ldflags '-w -s' .
```

