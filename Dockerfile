FROM golang:latest

LABEL name="gotrader"

ADD . /root/
WORKDIR /root/internal/gotrader/

RUN go get gopkg.in/yaml.v2 && go build -o gotrader

CMD ["./gotrader", "config", "../../configs/config.yml"]
