# build stage
FROM golang:alpine AS build
ADD . /src/
RUN export GOPROXY=https://goproxy.cn && \
    cd /src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o httpserver

# final stage
FROM busybox
WORKDIR /app
LABEL app=httpserver
ENV VERSION=1.0
COPY --from=build /src/httpserver /app/
EXPOSE 8080
ENTRYPOINT ["./httpserver"]
CMD ["-logtostderr=true"]
