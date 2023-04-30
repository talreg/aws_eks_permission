FROM golang
WORKDIR /goapp
COPY bin/ .
CMD  ./runner
