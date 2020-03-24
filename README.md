# protoc-gen-markdown

## install

```bash
go get github.com/lvht/protoc-gen-markdown
```

## generate markdown

```bash

protoc --plugin=protoc-gen-markdown=/Users/MS/Documents/goworkspace/src/protoc-gen-markdown/protoc-gen-markdown --markdown_out=. hello.proto

# set path prefix to /api
protoc --markdown_out=path_prefix=/api:. hello.proto
```
