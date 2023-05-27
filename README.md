# Go Authentication Service

The `go-authentication-service` package is a comprehensive authentication solution for Go applications. It provides various components and utilities to facilitate user authentication and session management. This README document serves as a guide to understand the structure and usage of the package.

## Table of Contents
- [Directory Structure](#directory-structure)
- [Installation](#installation)
- [Usage](#usage)
- [Components](#components)
  - [Auth Registry](#auth-registry)
  - [Cookie Manager](#cookie-manager)
  - [Session Manager](#session-manager)
  - [Utility](#utility)
- [License](#license)

## Directory Structure

The `go-authentication-service` package follows the following directory structure:

```
├── external
│   ├── cache_interface.go
│   └── encryptor_interface.go
└── service
    ├── authRegistry
    │   ├── authRegistry.go
    │   └── authRegistry_interface.go
    ├── cookieManager
    │   ├── cookieManager.go
    │   └── cookieManager_interface.go
    ├── sessionManager
    │   ├── option.go
    │   ├── session
    │   │   ├── session.go
    │   │   └── session_interface.go
    │   ├── sessionManager.go
    │   └── sessionManager_interface.go
    └── utility
        ├── cache.go
        ├── encryptor.go
        └── encryptorOption.go
```

## Installation

To install the `go-authentication-service` package, you can use the `go get` command:

```shell
go get github.com/Zhima-Mochi/go-authentication-service
```

## Components

The `go-authentication-service` package consists of the following components:

### Auth Registry

The `Auth Registry` component provides functionality to register and manage oauth2 authentication providers.
See [Auth Registry Interface](./service/authRegistry/authRegistry_interface.go) for more details.

### Cookie Manager

The `Cookie Manager` component handles the management of cookies related to user sessions.
See [Cookie Manager Interface](./service/cookieManager/cookieManager_interface.go) for more details.

### Session Manager

The `Session Manager` component is responsible for managing user sessions.
See [Session Manager Interface](./service/sessionManager/sessionManager_interface.go) for more details.

### Utility

The `Utility` package provides various utility functions used by the authentication service and also offers implementations of interfaces defined in the `external` package.
See [Utility Folder](./service/utility) for more details.

## Example

See [Example](./example/main.go) for a simple example of how to use the `go-authentication-service` package.

## License

The `go-authentication-service` package is distributed under the [LICENSE](./LICENSE) file. Please refer to the license for more details.