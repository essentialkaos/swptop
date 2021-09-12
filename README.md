<p align="center"><a href="#readme"><img src="https://gh.kaos.st/swptop.svg"/></a></p>

<p align="center">
  <a href="https://github.com/essentialkaos/swptop/actions"><img src="https://github.com/essentialkaos/swptop/workflows/CI/badge.svg" alt="GitHub Actions Status" /></a>
  <a href="https://github.com/essentialkaos/swptop/actions?query=workflow%3ACodeQL"><img src="https://github.com/essentialkaos/swptop/workflows/CodeQL/badge.svg" /></a>
  <a href="https://goreportcard.com/report/github.com/essentialkaos/swptop"><img src="https://goreportcard.com/badge/github.com/essentialkaos/swptop"></a>
  <a href="https://codebeat.co/projects/github-com-essentialkaos-swptop-master"><img alt="codebeat badge" src="https://codebeat.co/badges/21eb1670-e54a-4373-8f4b-cfb861198d4c" /></a>
  <a href="#license"><img src="https://gh.kaos.st/apache2.svg"></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#usage">Usage</a> • <a href="#build-status">Build Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

<br/>

`swptop` is simple utility for viewing swap consumption of processes.

### Installation

#### From source

To build the `swptop` from scratch, make sure you have a working Go 1.16+ workspace ([instructions](https://golang.org/doc/install)), then:

```
go get github.com/essentialkaos/swptop
```

If you want to update `swptop` to latest stable release, do:

```
go get -u github.com/essentialkaos/swptop
```

#### From [ESSENTIAL KAOS Public Repository](https://yum.kaos.st)

```bash
sudo yum install -y https://yum.kaos.st/get/$(uname -r).rpm
sudo yum install swptop
```

#### Prebuilt binaries

You can download prebuilt binaries for Linux from [EK Apps Repository](https://apps.kaos.st/swptop/latest).

To install the latest prebuilt version, do:

```bash
bash <(curl -fsSL https://apps.kaos.st/get) swptop
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
| `master` | [![CI](https://github.com/essentialkaos/swptop/workflows/CI/badge.svg?branch=master)](https://github.com/essentialkaos/swptop/actions) |
| `develop` | [![CI](https://github.com/essentialkaos/swptop/workflows/CI/badge.svg?branch=develop)](https://github.com/essentialkaos/swptop/actions) |

### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

### License

[Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
