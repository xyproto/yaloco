# YaLoCo

Yet another log colorizer

## Usage

Colorize build information while building:

    make | yaloco

Colorize logs:

    yaloco < /var/log/pacman.log

## Installation

### Linux

    go build
    sudo install -Dm755 yaloco /usr/bin/yaloco

## General info

* Version: 1.2.0
* License: MIT
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
