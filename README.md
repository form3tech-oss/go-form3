![Form3 logotype](form3.png)

# Form3 Go client library

A simple go client library for the [Form3 REST APIs](https://api-docs.form3.tech/api.html) built with custom operation names and swagger templates for a cleaner interface.

This is ideal for:

- Calling the API interactively (e.g. investigating resources or testing new features).
- Scripting (e.g. multi-call investigation or cleanup scripts)
- Testing, particularly when multiple calls to the API are involved.
- Load Testing, with K6 integration for initiating the requests

## What's included

A simple client library with:

- Automatic authentication and re-authentication based on environment variables
- Pre-populated request objects with configurable defaults
- Sensible operation naming
- A fluent API with reduced repetition compared with vanilla swagger-go generated clients
- All the power and flexibility of Go!

## Getting started

For interactive usage we recommend using [gore](https://github.com/motemen/gore).

Note that the usage of client IDs and secrets is deprecated but still supported using `form3.WithTokenTransport`.
Mutual TLS (not yet available in all environments) is supported using `form3.NewMTLSTransport`.

### Example usage

```go
u, err := url.Parse(ts.URL)
if err != nil {
    log.Fatal(err)
}

o, err := uuid.Parse(os.Getenv("FORM3_ORGANISATION_ID"))
if err != nil {
    log.Fatal(err)
}

f3, err := form3.New(
    form3.WithBaseURL(*u),
    form3.WithOrganisationID(o),
    form3.WithRequestSigningTransport(
        form3.NewRequestSigningTransport(
            form3.WithPrivateKey(privKey),
            form3.WithPublicKeyID(pubKeyID))))
if err != nil {
    log.Fatal(err)
}
```


### Environment variables

The following environment variables need to be set when constructing the client
using `form3.NewFromEnv`

| Environment variable   | Description                                   |
|:-----------------------|:----------------------------------------------|
| `FORM3_HOST`           | Form3 host URL, e.g. https://api.form3.tech   |
| `FORM3_PUBLIC_KEY_ID`  | Public key ID for request signing             |
| `FORM3_PRIVATE_KEY`    | Private key in PEM format for request signing |
| `FORM3_ORGANISATION_ID`| Organisation ID                               |
| `DEBUG`                | Output full HTTP calls and responses          |

## Updating the client

The client is generated using `go-swagger`.

### Prerequisites

* [yq](https://github.com/mikefarah/yq#install)
* Python 3 with `pyaml`

### Regenerate client

Tasks in the `Makefile` download and preprocess the latest API swagger version:

1. `make install-swagger` to install go-swagger if not already installed.
1. `make` to regenerate the client files.

## License

Copyright 2019-2021 Form3 Financial Cloud

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
