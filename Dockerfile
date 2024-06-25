FROM golang:1.22

WORKDIR /build

COPY . .

ENV CGO_ENABLED=1

RUN go get 
RUN go build -o bin .

ENTRYPOINT ["/build/bin"]

#docker build . -t threadcoredocker:latest
#docker run -e PORT=8080 -p 8080:8080 threadcoredocker:latest