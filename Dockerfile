FROM public.ecr.aws/docker/library/golang:alpine AS build

WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
RUN go build -o main .

FROM public.ecr.aws/docker/library/alpine:latest
COPY --from=build /app/main /usr/local/bin

ENTRYPOINT ["main"]
