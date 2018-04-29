# go-archive

[![CircleCI](https://circleci.com/gh/A-Hilaly/go-archive/tree/master.svg?style=svg&circle-token=e2065cf69b74bbb9357229b8cab69fe30ef6e25a)](https://circleci.com/gh/A-Hilaly/go-archive/tree/master)

Go tool for creating and reading archive files.

## Install

Get with go command

```
go get -u github.com/a-hilaly/go-archive/cmd/archiver
```

## Usage

```
Usage: archiver <flags> <subcommand> <subcommand args>

Subcommands:
	commands         list all command names
	compress         compress a file
	decompress       decompress a file
	flags            describe all known top-level flags
	help             describe subcommands and their syntax


Use "compress flags" for a list of top-level flags
exit status 2
```
## Subcommands

###### Decompress

```
archiver decompress [-t archive type] [-d destination] somefile:
    decompress somefile to wanted kind (zip by default).
    -d string
    	destination
    -t string
    	archive type (default "zip")
```

###### Compress

```
archiver compress [-d destination] [-k archive type] somefile:
    compress somefile to wanted kind (zip by default).
    -t string
    	archive type (default "zip")
```


## TODO

- ~~archiver commmand~~
- ~~zip support~~
- tar support
- gzib support
