FROM golang:1.22 AS build
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /server .

FROM gcr.io/distroless/base-debian12:latest
WORKDIR /
COPY --from=build /server ./server
EXPOSE 3000
USER nonroot:nonroot
ENTRYPOINT ["/server"]
