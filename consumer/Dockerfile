FROM golang:alpine

RUN apk add --no-cache git build-base

# Install librdkafka
RUN apk add --update --no-cache alpine-sdk bash ca-certificates \
      libressl \
      tar \
      git openssh openssl yajl-dev zlib-dev cyrus-sasl-dev openssl-dev build-base coreutils
WORKDIR /root
RUN git clone https://github.com/edenhill/librdkafka.git
WORKDIR /root/librdkafka
RUN ./configure
RUN make
RUN make install
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

WORKDIR /app
COPY . /app

RUN go mod init microservices/consumer
RUN go get github.com/confluentinc/confluent-kafka-go/v2/kafka
RUN CGO_ENABLED=1 GOOS=linux go build -a -tags musl -o consumer .

CMD ["./consumer"]
