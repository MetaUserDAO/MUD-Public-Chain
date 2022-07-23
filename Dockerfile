FROM alpine:latest

RUN apk add --no-cache bash
WORKDIR /root
COPY ./mudaod /root/.
#ENTRYPOINT ["tail", "-f", "/dev/null"]

