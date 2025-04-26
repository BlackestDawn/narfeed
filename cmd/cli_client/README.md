## Console client

Clone the repository:

```bash
git clone https://github.com/BlackestDawn/nar-feed.git
```

Building and running the console client.

```bash
go build cmd/client
./client <flags>
```

Available flags:

* `-s` | `--section` - Comma separated list of sections to pull from. Defaults to "newest"
* `-p` | `--pages` - Number of pages (per section) to pull, NAR has 5 posts per page. Defaults to 4
* `-a` | `--paginate` - Print only one post at a time
* `-h` | `--help` - Print the help screen
