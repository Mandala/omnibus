The Omnibus Project
====================

Omnibus is a collection of high-level web application API package for Go.
Whether you build REST API endpoint to full featured website or something in
between, it got you covered. Omnibus subpackages is designed with modularity in
mind so you can easily mix and match it with another package.

## Quick Start Guide

To use any subpackage in Omnibus, you should install it first with `go get`.

```
go get -u github.com/mandala/omnibus
```

Then, just use `import` statement in your code to use the omnibus subpackages.

```go
import (
    "github.com/mandala/omnibus/serve"
    "github.com/mandala/omnibus/browser"
    "github.com/mandala/omnibus/view"
)
```

## Documentation

Please refer to each subpackage folder for more information on its API docs
and usage examples.

## Benchmarks

Well, if you care so much about performance benchmarks please help us out to
tighten up the code and create proper benchmarking codes.

## Issues and Feature requests

Please let us know on <https://github.com/mandala/omnibus/issues>.

## Contributing

Please install Editor Config and Golang support on your favorite IDE before
start tinkering the code. Then, when you're done please send a pull request
with proper documentation.

## License

Copyright (c) 2017 Fadhli Dzil Ikram. All rights reserved.

This project is licensed under MIT license that can be found on the LICENSE
file.
