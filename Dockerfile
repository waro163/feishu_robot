FROM golang:1.15.2

ADD . /feishu_robot
WORKDIR /feishu_robot
ENV TZ Asia/Shanghai
ENV GOPROXY "https://goproxy.cn,direct"
RUN go mod vendor
EXPOSE 8001

RUN go build -mod=vendor -o main main.go

CMD ["./main"]