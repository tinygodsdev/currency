# Currency Normalization Package

The `currency` package provides functionality to normalize various currency symbols to their standard ISO 4217 codes. It also allows the use of custom mappings for additional flexibility.
Installation

To install the package, run:

```go get github.com/tinygodsdev/currency```

## Usage

### Normalizing Currency Symbols

You can use the Normalize function to normalize various currency symbols to their standard ISO 4217 codes. The function also supports custom mappings.

```go

package main

import (
    "fmt"
    "github.com/tinygodsdev/currency"
)

func main() {
    customMapping := map[string]string{
        "usd": "USDT",
        "custom": "CSTM",
    }

    fmt.Println(currency.Normalize("₽", nil))                // ("RUB", true)
    fmt.Println(currency.Normalize("usd", customMapping))    // ("USDT", true)
    fmt.Println(currency.Normalize("custom", customMapping)) // ("CSTM", true)
    fmt.Println(currency.Normalize("unknown", nil))          // ("unknown", false)
}
```

### Listing Supported Currencies

You can use the List function to get a list of all supported currencies.

```go

package main

import (
    "fmt"
    "github.com/tinygodsdev/currency"
)

func main() {
    fmt.Println(currency.List())
}
```

### Custom Mapping

If you have custom currency symbols or codes, you can pass them to the Normalize function using the customMapping parameter. The custom mapping takes precedence over the default mapping.

## Contributing

Contributions are welcome! Please submit a pull request on GitHub.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
