FROM golang:1.16.7-alpine

WORKDIR /var/www/app

RUN apk update && \
    apk add --update --no-cache --virtual=.build-dependencies \
    git

# no parameter
# CMD ["go", "run", "src/main.go"]
CMD ["go", "run", "src/main.go", "default"]
