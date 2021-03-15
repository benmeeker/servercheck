# ServerCheck

## Status

Early stage of development - currently only gets OS, Kernel, Hostname, Packages, and respective versions - needs a lot more

Note that we will try our best to avoid any backward incompatibility.  If we do run into something, it will be documented.

## About

ServerCheck is an app that can give you updated information on a database on servers, specifically what needs to be updated. This is a learning project - so quality will be questionable at best

## Configuration

No configuration is currently possible.  See [Todo](#todo)

```
TBD
```

## Installation

Ensure you have Go >= 1.16.0 installed and set up on your machine, then run the following command:

```
$ go install github.com/benmeeker/servercheck
```

## Usage

Usage is very straightforward

**CAUTION**: Servercheck currently only supports http NOT https

## Todo
- [ ] Fix navbar styling
- [ ] Make search bar work
- [ ] Create page updates without refreshing
- All Servers page
  - [ ] Containers for each server
  - [ ] Basic information in containers
- Outdated Servers page
  - [ ] Containers for each server
  - [ ] Basic information in containers
And so much more.... (please submit a feature request issue)

## License

MIT
