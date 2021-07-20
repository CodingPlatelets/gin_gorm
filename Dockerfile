FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /gin_web
COPY . /gin_web
RUN go build -o gin_web .

EXPOSE 8000
ENTRYPOINT ["./gin_web"]