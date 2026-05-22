# go-ciscosecureaccess

Go bindings for Cisco Secure Access (SSE) APIs

## Project Documentation

The `CODE_OF_CONDUCT.md` reflects our standards for interaction. 

The `CONTRIBUTING.md` file instructs new contributors on how to communicate with the project maintainers, report issues, provide pull requests, reviewing contributions, and how to version control releases.

The `LICENSE` file should contain the license you intend for the source code in the repo. 

The `SECURITY.md` file describes security policies and procedures including reporting a security-related bug and the policy on disclosure. 

## Package Documentation

[client](./client/docs/Client.md) - SSE OAuth2 base client

[ntg](./ntg/docs/NetworkTunnelGroupsAPI.md) - SSE Network Tunnel Groups API

[privateapps](./privateapps/docs/PrivateResourcesAPI.md) - SSE Private Resources API

[resconn](./resconn/docs/ConnectorGroupsAPI.md) - SSE Resource Connector Groups API

[rules](./rules/docs/AccessPoliciesAPI.md) - SSE Access Policies API

[virtualappliances](./virtualappliances/README.md) - SSE Virtual Appliances API

[identities](./identities/README.md) - SSE Identities Registration API

## Building

Build an API client by invoking it's `make` target

```
make generate-${API}
```

For example:

```
make generate-ntg
```

This will generate a package in a directory named after the API based off of the associated OpenAPI spec (hardcoded into the Makefile)

## Extending

1. Add support for additional SSE APIs by adding their spec file to `./specs`.
1. Create a `make` target in `Makefile` and generate the API using that target. 
1. Add a handle to get the client using OAuth2 in `./client/client.py`

