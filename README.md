# YaLoCo

Yet another log colorizer

## Usage

Colorize build information while building:

    make | yaloco

Colorize logs:

    yaloco < /var/log/pacman.log

## Requirements

Go 1.14 or 1.15 is recommended.

## Installation

### Linux

    go build -mod=vendor
    sudo install -Dm755 yaloco /usr/bin/yaloco

## General info

* Version: 2.0.0
* License: MIT
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
