FROM golang:1.25

WORKDIR /usr/src/app

RUN git config --global --add safe.directory /usr/src/app && \
    go install github.com/air-verse/air@latest

CMD ["air"]
