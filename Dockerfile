FROM golang

COPY ./ $GOPATH/src/git/red-coins/src/github.com/bayerlein/red-coins/app
WORKDIR $GOPATH/src/git/red-coins/src/github.com/bayerlein/red-coins/app

RUN go get github.com/bayerlein/red-coins
RUN go get github.com/go-chi/chi
RUN go get golang.org/x/crypto/bcrypt

RUN go install -v ./...
RUN go build
EXPOSE 8080
CMD ["./app"]
