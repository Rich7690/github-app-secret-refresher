FROM alpine:3.15
COPY github-app-secret-refresher app
ENTRYPOINT ["/app"]
