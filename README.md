# Form 3 Go Client

A simple go client for the Form 3 REST APIs built with custom operation names and swagger templates for a cleaner interface.

This is ideal for:

- Calling the API interactively (e.g. investigating resources or testing new features).
- Scripting (e.g. multi-call investigation or cleanup scripts)
- Testing, particularly when multiple calls to the API are involved.
- Load Testing, with K6 integration for initiating the requests

[gorun](https://github.com/erning/gorun) and [gore](https://github.com/motemen/gore) are highly recommended for running scripts and interactively developing scripts

## What's included

A simple client with:

- Automatic authentication and re-authentication based on environment variables (use `f3 creds` to set them)
- Pre-populated request objects with configurable defaults
- Sensible operation naming
- A fluent api with reduced repetition compared with vanilla swagger-go generated clients
- All the power and flexibility of Go!

## Getting Started

For interactive use with [gore](https://github.com/motemen/gore)
```go
go get -u github.com/form3tech-oss/go-form3
go get -u github.com/go-openapi/swag
go get -u github.com/go-openapi/strfmt
go get -u github.com/motemen/gore/cmd/gore
go get -u github.com/mdempsky/gocode
go get -u github.com/k0kubun/pp
go get -u github.com/google/uuid
```

### Required Environment Variables

| Environment variable | Description                                   |
|:---------------------|:----------------------------------------------|
| FORM3_HOST           | Form 3 host e.g. api.form3.tech               |
| FORM3_PUBLIC_KEY_ID  | Public key ID for request signing             |
| FORM3_PRIVATE_KEY    | Private key in PEM format for request signing |
| FORM3_ORGANISATION_ID| Organisation Id                               |
| DEBUG                | Output full HTTP calls and responses          |
| FORM3_CLIENT_ID      | Client id for token based auth (deprecated)   |
| FORM3_CLIENT_SECRET  | Secret for token based auth (deprecated)      |

### Local Tests

Run `gore -autoimport` then type commands interactively, with completion, history, and more. Gore will compile and run your script after each line, printing the output to the screen. Scripts can be output using `:print` and reset with `:clear`. Use `:help` for more details or [github](https://github.com/motemen/gore).

```go
    f3, _ := form3.NewFromEnv()
	units := f3.Organisations.GetOrganisationUnits().MustDo()
	for _, unit := range units.Data {
		fmt.Printf("%s - %s\n", unit.ID, unit.Attributes.Name)
	}
	:print
```

### Scripts

`go run examples/print_organisation_names.go`

See [examples](examples/) for a small subset of sample scripts.

## Updating the Client

### Pre-requisites

* [yq](https://github.com/mikefarah/yq) (:warning: the `python-yq` package is not compatible)
  * macOS: `brew install yq`
  * linux: `sudo add-apt-repository ppa:rmescandon/yq && sudo apt-get install yq`
* Python 3:
  * macOS: `brew install python`
  * linux: `sudo apt-get install python3.8`
* Python YAML package:
  * `pip3 install pyyaml`

### Regenerate client

Tasks in the `Makefile` download and preprocess the latest API swagger version:

1. `make install-swagger` to install go-swagger if not already installed.
1. `make` to regenerate the client files.

## License
Copyright 2019-2020 Form3 Financial Cloud

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
