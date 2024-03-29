flags="-X main.version=0.1.0  -X 'TaskRESTfulExercise/cmd.goversion=$(go version)' -X TaskRESTfulExercise/cmd.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X TaskRESTfulExercise/cmd.githash=`git describe --always --long --dirty --abbrev=14`"

echo "$flags"

# # build current
go build -ldflags "$flags" -o build/TaskRESTfulExercise -o build/TaskRESTfulExercise

# build for Linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$flags" -o build/TaskRESTfulExercise.linux.x64 main.go 

# build for Windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "$flags" -o build/TaskRESTfulExercise.windows.x64.exe main.go 

# build for Mac (intel)
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "$flags" -o build/TaskRESTfulExercise.mac.x64 main.go  

# build for Mac (apple)
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "$flags" -o build/TaskRESTfulExercise.mac.arm64 main.go