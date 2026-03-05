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

#### Available settings:
1. -loggers `comma sep string`. Supported: `zap`, `slog`. Example setting golangci-lint
   [here](https://github.com/s-khechnev/logcheck/blob/master/.golangci.yml#L13)

### Run tests
```bash
make test
```
