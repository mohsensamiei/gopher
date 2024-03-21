GOPATH=${HOME}/go

%:
	@true

.PHONY: fmt
fmt:
	scripts/fmt.sh
