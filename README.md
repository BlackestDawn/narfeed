# NotAlwaysRight feed generator/displayer

Minimalistic program to fetch posts from Notalwaysright.com and either display them directly in a console terminal or create an atom-feed from them. Feed generation can be done by running it as a local server or run as a server-less function in a public cloud.

## Motivation

Was tired of not having an easy way to keep track of which posts I had already read on Not Always Right (NAR) so wanted to create a way for me to do so. I also wanted to apply all that I had learned from Boot.dev and make it as adaptable as possible.

## Installation

Clone the repository:

```bash
git clone https://github.com/BlackestDawn/nar-feed.git
```

### Console client

Building and running the console client.

```bash
go build cmd/client
./client
```

Available flags:

* `-s` - Comma separated list of sections to pull from. Defaults to "newest"
* `-p` - Number of pages (per section) to pull, NAR has 5 posts per page. Defaults to 4

### Server version

Comming

### Cloud version

Comming

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
./client
```

or

```bash
go run ./cmd/[client, server]
```

### Submit a pull request

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.

## TODO

* ~~Add pagination for console output~~
* Actually use CLI parameters
* Ability to choose/filter sections
* Some form of "memory" so it only displays new posts since last run
* Server version that will continously and peridocally update
* Support for running as GCP Gloud Function and AWS Lambda
* Ansible and Terraform routines for setting up GCP and AWS