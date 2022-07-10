FROM golang:1.18-bullseye

WORKDIR /app

COPY meiki_server/ .
RUN go mod download
RUN go build -o meiki

EXPOSE 443
ENV PORT=443

CMD [ "./meiki" ]
