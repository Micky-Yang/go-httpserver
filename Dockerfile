#FROM golang:1.16 as build
#WORKDIR /app
#COPY ./go.mod .
#COPY ./go.sum .
#ENV GOPROXY https://proxy.golang.com.cn,direct
#RUN go mod download
#COPY main.go .
#RUN go build -o /build/httpserver .

#FROM centos:centos7 as run
#COPY --from=build /build/httpserver /httpserver
#ENTRYPOINT ["/httpserver"]

FROM centos:centos7 as run
COPY httpserver /httpserver
ENTRYPOINT ["/httpserver"]