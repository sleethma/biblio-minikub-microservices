# image
FROM golang:1.12.7-alpine

# get git
RUN apk update && apk upgrade && apk add --no-cache bash git

# grab gin library
RUN go get github.com/gin-gonic/gin

# set image env
ENV SOURCES /go/src/samples/go-kub-dock-services/
 
# copy from curr dir to run env
COPY . ${SOURCES}

# build service in working directory
RUN cd ${SOURCES} && CGO_ENABLED=0 go build

WORKDIR ${SOURCES}

# run service
CMD ./go-kub-dock-services

EXPOSE 8080