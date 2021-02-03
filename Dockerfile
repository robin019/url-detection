FROM golang:1.15.7

WORKDIR /url_detection

COPY . /url_detection

CMD sh run.sh