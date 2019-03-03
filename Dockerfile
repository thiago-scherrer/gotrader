FROM golang:latest

LABEL name="gotrader"

WORKDIR /opt/gotrader

COPY internal/ /opt/
COPY configs/config.yml /opt/gotrader

RUN go get gopkg.in/yaml.v2 && \
    go build -o /bin/gotrader 

CMD ["gotrader", "config", "/opt/gotrader/config.yml"]
