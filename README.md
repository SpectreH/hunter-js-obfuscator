<h1>Hunter - Go Javascript Obfuscator</h1>

This package is a Hunter Obfuscator written in Go to protect your JavaScript source code with the simplest and fastest way.

## Installation
```
go get github.com/spectreh/hunter-js-obfuscator
```
## Usage
```golang
import "github.com/spectreh/hunter-js-obfuscator"

rand.Seed(time.Now().UnixNano())

code := "alert('hello');"

obf := NewObfuscator(code)
obfuscatedCode := obf.Obfuscate()
```

## Additional information

Only standard go packages were in use.

## Author

* SpectreH (https://github.com/SpectreH)

## License

This package is licensed under **MIT** license.
