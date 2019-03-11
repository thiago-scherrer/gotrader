FROM golang:latest as test
LABEL name="gotrader_test"
ENV BOTDEV /src/gotrader
RUN mkdir -p $BOTDEV
WORKDIR ${BOTDEV}/
COPY internal/* ${BOTDEV}/
COPY configs/config-test.yml /opt/
RUN go test -args config /opt/config-test.yml

FROM golang:latest as build
LABEL name="gotrader"
ENV BOTDEV /src/gotrader
RUN mkdir -p $BOTDEV
WORKDIR ${BOTDEV}/
COPY internal/* ${BOTDEV}/
COPY configs/* /opt/
RUN go build -o /bin/gotrader 

CMD ["gotrader", "config", "/opt/config.yml"]
