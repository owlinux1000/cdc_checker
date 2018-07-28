# CAFÉ de CRIÉ Point Card Checker

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE.txt)

## What's this

You can easily confirm your CAFÉ de CRIÉ point card information such as balance, expiration date !

## Usage

```
$ git clone https://github.com/owlinux1000/cdc_checker
$ cd cdc_checker
$ go build
$ emacs .env
CARD_ID='XXXXXXXXX'
PIN_ID='YYYYYY'
$ ./cdc_checker
{"card_id":"XXXXXXXXXX","balance":"￥551","expiration_date":"2020/06/19"}
```
