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

### Example of usage

1. Non-English messages
```go
slog.Info("Привет мир") // "Message contains non-English letter: Привет мир"
slog.Info("User created", "роль", "admin") // "Message contains non-English letter: роль"
```

2. Lowercase messages
```go
slog.Info("Привет мир") // "Message starts with capital letter:: Привет мир"
```

3. No sensitive Data
```go
slog.Info("user password: " + password) // "Message contains sensitive data: password"
slog.Info("user auth", slog.String("password", password)) // "Message contains sensitive data: password"
```

3. No special chars and emoji
```go
slog.Info("Hello 👋") // "Message contains special char or emoji: Hello 👋"
slog.Info("Temperature: 25°C") // "Message contains special char or emoji: Temperature: 25°C"
```

Try: 
```bash
build/logcheck internal/logcheck/nosensitivedata/testdata/slog/slog.go
```

And a lot of others in `internal/logcheck/*/testdata/*/*.go`
