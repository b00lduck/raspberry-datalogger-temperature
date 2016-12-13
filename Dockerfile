FROM rem/rpi-golang-1.7:latest

WORKDIR /gopath/src/github.com/b00lduck/raspberry-datalogger-temperature
ENTRYPOINT ["raspberry-datalogger-temperature"]

ADD . /gopath/src/github.com/b00lduck/raspberry-datalogger-temperature
RUN go get
RUN go build
