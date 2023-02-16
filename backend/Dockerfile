# syntax=docker/dockerfile:1
##
## Build
##
FROM golang:1.19-alpine as build

WORKDIR /go/src

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

## git追加
RUN apk add --no-cache git && \  
    ##   ビルドするOSをLinux、アーキテクチャをamd64に指定
    env GOOS=linux GOARCH=amd64  

RUN go build -o /go/bin/base ./main.go

CMD [/go/bin/base]

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /go/src

COPY --from=build /go/bin/base /go/bin/base

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/go/bin/base"]