FROM golang:1.21 AS builder
MAINTAINER lory <civet148@outlook.com>


RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 3B4FE6ACC0B21F32
RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 871920D1991BC93C

RUN echo "deb http://mirrors.aliyun.com/ubuntu/ focal main restricted universe multiverse" > /etc/apt/sources.list
RUN echo "deb http://mirrors.aliyun.com/ubuntu/ focal-security main restricted universe multiverse" >> /etc/apt/sources.list
RUN echo "deb http://mirrors.aliyun.com/ubuntu/ focal-updates main restricted universe multiverse" >> /etc/apt/sources.list
RUN echo "deb http://mirrors.aliyun.com/ubuntu/ focal-proposed main restricted universe multiverse" >> /etc/apt/sources.list
RUN echo "deb http://mirrors.aliyun.com/ubuntu/ focal-backports main restricted universe multiverse" >> /etc/apt/sources.list
RUN apt-get clean


RUN apt-get update && apt-get install -y ca-certificates make
ENV SRC_DIR /coin-ant
RUN set -x \
  && cd /tmp

RUN go env -w GOPROXY=https://goproxy.io

COPY . $SRC_DIR
RUN cd $SRC_DIR && export GIT_SSL_NO_VERIFY=true && git config --global http.sslVerify "false" && make

FROM ubuntu:22.04

RUN ln -fs /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" > /etc/timezone && apt-get update && apt-get install -y tzdata
ENV TZ Asia/Shanghai
ENV SRC_DIR /coin-ant


COPY --from=builder $SRC_DIR/coin-ant /usr/local/bin/coin-ant
COPY --from=builder /etc/ssl/certs /etc/ssl/certs


ENV HOME_PATH /data

VOLUME $HOME_PATH
