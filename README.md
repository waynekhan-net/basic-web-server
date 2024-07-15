# README

```text
go test -v -cover -covermode=count -mod=vendor -coverprofile=coverage.out -json > report.json && \
  sonar-scanner -X
```
