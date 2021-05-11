# socks (binary releases of https://github.com/fangdingjun/socks-go server)

Socks server.

## Usage

```
> socks
```

or

```
> socks -p 1081
```

## Installation

### Precompiled executables

Download the executable for the desired platform from the [release page](https://github.com/kpym/spcks/releases).

### Compile it yourself

#### Using Go

```
$ go get github.com/kpym/spcks
```

#### Using goreleaser

After cloning this repo you can compile the sources with [goreleaser](https://github.com/goreleaser/goreleaser/) for all available platforms:

```
git clone https://github.com/kpym/spcks.git .
goreleaser --snapshot --skip-publish --rm-dist
```

You will find the resulting binaries in the `dist/` sub-folder.

## License

[MIT](LICENSE) for this code (all used libraries may have different licence).
