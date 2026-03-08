# Gator

Gator is a terminal RSS feed aggregator. It lets you add RSS feeds, follow/unfollow them across multiple users, and browse posts -- all from the command line. A background aggregation loop continuously fetches new posts and stores them in a PostgreSQL database.

## Prerequisites

- [Go](https://go.dev/dl/) (1.22 or later)
- [PostgreSQL](https://www.postgresql.org/download/) (15 or later recommended)

Make sure Postgres is running and you have a database created for Gator:

```sql
CREATE DATABASE gator;
```

## Installation

```bash
go install github.com/<your-username>/gator@latest
```

This compiles and installs the `gator` binary to your `$GOPATH/bin`. Make sure that directory is in your `PATH`.

## Configuration

Gator reads its configuration from `~/.gatorconfig.json`. Create the file with the following structure:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable"
}
```

Replace `username` and `password` with your Postgres credentials.

Database migrations run automatically on startup -- no manual migration step is needed.

## Commands

### User management

| Command | Description |
|---|---|
| `gator register <name>` | Create a new user and set it as the current user |
| `gator login <name>` | Switch to an existing user |
| `gator users` | List all users (current user is marked) |

### Feed management

| Command | Description |
|---|---|
| `gator addfeed <name> <url>` | Add an RSS feed and automatically follow it |
| `gator feeds` | List all feeds in the database |

### Following

| Command | Description |
|---|---|
| `gator follow <url>` | Follow an existing feed by URL |
| `gator unfollow <url>` | Unfollow a feed by URL |
| `gator following` | List feeds the current user is following |

### Aggregation

| Command | Description |
|---|---|
| `gator agg <duration>` | Start the aggregation loop (e.g. `gator agg 1m`) |

The `agg` command runs continuously, fetching one feed per tick. Leave it running in a background terminal while you interact with the CLI in another.

### Browsing posts

| Command | Description |
|---|---|
| `gator browse [limit]` | View recent posts from followed feeds (default: 2) |

### Admin

| Command | Description |
|---|---|
| `gator reset` | Delete all users (cascades to feeds, follows, and posts) |

## Example workflow

```bash
# Register and add some feeds
gator register alice
gator addfeed "Hacker News" "https://news.ycombinator.com/rss"
gator addfeed "Boot.dev Blog" "https://www.boot.dev/blog/index.xml"

# Start the aggregator in a separate terminal
gator agg 1m

# Browse the latest posts
gator browse 5
```
