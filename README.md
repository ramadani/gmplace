# gmplace

Golang package for [Google Maps Place Autocomplete](https://developers.google.com/places/web-service/autocomplete)

## Installation

```cmd
go get github.com/ramadani/gmplace
```

## Usage

```go
import "github.com/ramadani/gmplace"

gmplace := gmplace.New("YOUR_GOOGLE_API_KEY")
res, err := gmplace.Autocomplete(map[string]string{
  "input":    "your keyword",
  "language": "your language id",
})

if err != nil {
  log.Println(err)
}

// do something with res
```

For result response, you can see it at https://developers.google.com/places/web-service/autocomplete#place_autocomplete_responses

## License

This library is distributed under the [MIT](LICENSE) license.