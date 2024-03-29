config=local.toml

all: build

build:
	./build.sh

clean:
	rm -rf build/TaskRESTfulExercise build/TaskRESTfulExercise.linux.x64 build/TaskRESTfulExercise.windows.x64.exe build/TaskRESTfulExercise.mac.x64 build/TaskRESTfulExercise.mac.arm64

run:
	docker-compose up --build

run-local:
	docker-compose up --build -d
	go run main.go --config $(config) http --port 5566

.PHONY: clean build all