CURRENT_DIR=$(shell pwd)

proto-gen:
		./scripts/proto-gen.sh ${CURRENT_DIR}

start:
	go run cmd/main.go

.PHONY:compile