FROM golang:1.20-bookworm

# install SQLite
RUN apt-get update && apt-get install -y sqlite3

# set current working directory inside the container
WORKDIR /app


# leverage layers: copy relevant files to install dependencies and thus cache them!
COPY /scraper/go.mod .
COPY /scraper/go.sum .
RUN go mod download

# then, copy the source code
COPY /scraper .

# build 
RUN go build -o scraper
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o scraper

# run
CMD ["./scraper"]
