<p align="center"><a href="#readme"><img src=".github/images/card.svg"/></a></p>

<p align="center">
  <a href="https://kaos.sh/r/swptop"><img src="https://goreportcard.com/badge/github.com/essentialkaos/swptop" alt="GoReportCard" /></a>
  <a href="https://kaos.sh/y/swptop"><img src="https://app.codacy.com/project/badge/Grade/1c97a07461d34092b5fded246af36f59" alt="Codacy badge" /></a>
  <a href="https://kaos.sh/w/swptop/ci"><img src="https://github.com/essentialkaos/swptop/actions/workflows/ci.yml/badge.svg" alt="GitHub Actions CI Status" /></a>
  <a href="https://kaos.sh/w/swptop/codeql"><img src="https://github.com/essentialkaos/swptop/actions/workflows/codeql.yml/badge.svg" alt="GitHub Actions CodeQL Status" /></a>
  <a href="#license"><img src=".github/images/license.svg"/></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#usage">Usage</a> • <a href="#ci-status">CI Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

<br/>

`swptop` is simple utility for viewing swap consumption of processes.

### Installation

#### From source

To build the `swptop` from scratch, make sure you have a working Go 1.25+ workspace ([instructions](https://go.dev/doc/install)), then:

```
go install github.com/essentialkaos/swptop@latest
```

#### From [ESSENTIAL KAOS Public Repository](https://kaos.sh/kaos-repo)

```bash
sudo dnf install -y https://pkgs.kaos.st/kaos-repo-latest.el$(grep 'CPE_NAME' /etc/os-release | tr -d '"' | cut -d':' -f5).noarch.rpm
sudo dnf install swptop
```

#### Prebuilt binaries

You can download prebuilt binaries for Linux from [EK Apps Repository](https://apps.kaos.st/swptop/latest).

To install the latest prebuilt version, do:

```bash
bash <(curl -fsSL https://apps.kaos.st/get) swptop
```

### Usage

<img src=".github/images/usage.svg" />

### CI Status

| Branch | Status |
|--------|--------|
| `master` | [![CI](https://github.com/essentialkaos/swptop/actions/workflows/ci.yml/badge.svg?branch=master)](https://kaos.sh/w/swptop/ci?query=branch:master) |
| `develop` | [![CI](https://github.com/essentialkaos/swptop/actions/workflows/ci.yml/badge.svg?branch=develop)](https://kaos.sh/w/swptop/ci?query=branch:develop) |

### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/.github/blob/master/CONTRIBUTING.md).

### License

[Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://kaos.dev"><img src="https://raw.githubusercontent.com/essentialkaos/.github/refs/heads/master/images/ekgh.svg"/></a></p>
