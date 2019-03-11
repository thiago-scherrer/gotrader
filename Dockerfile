FROM golang:latest as test
LABEL name="gotrader_test"
ENV BOTDEV /src/gotrader
RUN mkdir -p $BOTDEV
WORKDIR ${BOTDEV}/
COPY internal/* ${BOTDEV}/
COPY configs/config-test.yml /opt/
RUN go test -args config /opt/config-test.yml

FROM golang:latest as build
LABEL name="gotrader-build"
ENV BOTDEV /src/gotrader
RUN mkdir -p $BOTDEV
WORKDIR ${BOTDEV}/
COPY internal/* ${BOTDEV}/
COPY configs/config.yml /opt/
RUN go build -o /bin/gotrader 

FROM golang:latest as run
LABEL name="gotrader"
COPY --from=build /opt/config.yml /opt/
COPY --from=build /bin/gotrader /bin/gotrader

CMD ["gotrader", "config", "/opt/config.yml"]
