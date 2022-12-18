# Disposable

[![Go Reference](https://pkg.go.dev/badge/github.com/prophittcorey/disposable.svg)](https://pkg.go.dev/github.com/prophittcorey/disposable)

A golang package and command line tool for the analysis and identification of
disposable email addresses.

## Package Usage

```golang
import "github.com/prophittcorey/disposable"

disposable.Check("someone@10minutemail.ru") // => true, err

disposable.Domains() // => ["10minutemail.ru", ...]
```

## Tool Usage

```bash
# Install the latest tool.
$ go install github.com/prophittcorey/disposable/cmd/disposable@latest

# Dump all domains.
$ disposable --domains

# Check a specific email address or domain.
$ disposable --check someone@10minutemail.ru
```

## License

The source code for this repository is licensed under the MIT license, which you can
find in the [LICENSE](LICENSE.md) file.
