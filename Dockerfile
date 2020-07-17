FROM golang:1.14

WORKDIR $GOPATH/src/github.com/caesargusti/retrospective

COPY . .
