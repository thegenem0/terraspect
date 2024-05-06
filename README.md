# TerraSpect

## Required Software

### Node

#### Linux

Install [https://github.com/nvm-sh/nvm](nvm) by following the instructions from the web site.
Install the latest lts version by running:

```bash
nvm ls-remote --lts
```

Note the version and running the following command to install it:

```bash
nvm install v18.18.0
```

### Docker

Install [https://docs.docker.com/engine/install/](Docker) by following the instructions from the web site.

### AWS CLI v2

### Go
This is required tooling to for the API's local development.
Install [https://golang.org/doc/install](Go) by following the instructions from the web site.

### Yarn

This is preferred over `npn / pnpm` as scripts depend on it.
Install [https://yarnpkg.com/getting-started/install](Yarn) by following the instructions from the web site.

## Configuration
Create a `.env` file in the web project's root, and add the following environment variables:

```bash
VITE_CLERK_PUBLISHABLE_KEY=<token>
VITE_API_BASE_URI=http://localhost:8080
```

Add API environment variables to the `docker-compose.yml` file in the api project section

## Development

Run `make dev` to start the development server.
`make dev` will start the api, and the web server, and will also bootstrap a local PostgreSQL database.

## Committing and Pushing code
Use conventional commits to commit code. This will allow for automatic versioning and changelog generation.

## Deployment

The system's infrastructure is configured in `/infra` and is managed through AWS CloudFormation.

The entire system can be deployed to a CloudFormation stack by running `make deploy`.
Access to AWS, and the correct permissions are required to deploy the stack.

