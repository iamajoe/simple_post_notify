# build stage
FROM golang:1.19-alpine AS builder

LABEL maintainer="Joel Santos <joe@joesantos.io>"

WORKDIR /app
COPY . .

RUN apk update && apk add --virtual build-dependencies build-base gcc git make

RUN go mod download
RUN go build -v -o ./app

# final stage
FROM alpine

ARG TZ="Etc/UTC"
ARG ENV="production"
ARG ALLOWED_ORIGINS="*"

ARG TELEGRAM_ENABLE
ARG TELEGRAM_ID
ARG TELEGRAM_SECRET

ENV PORT=4040
ENV TZ=${TZ}
ENV ENV=${ENV}
ENV ALLOWED_ORIGINS=${ALLOWED_ORIGINS}

ENV TELEGRAM_ENABLE=${TELEGRAM_ENABLE}
ENV TELEGRAM_ID=${TELEGRAM_ID}
ENV TELEGRAM_SECRET=${TELEGRAM_SECRET}

WORKDIR /root/
COPY --from=builder /app/app ./app

CMD ["./app"]
