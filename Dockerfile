FROM golang:1.8

RUN go get github.com/hjma29/ovcli
WORKDIR /go/src/github.com/hjma29/ovcli
RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."
#CMD ["go-wrapper", "run"] # ["app"]
