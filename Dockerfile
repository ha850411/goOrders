FROM --platform=linux/amd64 centos:7

EXPOSE 80 
EXPOSE 443

RUN mkdir /app \
    && yum -y update \
    && yum install gcc git wget ImageMagick -y && yum clean all

ENV GO_VERSION 1.19.5

RUN wget https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz  \
    && tar -C /usr/local -zxvf go${GO_VERSION}.linux-amd64.tar.gz \
    && rm /go${GO_VERSION}.linux-amd64.tar.gz

ENV PATH /usr/local/go/bin:$PATH

ENV GOROOT /usr/local/go
RUN \cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

ENTRYPOINT go run main.go >> ./logs/run.log 2>&1
