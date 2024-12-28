# NotAlwaysRight feed generator/displayer

Minimalistic program to read posts on notalwaysright.com

## Installation

Clone the repository:

`git clone https://github.com/BlackestDawn/nar-feed.git`

#### Console versions

Build and run the client.

First got the repository and issue:

`go build cmd/client`

Run it by issuing:

`./client`

## TODO

* Add pagination for console output
* Add ability to choose/filter sections
* Add server version that will continously and peridocally update
* Add support for running as GCP Gloud Function and AWS Lambda
* Add Ansible and Terraform routines for setting up GCP and AWS