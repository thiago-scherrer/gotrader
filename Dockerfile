FROM golang:latest as test
LABEL name="gotrader_test"
ENV GOPATH /src/gotrader
ENV GO111MODULE on

RUN mkdir -p $GOPATH/src/gotrader
WORKDIR ${GOPATH}/src/gotrader

COPY . ${GOPATH}/src/gotrader
COPY configs/config-test.yml /opt/

RUN cd internal/central \
    && go test -args config /opt/config-test.yml 
RUN cd internal/convert \
    && go test -args config /opt/config-test.yml 
RUN cd internal/display \
    && go test -args config /opt/config-test.yml 

FROM golang:latest as builder
LABEL name="gotrader-build"
ENV GOPATH /src/gotrader
ENV GO111MODULE on

RUN mkdir -p $GOPATH/src/gotrader
COPY . ${GOPATH}/src/gotrader
WORKDIR ${GOPATH}/src/gotrader/cmd/main/
COPY configs/config.yml /opt/
RUN go build -o /bin/gotrader 

FROM golang:latest as run
LABEL name="gotrader"
COPY --from=builder /opt/config.yml /opt/
COPY --from=builder /bin/gotrader /bin/gotrader

CMD ["gotrader", "config", "/opt/config.yml"]
