# Log colorizer

Bring some color to your build output and logs, to improve clarity and readability.

## Usage

When building with `make`:

    make | yaloco

Viewing `pacman` logs:

    yaloco < /var/log/pacman.log

## Requirements

* Go 1.14 or later

## Installation

### Linux

    go build -mod=vendor
    sudo install -Dm755 yaloco /usr/bin/yaloco

## General info

* Version: 1.3.0
* License: MIT
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
