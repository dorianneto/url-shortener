FROM golang:1.20

WORKDIR /usr/src/app

RUN git config --global --add safe.directory /usr/src/app && \
    go install github.com/cosmtrek/air@latest

CMD ["air"]
