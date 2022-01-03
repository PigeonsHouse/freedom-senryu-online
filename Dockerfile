FROM golang:1.17.5

RUN mkdir -p $GOPATH/bin
RUN mkdir -p $GOPATH/pkg
RUN mkdir -p $GOPATH/src/fso

COPY . $GOPATH/src/fso
WORKDIR $GOPATH/src/fso

RUN go install

RUN go install github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]