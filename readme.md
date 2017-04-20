# `swptop` [![Build Status](https://travis-ci.org/essentialkaos/swptop.svg?branch=master)](https://travis-ci.org/essentialkaos/swptop) [![Go Report Card](https://goreportcard.com/badge/github.com/essentialkaos/swptop)](https://goreportcard.com/report/github.com/essentialkaos/swptop) [![License](https://gh.kaos.io/ekol.svg)](https://essentialkaos.com/ekol)

`swptop` is simple utility for viewing swap consumption of processes.

## Installation

### From source

Before the initial install allows git to use redirects for [pkg.re](https://github.com/essentialkaos/pkgre) service (reason why you should do this described [here](https://github.com/essentialkaos/pkgre#git-support)):

```
git config --global http.https://pkg.re.followRedirects true
```

To build the MDToc from scratch, make sure you have a working Go 1.6+ workspace ([instructions](https://golang.org/doc/install)), then:

```
go get github.com/essentialkaos/swptop
```

If you want to update MDToc to latest stable release, do:

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

### Prebuilt binaries

You can download prebuilt binaries for Linux and OS X from [EK Apps Repository](https://apps.kaos.io/swptop/latest).

## Usage

```
Usage: swptop {options}

Options

  --no-color, -nc    Disable colors in output
  --help, -h         Show this help message
  --version, -v      Show version

```

## Build Status

| Branch | Status |
|------------|--------|
| `master` | [![Build Status](https://travis-ci.org/essentialkaos/swptop.svg?branch=master)](https://travis-ci.org/essentialkaos/swptop) |
| `develop` | [![Build Status](https://travis-ci.org/essentialkaos/swptop.svg?branch=develop)](https://travis-ci.org/essentialkaos/swptop) |

## Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

## License

[EKOL](https://essentialkaos.com/ekol)
