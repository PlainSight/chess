# Chess

A chess game written in Go

## Building and running
```
go get github.com/go-gl/gl/v2.1/gl
go get github.com/go-gl/glfw/v3.2/glfw
go get -u github.com/jteeuwen/go-bindata/...
go-bindata pieces.png tile.png 
go build -o chess
.\chess
```
