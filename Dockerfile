FROM golang:alpine

ENV DB_NAME     postgres
ENV DB_USER     postgres
# replace with postgres_db if you run with containe
ENV DB_HOST     localhost
ENV DB_PASSWORD mysecretpassword
ENV DB_PORT     5433
# development | testing | productio
ENV STAGE       testing

# create working directory
# copy go.sum if exist
# and do cache if exist
WORKDIR /app
COPY go.* ./
RUN go mod download

COPY . .
CMD go run main.go