## Go Api Template

Api in go template repository.

Requirements:

- go >= 1.22.2
- mkcert - https://github.com/FiloSottile/mkcert
- pre-commit - https://pre-commit.com/
- golang-ci - https://github.com/golangci/golangci-lint
- yq - https://github.com/mikefarah/yq
- Migoro - https://github.com/brownhounds/migoro
- make

### Working with local packages

```bash
go mod edit -replace=github.com/username/swift=/home/brownhounds/dev/local-package
```

Alternatively add section in `go.mod` file:

```
replace github.com/username/swift => /home/brownhounds/dev/local-package
```
