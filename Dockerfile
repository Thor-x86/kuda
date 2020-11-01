FROM golang:1-alpine

COPY . /usr/src
WORKDIR /usr/src
RUN go build -o ../bin

ENV KUDA_PUBLIC_DIR "/srv"
ENV KUDA_DOMAIN "localhost"
ENV KUDA_PORT "8080"
ENV KUDA_ORIGINS ""
ENV KUDA_PORT_TLS ""
ENV KUDA_CERT ""
ENV KUDA_KEY ""

CMD kuda ${KUDA_PUBLIC_DIR} --domain=${KUDA_DOMAIN} --port=${KUDA_PORT} --origins=${KUDA_ORIGINS} --portTLS=${KUDA_PORT_TLS} --cert=${KUDA_CERT} --key=${KUDA_KEY}