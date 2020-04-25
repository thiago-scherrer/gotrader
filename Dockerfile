# ------------------------------------------------------------------------------
# build
# ------------------------------------------------------------------------------
FROM golang:1.14.2 AS builder

LABEL NAME golang
LABEL VERSION 2.0

ENV APP /src/gotrader

WORKDIR ${APP}/src/gotrader

RUN mkdir -p ${APP}/src/gotrader

COPY . ${APP}/src/gotrader

RUN cd ${APP}/src/gotrader \
    && go mod download 

RUN mv configs/*.yml /opt/ \
    && cd ${APP}/src/gotrader \
    && cd cmd \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/gotrader \
    && useradd gotrader

# ------------------------------------------------------------------------------
# daemon image
# ------------------------------------------------------------------------------
FROM scratch AS runner

LABEL NAME scratch
LABEL VERSION 2.0

USER gotrader

ENV GOTRADER_MODE testnet

COPY --from=builder /etc/ssl /etc/ssl
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /bin/gotrader /bin/gotrader

CMD ["gotrader"]
