FROM golang:1.11 AS builder
WORKDIR ./
COPY . ./
RUN make build

FROM scratch
COPY --from=builder ./builds/topnservice /
ENTRYPOINT [ "./topnservice" ]

