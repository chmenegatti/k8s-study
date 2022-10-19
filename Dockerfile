FROM golang
WORKDIR /go/src/github.com/k8s-study
COPY . .
RUN go build
EXPOSE 3001
CMD ["./k8s-study"]