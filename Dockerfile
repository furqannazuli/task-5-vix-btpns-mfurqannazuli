FROM golang:alpine

ENV DB_NAME     postgres
ENV DB_USER     postgres
ENV DB_HOST     localhost
ENV DB_PASSWORD passwordnya
ENV DB_PORT     5433
# development | testing | production
ENV STAGE       development

WORKDIR /app
COPY go.* ./
RUN go mod download

COPY . .
CMD go run main.go