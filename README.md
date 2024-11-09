# About
Currency Wallet is the part of the complex of microsirvices provides API to manage wallets and exchange currencies
Exchanger could be find here https://github.com/tyagnii/gw-exchanger/tree/entire_project

# Installation
```go
go get "github.com/tyagnii/gw-currency-wallet.git"
```

# Server
## build
Use makefile to build
`make build`

**! REMEMBER**
        
    The build is optimized only for Linux systems. Use your own build parameter for different OS

## run
`server serve`

## config
Configuration is placed in config.env file

# Docker-Compose Deployment
To deploy entire project with exchanger and database there is a docker-compose config file 
https://github.com/tyagnii/gw-cicd