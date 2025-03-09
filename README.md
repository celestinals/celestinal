![ico](favicon.ico)

# 🎫 Tickex

[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0)

Buying, selling, exchanging, and sharing all types of tickets and game cards.

Tickex is a platform that allows users to buy, sell, exchange, and share various types of tickets and game cards easily and securely. Our goal is to provide a seamless experience for ticket enthusiasts while ensuring transparency and trust among users.

## Requirement

- Docker engine
- Go 1.24.0
- Protocol Buffer
- Make (for running commands efficiently)

## Installation

**Run the command below to clone all project's source code or Clone code with SSH Key** 
- HTTPs
    ```sh
    git clone --recurse-submodules -j8 https://github.com/tickexvn/tickex.git $GOPATH/src/github.com/tickexvn/tickex
    ```
- SSH Key
    ```sh
    git clone --recurse-submodules -j8 git@github.com:tickexvn/tickex.git $GOPATH/src/github.com/tickexvn/tickex
    ```

**Navigate to the project directory**
```sh
cd $GOPATH/src/github.com/tickexvn/tickex
```

## Running the Service

To build and run the service using `make`, use the following commands:

- **Build the Tickex Edge (included API Gateway):**
  ```sh
  make run.tickex
  ```

- **Run the service:**
  ```sh
  make run.x.<service>
  ```

## License

Copyright (c) Tickex Labs. All rights reserved.

Licensed under the [Apache 2.0](LICENSE) license.

#

Made in 🇻🇳 🚀