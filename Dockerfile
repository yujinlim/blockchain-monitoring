FROM golang:1.8.3

LABEL maintainer="Yu Jin"

ARG package_name=github.com/yujinlim/blockchain-monitoring
ARG workdir=$GOPATH/src/$package_name

WORKDIR $GOPATH

ADD . $workdir
RUN go install $package_name

ENTRYPOINT ["blockchain-monitoring"]
