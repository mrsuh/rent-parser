# Rent parser

[![Build Status](https://travis-ci.org/mrsuh/rent-parser.svg?branch=master)](https://travis-ci.org/mrsuh/rent-parser)

## Installation
```
sh bin/deploy
set parameters in a config/config.yml
set parameter 'Dictionary' in a tomita/{area, contact, price, type}/config.proto
```

## Compilation
```
sh bin/compile
```

## Run
```
bin/main /config/config.yml
```

## Use
```
-X POST -d 'сдаю двушку 50.4 кв.м за 30 тыс в месяц. телефон + 7 999 999 9999' 'http://localhost:/parse'
{"type":2,"phone":["9999999999"],"area":50.4,"price":30000}
```

## Types
```
0 - комната
1 - 1 комнатная квартира
2 - 2 комнатная квартира
3 - 3 комнатная квартира
4 - 4+ комнатная квартира
```