test:
	$(call announce,"Running tests")
	@go test -cover ./...

