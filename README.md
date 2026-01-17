Explain to the user how to set up the config file and run the program. Tell them about a few of the commands they can run.
Push gator up to GitHub, then submit the link to your remote repo. Your link should look something like this: https://github.com/github-username/repo-name.
Remember: Go programs are statically compiled binaries. After running go build or go install, you should be able to run your code without needing the go toolchain.


# Gator - RSS Aggregator

A multi-user RSS feed aggregator written in Go. A guided project from [Boot.dev](https://www.boot.dev)

## Requirements

To use gator, you'll need:

1. Go [Download and Install Go](https://go.dev/doc/install)
2. Postgres [Download](https://www.postgresql.org/download/)

## Installation

Once you have Go installed:

go install github.com/KenShabby/gator

or build from source:

git clone https://github.com/KenShabby/gator.git
cd gator
go build -o gator

## Commands


register <username> -- add a new user
login <username> -- make a user the current user
users -- list users
agg [time]
addfeed <nickname> <rss url>
feeds -- list current feeds
follow <rss url> -- follow a feed added by another user
following -- print the feeds that the current user is following
unfollow <rss url> -- stop following a feed for the current user
browse <limit> -- print <limit> number of posts to the terminal. Default 2.
reset -- reset all user/feed data

