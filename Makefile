test:
	go test $$(go list ./...)

fmt:
	go fmt $$(go list ./...)
