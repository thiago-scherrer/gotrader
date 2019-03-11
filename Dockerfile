FROM golang:latest

LABEL name="gotrader"

ENV BOTDEV /src/gotrader

RUN mkdir -p $BOTDEV

WORKDIR ${BOTDEV}/

COPY internal/* ${BOTDEV}/
COPY configs/* /opt/

RUN go build -o /bin/gotrader 

CMD ["gotrader", "config", "/opt/config.yml"]
