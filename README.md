# Disposable

A golang package and command line tool for the analysis and identification of
disposable email addresses.

## Package Usage

```golang
import "github.com/prophittcorey/disposable"

disposable.Check("someone@10minutemail.ru") // => true, err

disposable.Domains() // => ["10minutemail.ru", ...]
```

## License

The source code for this repository is licensed under the MIT license, which you can
find in the [LICENSE](LICENSE.md) file.
