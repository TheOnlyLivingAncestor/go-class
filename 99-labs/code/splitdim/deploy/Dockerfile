# 1. Compile the go binary in a golang environment

FROM docker.io/golang:latest as build

COPY . /

WORKDIR /

RUN CGO_ENABLED=0 GOOS=linux go build -a -o splitdim

# 2. Build the container image

FROM scratch

COPY --from=build /splitdim .
COPY --from=build /static ./static

CMD ["/splitdim"]