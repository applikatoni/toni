[![Build Status](https://travis-ci.org/applikatoni/toni.svg?branch=master)](https://travis-ci.org/applikatoni/toni)

# toni

toni is the CLI for Applikatoni.

## Installation

### Installing using `go get`

        go get github.com/applikatoni/toni

## Usage

    toni [-vh] [-c=<commit SHA>] [-b=<commit branch>] -t <target> -m <comment>

### Arguments

    REQUIRED:

      -t    Target of the deployment

      -m    The deployment comment

    OPTIONAL:

      -c    The deployment commit SHA (if unspecified toni uses the current git HEAD)
      -b    The branch of the commit SHA (if unspecified toni uses the current git HEAD)
      -h    Print the help and usage information
      -v    Print the version of the toni executable

## Configuration

When starting up, toni tries to read the ".toni.yml" configuration file in the
current working directoy. The configuration MUST specify the HOST, APPLICATION,
API TOKEN and the STAGES of deployments.

An example ".toni.yml" looks like this:

```yaml
host: http://toni.shippingcompany.com
application: shippingcompany-main-application
api_token: 4fdd575f-FOOO-BAAR-af1e-ce3e9f75367d
stages:
production:
  - CHECK_CONNECTION
  - PRE_DEPLOYMENT
  - CODE_DEPLOYMENT
  - POST_DEPLOYMENT
staging:
  - CHECK_CONNECTION
  - PRE_DEPLOYMENT
  - CODE_DEPLOYMENT
  - POST_DEPLOYMENT
```

# Contributing

All contributions are welcome! Create a fork, a new branch and send a pull
request!

Is the documentation lacking something? Did you find a bug? Do you have a
question about how Applikatoni works? Open a issue!


## License

MIT License. See [LICENSE](LICENSE).

## Authors
* [Thorsten Ball](https://github.com/mrnugget) ([@thorstenball](https://twitter.com/thorstenball))
