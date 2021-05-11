# socks

Socks server, binary releases of https://github.com/fangdingjun/socks-go example server.

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

Download the executable for the desired platform from the [release page](https://github.com/kpym/socks/releases).

### Compile it yourself

#### Using Go

```
$ go get github.com/kpym/socks
```

#### Using goreleaser

After cloning this repo you can compile the sources with [goreleaser](https://github.com/goreleaser/goreleaser/) for all available platforms:

```
git clone https://github.com/kpym/socks.git .
goreleaser --snapshot --skip-publish --rm-dist
```

You will find the resulting binaries in the `dist/` sub-folder.

## My use case

I want to use socks5 over ssh but the dynamic port forward (`-D` of `ssh`) is disabled on a server.

So the plan is:
```
> ssh  ssh -L 1080:localhost:1080 user@server
> wget -c https://github.com/kpym/socks/releases/download/v0.1.0/socks_0.1.0_Linux_64bit.tar.gz -O - | tar -xz socks
> ./socks
```
and then use `localhost:1080` as socks5 proxy.

## License

[MIT](LICENSE) for this code _(but all used libraries may have different licence)_.
