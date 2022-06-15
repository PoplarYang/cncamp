# build stage
FROM golang:alpine AS build
ADD . /src/
WORKDIR /src/
ENV GOPROXY=https://goproxy.cn
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o httpserver

# final stage
FROM busybox
WORKDIR /app
LABEL app=httpserver
ENV VERSION=1.2
COPY --from=build /src/httpserver /app/
EXPOSE 8080
ENTRYPOINT ["./httpserver"]
CMD ["-logtostderr=true"]
