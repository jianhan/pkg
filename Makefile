.PHONY: pb data lint run

pb:
	for f in proto/**/*.proto; do \
		protoc --go_out=plugins=grpc,micro:. $$f; \
		echo compiled: $$f; \
	done
