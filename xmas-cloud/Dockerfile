FROM node:15.2 as nodemodules
WORKDIR /src
COPY package.json package-lock.json ./
# Download all dependencies. Dependencies will be cached if the package*.json are not changed
RUN npm install
RUN chmod -R go+rX .

FROM golang:1.15.5 as builder
WORKDIR /src
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .
RUN cd cryptor; CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .
RUN chmod -R go+rX .

FROM scratch
WORKDIR /app
USER 1001
COPY --from=nodemodules /src/node_modules ./node_modules
COPY --from=builder /src/static ./static
COPY --from=builder /src/xmas-cloud .
COPY --from=builder /src/cryptor/cryptor /bin/xmas-cryptor
COPY --from=builder /src/flag /flag
EXPOSE 8080
CMD ["./xmas-cloud"]