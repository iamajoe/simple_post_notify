# Simple post notify

> A simple server to handle a post and notify to one of the available modules

***

#### Run

```sh
go get .
DEBUG=true go run .
```

#### Available ENV variables

```
ENV                 # development | production
DEBUG               # true | false
ALLOWED_ORIGINS     # "http://*;https://*"

TELEGRAM_ENABLE     # true | false
TELEGRAM_ID         # <telegram_chat_id>
TELEGRAM_SECRET     # <telegram_chat_secret>
```

#### Build with docker

```sh
docker build -t simple_post_notify:latest --build-arg TELEGRAM_ENABLE=true --build-arg TELEGRAM_ID=example_id --build-arg TELEGRAM_SECRET=shh_secret .
```
