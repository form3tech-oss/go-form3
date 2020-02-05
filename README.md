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
    f3 := form3.NewFromEnv()
	units := f3.Organisations.GetOrganisationUnits().MustDo()
	for _, unit := range units.Data {
		fmt.Printf("%s - %s\n", unit.ID, unit.Attributes.Name)
	}
	:print
```

### Scripts

`gorun cmd/sample_scripts/print_organisation_names.go`

See [cmd/sample_scripts](cmd/sample_scripts/) for a small subset of sample scripts

## Updating the Client

Tasks in the `Make` file download and preprocess the latest API swagger version.

* `make install-swagger` to install go-swagger if not already installed.
* `make` to regenerate the client files.
