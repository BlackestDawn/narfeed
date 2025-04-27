## Server

Clone the repository:

```bash
git clone https://github.com/BlackestDawn/nar-feed.git
```

Building and running the console client.

```bash
go build cmd/server
./server <flags>
```

Available flags:

* `-m` | `--mode` - Sets output mode. Currently supports 'atom' and 'console', defaults to 'atom'
* `-s` | `--section` - Comma separated list of sections to pull from. Defaults to 'newest'
* `-t` | `-tags` - Comma separated list of tags to pull from. Using this will disregard sections.
* `-p` | `--pages` - Number of pages (per section) to pull, NAR has 5 posts per page. Defaults to 4
* `-l` | `--listenport` - Port that the server will listen on, only used on atom mode. Defaults to 8181
* `-a` | `--paginate` - Print only one post at a time, only used in console mode
* `-h` | `--help` - Print the help screen
