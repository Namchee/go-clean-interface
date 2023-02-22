# go-clean-interface

Checks interface declarations for unnecessary `I` prefixes or `Itf` suffixes.

## Why?

[Main Reasoning](https://stackoverflow.com/a/5817904)

Prefixing interfaces with `I` or suffixing them with `Itf` is leftover convention from Java and C#. When we are writing code with clean architecture as a guideline, whatever we pass as a dependency should be expected as an interface.

Interfaces shouldn't be treated as a special language feature that requires special naming conventions.

## Installation

```shell
go install github.com/Namchee/go-clean-interface@latest
```

## Usage

```shell
go-clean-interface [-flag] [package]
```

## Options

All options are provided through CLI flags.

Name | Type | Default | Description
`no-prefix` | `bool` | `true` | Disallow interface declarations to be prefixed with `I`
`no-suffix` | `bool` | `true` | Disallow interface declarations to be suffixed with `Itf`
