FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO11MODULE=on
RUN go get github.com/mrifqi04/gomitra24/learn
RUN cd /build && git clone https://github.com/mrifqi04/gomitra24.git

RUN cd /build && go build

EXPOSE 8080

ENTRYPOINT [ "/build/server" ]
