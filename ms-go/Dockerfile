FROM golang:1.20-alpine as golang

RUN apk add -U tzdata
RUN apk --update add ca-certificates

WORKDIR /app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server .

FROM scratch

COPY --from=golang /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=golang /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=golang /etc/passwd /etc/passwd
COPY --from=golang /etc/group /etc/group
ENV TZ=America/Sao_Paulo

COPY --from=golang /server .

EXPOSE 3030

CMD ["/server"]