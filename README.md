# olsu

olsu is a tool that makes it easy to create releases on github from cli. Use this if git tags aren't enough for you or you need to attach binaries to your release.

## Installation

There's binaries for Linux and macOS on the [releases](https://github.com/Telling/olsu/releases) page. 

## Requirements

You'll need something like `ca-certificates` on debian before olsu runs.

## Usage

```
name, version and description are required arguments
Usage of ./olsu:
   name version description [attachment ... attachment]
  -b string
    	backend (default "github")
  -d	if it's a draft.
  -dr
    	delete release if it exists before creating new.
  -o string
    	owner of the repository.
  -p	if it's a prerelease.
  -q	quiet, only output release id.
  -r string
    	name of the repository.
  -t string
    	access token for github.
```

A few flags can be provided as environment variables as well:

* `OLSU_OWNER` (-o)
* `OLSU_REPOSITORY` (-r)
* `OLSU_TOKEN` (-t)
* `OLSU_DELETE_RELEASE` (-dr)

## Todo

* Releases on bitbucket
* Description from file