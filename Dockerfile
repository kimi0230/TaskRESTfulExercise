# build stage
FROM golang:1.22-alpine AS build-env
WORKDIR '/go/src/TaskRESTfulExercise/api'
ADD . .
RUN apk add git
RUN go build -ldflags "-X main.version=0.1.0  -X 'TaskRESTfulExercise/cmd.goversion=$(go version)' -X TaskRESTfulExercise/cmd.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X TaskRESTfulExercise/cmd.githash=`git describe --always --long --dirty --abbrev=14`" -o APPSERVICE

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/TaskRESTfulExercise/api/APPSERVICE /app/
COPY --from=build-env /go/src/TaskRESTfulExercise/api/config.toml /app/
# RUN ["mkdir","logs"]
RUN ["chmod", "+x", "APPSERVICE"]
RUN ["ls", "-al"]
EXPOSE 8080

# ENTRYPOINT ./APPSERVICE
CMD ./APPSERVICE http 8080