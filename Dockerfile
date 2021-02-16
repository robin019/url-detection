FROM golang:1.15.7-alpine3.13

WORKDIR /url_detection

COPY . /url_detection

CMD sh run.sh