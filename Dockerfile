#Golang version 1.18.1
FROM golang:latest

# 工作路径
WORKDIR /app

#Copy Code
COPY src/ .

#Build
ENV GOPROXY https://goproxy.cn
RUN go mod download
RUN go build -o main .

# 设置命令
CMD ["./main"]
