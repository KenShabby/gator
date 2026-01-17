# Gator - RSS Aggregator

A multi-user RSS feed aggregator written in Go. A guided project from [Boot.dev](https://www.boot.dev)

## Requirements

To use gator, you'll need:

1. Go [Download and Install Go](https://go.dev/doc/install)
2. Postgres [Download](https://www.postgresql.org/download/)

## Installation

Once you have Go installed:

```
git clone https://github.com/KenShabby/gator.git
cd gator
go build -o gator
```

## Commands

- register [username] -- add a new user
- login [username] -- make a user the current user
- users -- list users
- agg [time]
- addfeed [nickname] [rss url]
- feeds -- list current feeds
- follow [rss url] -- follow a feed added by another user
- following -- print the feeds that the current user is following
- unfollow [rss url] -- stop following a feed for the current user
- browse [limit] -- print [limit] number of posts to the terminal. Default 2.
- reset -- reset all user/feed data
