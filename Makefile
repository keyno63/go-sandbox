BASEDIR=${PWD}
WORKDIR="${BASEDIR}/src/main/webapi"

.PHONY: build
build:
	@echo "Building now."
	@cd ${WORKDIR} && \
	  go build main.go person.go

.PHONY: run
run: build
	@echo "run"
	@cd ${WORKDIR} && \
	  go run main.go person.go
