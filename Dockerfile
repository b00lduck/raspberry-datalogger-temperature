FROM hypriot/rpi-golang
WORKDIR /gopath1.5/src/b00lduck/raspberry-datalogger-temperature
ENTRYPOINT ["raspberry-datalogger-temperature"]

ADD . /gopath1.5/src/b00lduck/raspberry-datalogger-temperature
RUN go get
RUN go build
