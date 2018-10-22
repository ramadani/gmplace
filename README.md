# gmplace

Golang package for [Google Maps Place Autocomplete](https://developers.google.com/places/web-service/autocomplete)

## Installation

```cmd
go get github.com/ramadani/gmplace
```

## Usage

```go
gmplace := gmplace.New("YOUR_GOOGLE_API_KEY")
res, err := gmplace.Autocomplete("your keyword")

if err != nil {
  log.Println(err)
}

// do something with res
```

For result response, you can see it at https://developers.google.com/places/web-service/autocomplete#place_autocomplete_responses

## To Do

- [ ] Support for search by [Optional parameters](https://developers.google.com/places/web-service/autocomplete#place_autocomplete_requests)


## License

This library is distributed under the [MIT](LICENSE) license.