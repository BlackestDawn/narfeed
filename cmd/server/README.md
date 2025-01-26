## Server

Clone the repository:

```bash
git clone https://github.com/BlackestDawn/nar-feed.git
```

Building and running the console client.

```bash
go build cmd/server
./client [mode] <flags>
```

Available modes are:

* `atom` - Will run a simple webserver and serve an atom-feed on the listening port
* `console` - Will run persitently and print directly to the console

Available flags:

* `-s` | `--section` - Comma separated list of sections to pull from. Defaults to "newest"
* `-p` | `--pages` - Number of pages (per section) to pull, NAR has 5 posts per page. Defaults to 4
* `-l` | `--listenport` - Port that the server will listen on, only used on atom mode. Defaults to 8181
* `-a` | `--paginate` - Print only one post at a time, only used in console mode
* `-h` | `--help` - Print the help screen
