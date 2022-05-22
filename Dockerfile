FROM golang as builder
WORKDIR /app/

COPY go.work go.work.sum /app/

ADD ./apps apps/
ADD ./engines engines/
ADD ./pkg pkg/

RUN CGO_ENABLED=0 GOOS=linux go build blackgo/api

FROM alpine:latest as api
WORKDIR /root/
COPY --from=builder /app/api /root/
EXPOSE 8080
CMD [ "./api" ]