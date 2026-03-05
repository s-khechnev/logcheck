# logcheck

## Checks los messages


### Building
```bash
git clone https://github.com/s-khechnev/logcheck
cd logcheck
make build
build/logcheck help
```

### As a golangci-lint plugin

Build a custom version of golangci-lint with the plugin:
```bash
make gen_golangci_lint
```
This creates an executable `./custom-gcl` that can be used instead of the standard golangci-lint

### Run tests
```bash
make test
```
