## Yet another log colorizer

Bring some color to your build output and logs, to improve clarity and readability.

## Usage

When building with `make`:

    make | yaloco

Viewing `pacman` logs:

    yaloco < /var/log/pacman.log

## Requirements

* Go 1.14 or later

## Installation

### With Go 1.17 or later

    go install github.com/xyproto/yaloco@latest

### From source, on Linux

    go build -mod=vendor
    sudo install -Dm755 yaloco /usr/bin/yaloco

### Using the static executable, on 64-bit x86 Linux

Just download the statically compiled executable from the "Releases" page, and use that.

## General info

* Version: 1.3.2
* License: BSD-3
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
