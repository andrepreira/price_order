FROM golang:1.20.1

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN apt-get update && apt-get install -y build-essential librdkafka-dev

CMD ["tail", "-f", "/dev/null"]