# `swptop` [![Build Status](https://travis-ci.org/essentialkaos/swptop.svg?branch=master)](https://travis-ci.org/essentialkaos/swptop) [![Go Report Card](https://goreportcard.com/badge/github.com/essentialkaos/swptop)](https://goreportcard.com/report/github.com/essentialkaos/swptop) [![codebeat badge](https://codebeat.co/badges/21eb1670-e54a-4373-8f4b-cfb861198d4c)](https://codebeat.co/projects/github-com-essentialkaos-swptop-master) [![License](https://gh.kaos.io/ekol.svg)](https://essentialkaos.com/ekol)

`swptop` is simple utility for viewing swap consumption of processes.

### Installation

#### From source

Before the initial install allows git to use redirects for [pkg.re](https://github.com/essentialkaos/pkgre) service (reason why you should do this described [here](https://github.com/essentialkaos/pkgre#git-support)):

```
git config --global http.https://pkg.re.followRedirects true
```

To build the `swptop` from scratch, make sure you have a working Go 1.6+ workspace ([instructions](https://golang.org/doc/install)), then:

```
go get github.com/essentialkaos/swptop
```

If you want to update `swptop` to latest stable release, do:

```
go get -u github.com/essentialkaos/swptop
```

#### From ESSENTIAL KAOS Public repo for RHEL6/CentOS6

```
[sudo] yum install -y https://yum.kaos.io/6/release/x86_64/kaos-repo-8.0-0.el6.noarch.rpm
[sudo] yum install swptop
```

#### From ESSENTIAL KAOS Public repo for RHEL7/CentOS7

```
[sudo] yum install -y https://yum.kaos.io/7/release/x86_64/kaos-repo-8.0-0.el7.noarch.rpm
[sudo] yum install swptop
```

### Usage

```
Usage: swptop {options}

Options

  --user, -u         Filter output by user
  --filter, -f       Filter output by part of command
  --no-color, -nc    Disable colors in output
  --help, -h         Show this help message
  --version, -v      Show version

Examples

  swptop
  Show current swap consumption of all processes

  swptop -u redis
  Show current swap consumption by webserver user processes

  swptop -f redis-server
  Show current swap consumption by processes with 'redis-server' in command

  swptop | wc -l
  Count number of processes which use swap

```

### Build Status

| Branch | Status |
|--------|--------|
| `master` | [![Build Status](https://travis-ci.org/essentialkaos/swptop.svg?branch=master)](https://travis-ci.org/essentialkaos/swptop) |
| `develop` | [![Build Status](https://travis-ci.org/essentialkaos/swptop.svg?branch=develop)](https://travis-ci.org/essentialkaos/swptop) |

### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

### License

[EKOL](https://essentialkaos.com/ekol)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.io/ekgh.svg"/></a></p>
