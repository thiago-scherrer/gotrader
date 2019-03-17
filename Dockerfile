# ------------------------------------------------------------------------------
# Test image
# ------------------------------------------------------------------------------
FROM golang@sha256:908ea6b956394d7a7006453e6a16011a6f86fd47996f2ccc32711f1eeff6b9fc AS test
ENV APP /src/gotrader
WORKDIR ${APP}/src/gotrader
RUN mkdir -p ${APP}/src/gotrader
COPY . ${APP}/src/gotrader
COPY configs/config-test.yml /opt/
RUN cd ${APP}/src/gotrader/modules \
    && go mod download
RUN cd internal/central \
    && go test -args config /opt/config-test.yml 
RUN cd internal/convert \
    && go test -args config /opt/config-test.yml 
RUN cd internal/display \
    && go test -args config /opt/config-test.yml 

# ------------------------------------------------------------------------------
# build image
# ------------------------------------------------------------------------------
FROM test AS builder
ENV APP /src/gotrader
WORKDIR ${APP}/src/gotrader/cmd/main/
COPY configs/config.yml /opt/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/gotrader 
RUN useradd gotrader

# ------------------------------------------------------------------------------
# daemon image
# ------------------------------------------------------------------------------
FROM scratch AS runner
USER gotrader
COPY --from=builder /etc/ssl /etc/ssl
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /opt/config.yml /opt/
COPY --from=builder /bin/gotrader /bin/gotrader
CMD ["gotrader", "config", "/opt/config.yml"]
