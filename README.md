<p align="center"><a href="#readme"><img src="https://gh.kaos.st/swptop.svg"/></a></p>

<p align="center">
  <a href="https://kaos.sh/w/swptop/ci"><img src="https://kaos.sh/w/swptop/ci.svg" alt="GitHub Actions CI Status" /></a>
  <a href="https://kaos.sh/r/swptop"><img src="https://kaos.sh/r/swptop.svg" alt="GoReportCard" /></a>
  <a href="https://kaos.sh/b/swptop"><img src="https://kaos.sh/b/21eb1670-e54a-4373-8f4b-cfb861198d4c.svg" alt="codebeat badge" /></a>
  <a href="https://kaos.sh/w/swptop/codeql"><img src="https://kaos.sh/w/swptop/codeql.svg" alt="GitHub Actions CodeQL Status" /></a>
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
sudo yum install -y https://yum.kaos.st/kaos-repo-latest.el$(grep 'CPE_NAME' /etc/os-release | tr -d '"' | cut -d':' -f5).noarch.rpm
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
| `master` | [![CI](https://kaos.sh/w/swptop/ci.svg?branch=master)](https://kaos.sh/w/swptop/ci?query=branch:master) |
| `develop` | [![CI](https://kaos.sh/w/swptop/ci.svg?branch=master)](https://kaos.sh/w/swptop/ci?query=branch:develop) |

### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

### License

[Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
