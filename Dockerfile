FROM golang:1.20-rc-bullseye as builder
WORKDIR /app

RUN apt-get install ca-certificates && update-ca-certificates
RUN curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh | bash -s -- --to /usr/local/bin
COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN just build

FROM alpine:3.15

RUN apk add --no-cache ca-certificates
COPY --from=builder /app/dist/diagon-alley /var/run/diagon-alley
COPY --from=builder /app/dist/diagon-alley-create-user /var/run/diagon-alley-create-user
COPY --from=builder /app/dist/ /bin

ENV APP_USER=appuser
RUN addgroup -S $APP_USER && adduser -S $APP_USER -G $APP_USER
RUN chown -R $APP_USER:$APP_USER /var/run/diagon-alley
USER $APP_USER

EXPOSE 9000 9000
ARG release
ENV RELEASE_SHA $release

CMD ["sh", "-c", "/var/run/diagon-alley"]
