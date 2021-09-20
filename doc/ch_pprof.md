```
cd src/ch_pprof
go run ch_pprof.go > cpu.pprof
```

1. web 查看
需要先安装 Graphviz
```
// ubuntu
apt install graphviz
// macOS
brew install graphviz
```

```
go tool pprof -http=:9999 cpu.pprof
```

2. 工具命令行
```
go tool pprof cpu.pprof
top
top --cum
help
```