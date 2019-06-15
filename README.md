<p align="center"><a href="#readme"><img src="https://gh.kaos.st/swptop.svg"/></a></p>

<p align="center">
  <a href="https://travis-ci.org/essentialkaos/swptop"><img src="https://travis-ci.org/essentialkaos/swptop.svg"></a>
  <a href="https://goreportcard.com/report/github.com/essentialkaos/swptop"><img src="https://goreportcard.com/badge/github.com/essentialkaos/swptop"></a>
  <a href="https://codebeat.co/projects/github-com-essentialkaos-swptop-master"><img alt="codebeat badge" src="https://codebeat.co/badges/21eb1670-e54a-4373-8f4b-cfb861198d4c" /></a>
  <a href="https://essentialkaos.com/ekol"><img src="https://gh.kaos.st/ekol.svg"></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#usage">Usage</a> • <a href="#build-status">Build Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

<br/>

`swptop` is simple utility for viewing swap consumption of processes.

### Installation

#### From source

Before the initial install allows git to use redirects for [pkg.re](https://github.com/essentialkaos/pkgre) service (_reason why you should do this described [here](https://github.com/essentialkaos/pkgre#git-support)_):

```
git config --global http.https://pkg.re.followRedirects true
```

To build the `swptop` from scratch, make sure you have a working Go 1.10+ workspace ([instructions](https://golang.org/doc/install)), then:

```
go get github.com/essentialkaos/swptop
```

If you want to update `swptop` to latest stable release, do:

```
go get -u github.com/essentialkaos/swptop
```

#### From ESSENTIAL KAOS Public repo for RHEL6/CentOS6

```
[sudo] yum install -y https://yum.kaos.st/kaos-repo-latest.el6.noarch.rpm
[sudo] yum install swptop
```

#### From ESSENTIAL KAOS Public repo for RHEL7/CentOS7

```
[sudo] yum install -y https://yum.kaos.st/kaos-repo-latest.el7.noarch.rpm
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

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
