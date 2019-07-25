# ------------------------------------------------------------------------------
# Test and build
# ------------------------------------------------------------------------------
FROM golang:latest AS builder
LABEL NAME golang
LABEL VERSION 1.0
ENV APP /src/gotrader
WORKDIR ${APP}/src/gotrader
RUN mkdir -p ${APP}/src/gotrader
COPY . ${APP}/src/gotrader
COPY configs/config*.yml /opt/
RUN cd ${APP}/src/gotrader \
    && go mod download
RUN go test ./... -args config /opt/config-test.yml 
RUN cd cmd \
    && CGO_ENABLED=0 GOOS=linux \
    go build -a -installsuffix cgo -o /bin/gotrader \
    && useradd gotrader

# ------------------------------------------------------------------------------
# daemon image
# ------------------------------------------------------------------------------
FROM scratch AS runner
LABEL NAME scratch
LABEL VERSION 1.0
USER gotrader
COPY --from=builder /etc/ssl /etc/ssl
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /opt/*.yml /opt/
COPY --from=builder /bin/gotrader /bin/gotrader
CMD ["gotrader", "config", "/opt/config.yml"]
