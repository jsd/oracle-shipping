<p align="center">
  <img src="https://s3.linkpool.io/images/bridgestype.png">
</p>

[![Build Status](https://travis-ci.org/linkpoolio/bridges.svg?branch=master)](https://travis-ci.org/linkpoolio/bridges)
[![codecov](https://codecov.io/gh/linkpoolio/bridges/branch/master/graph/badge.svg)](https://codecov.io/gh/linkpoolio/bridges)
[![Go Report Card](https://goreportcard.com/badge/github.com/linkpoolio/bridges)](https://goreportcard.com/report/github.com/linkpoolio/bridges)
-----------------------

Bridges is a Chainlink adaptor framework, lowering the barrier of entry for anyone to create their own:

- A tested hardened library that removes the need to build your own HTTP server, allowing you to just focus on 
adapter requirements.
- Simple interface to allow you to build an adapter that confides to Chainlink schema.
- Kept up to date with any changes, meaning no extra work for existing adapters to support new schema changes or 
features.
- Supports running in serverless environments such as AWS Lambda & GCP functions with minimal effort.

## Contents
1. [Code Examples](#code-examples)
2. [Running in AWS Lambda](#running-in-aws-lambda)
3. [Running in GCP Functions](#running-in-gcp-functions)


## Running in Docker
After implementing your bridge, if you'd like to run it in Docker, you can reference the Dockerfiles in 
[examples](examples/cryptocompare/Dockerfile) to then use as a template for your own Dockerfile.

## Running in AWS Lambda
After you've completed implementing your bridge, you can then test it in AWS Lambda. To do so:

1. Build the executable:
    ```bash
    GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o bridge
    ```
2. Add the file to a ZIP archive:
    ```bash
    zip bridge.zip ./bridge
    ```
3. Upload the the zip file into AWS and then use `bridge` as the
handler.
4. Set the `LAMBDA` environment variable to `true` in AWS for
the adaptor to be compatible with Lambda.

## Running in GCP Functions
Due to the difference in running Go within GCP Functions, it requires specific considerations for it be supported 
within your bridge:
- Bridge implementation cannot be within the `main` package
- An extra `Handler` function within your implementation:
    ```go
    func Handler(w http.ResponseWriter, r *http.Request) {
        bridges.NewServer(&Example{}).Handler(w, r)
    }
    ```
- A `go.mod` and `go.sum` within the sub-package that contains the `Handler` function

For an example implementation for GCP Functions, view the 
[asset price adapter](https://github.com/linkpoolio/asset-price-cl-ea).

You can then use the gcloud CLI tool to deploy it, for example:
```bash
gcloud functions deploy bridge --runtime go111 --entry-point Handler --trigger-http
```






### Contributing

We welcome all contributors, please raise any issues for any feature request, issue or suggestion you may have.
