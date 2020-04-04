# protoc-gen-markdown

## install

```bash
go get github.com/withgame/protoc-gen-markdown
```

## generate markdown

```bash
# set path prefix to /api
protoc --markdown_out=path_prefix=/api:. hello.proto
```
