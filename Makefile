RUN_NAME="messengerBot"

build:
	mkdir -p output/
	go build -o output/${RUN_NAME}
test: build
	export GIN_MODE=debug && export env=test && exec output/${RUN_NAME} :80
prod: build
	export GIN_MODE=release && export env=prod && exec output/${RUN_NAME} :80