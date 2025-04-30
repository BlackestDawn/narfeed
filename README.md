# NotAlwaysRight feed generator/displayer

Minimalistic program to fetch posts from Notalwaysright.com and either display them directly in a console terminal or create an atom-feed from them. Feed generation can be done by running it as a local server or run as a server-less function in a public cloud.

## Motivation

Was tired of not having an easy way to keep track of which posts I had already read on Not Always Right (NAR) so wanted to create a way for me to do so. I also wanted to apply all that I had learned from Boot.dev and make it as adaptable as possible.

## Installation and Usage

See the individual pages for the respective variants:

[CLI client](https://github.com/BlackestDawn/nar-feed/tree/main/cmd/client/) | [Server](https://github.com/BlackestDawn/nar-feed/tree/main/cmd/server/) | [GCP Cloud Run function](https://github.com/BlackestDawn/nar-feed/tree/main/cmd/gcp/) | [AWS Lambda](https://github.com/BlackestDawn/nar-feed/tree/main/cmd/aws/)

## Contributing

### Clone the repo

```bash
git clone https://github.com/BlackestDawn/nar-feed
cd nar-feed
```

### Build and run

You can use the client version to test the general logic, and the server version for atom-feed generation.

```bash
go build ./cmd/[client, server]
./[cli_client, server]
```

or

```bash
go run ./cmd/[cli_client, server]
```

### Submit a pull request

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.

## TODO

* ~~Add pagination for console output~~
* ~~Actually use CLI parameters~~
* ~~Ability to choose/filter sections~~
* Some form of "memory" so it only displays new posts since last run
* ~~Server version that will continously and peridocally update~~
* Support for running as GCP Gloud Function and AWS Lambda
* Ansible and Terraform routines for setting up GCP and AWS