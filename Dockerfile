FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR .
ADD main .
ADD ./Config/application.yaml ./Config/application.yaml


EXPOSE 8001
CMD ["./main","--config=./Config/application.yaml"]
