FROM golang:1.20.4

WORKDIR /usr/src/job-tracker/server
COPY . .
RUN go mod tidy
EXPOSE 8080 
# configure ports later