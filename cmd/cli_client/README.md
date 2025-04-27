## Console client

Clone the repository:

```bash
git clone https://github.com/BlackestDawn/nar-feed.git
```

Building and running the console client.

```bash
go build cmd/cli_client
./cli_client <flags>
```

Available flags:

* `-s` | `--section` - Comma separated list of sections to pull from. Defaults to "newest"
* `-t` | `-tags` - Comma separated list of tags to pull from. Using this will disregard sections.
* `-p` | `--pages` - Number of pages (per section) to pull, NAR has 5 posts per page. Defaults to 4
* `-a` | `--paginate` - Print only one post at a time
* `-h` | `--help` - Print the help screen
