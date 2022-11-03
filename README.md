# mynona.

*NOTE: This project is still under very early development.*

The mynona project aims to provide a simple tool to create throwaway mail addresses using the [Migadu](https://www.migadu.com/api) API. It provides an alternative to services like [AnonAddy](https://anonaddy.com) or [SimpleLogin](https://simplelogin.io).

The differences are:

* Mynona itself is free BUT it requires a [aid Migadu subscription](https://www.migadu.com/pricing)
* No need to host anything yourself, since it uses Migadu's servers

## Requirements

To use this tool a Migadu API token is required with at least one registered domain.

## Usage

### Configuration

1. Copy `config.json.template` to `config.json`
2. Fill in your Migadu credentials (user and API token)
3. Fill in the domains registered in your migadu account you would like to use for throwaway mail addresses.

### Setup

1. Build the Vue frontend: `( cd frontend/ && yarn install && yarn build )`
2. Build the Go backend: `go build`
3. Start the service: `./mynona`
4. The UI is now available at http://localhost:3000

## Feature Status

- [x] Get index of: mailboxes, identities and aliases
- [ ] Create new identity 
- [ ] Delete identity
- [ ] Update identity
- [ ] User Management
- [ ] Authentication