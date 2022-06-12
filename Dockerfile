FROM golang:1.18-bullseye

WORKDIR /app

COPY meiki_server/ .

RUN go mod download


RUN go build -o meiki

# TODO: ENABLE MULTI STAGE
##
## Deploy
##
# FROM gcr.io/distroless/base-debian10

# WORKDIR /app

# COPY --from=build /meiki /meiki

EXPOSE 80
ENV PORT=80

# ENV MEIKI_DATABASE_URL

# USER nonroot:nonroot

# ENTRYPOINT ["/meiki"]
CMD [ "./meiki" ]