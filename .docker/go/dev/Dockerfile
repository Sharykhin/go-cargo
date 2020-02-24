FROM golang:1.13

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

ARG install_compile_deamon

RUN if [ $install_compile_deamon = "on" ] ; then go get github.com/githubnemo/CompileDaemon ; echo "Installing CompileDeamon is turned off" ; fi

COPY . .
EXPOSE 8000