# Rent parser

[![Build Status](https://travis-ci.org/mrsuh/rent-parser.svg?branch=master)](https://travis-ci.org/mrsuh/rent-parser)

## Installation
```
sh bin/deploy.sh
set parameters in a config/config.yml
set parameters in a tomita/{area, contact, price, type}/config.proto
```

## Compilation
```sh
sh bin/compile.sh
```

## Run
```sh
bin/main /config/config.yml
```

## Use
```sh
curl -X POST -d 'сдаю двушку 50.4 кв.м за 30 тыс в месяц. телефон + 7 999 999 9999' 'http://localhost:/parse'
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

## Configuration
config/config.yml
```yml
server.host: '127.0.0.1'
server.port: '9080'
tomita.bin: '/path/to/bin/tomita'
tomita.conf.type: '/path/to/tomita/type/config.proto'
tomita.conf.contact: '/path/to/tomita/contact/config.proto'
tomita.conf.area: '/path/to/tomita/area/config.proto'
tomita.conf.price: '/path/to/tomita/price/config.proto'
```

tomita/{area, contact, price, type}/config.proto
```proto
encoding "utf8";

TTextMinerConfig {

    Dictionary = "/path/to/tomita/area/dict.gzt";

    Output = {
        Format = xml;
    }

    Articles = [
       { Name = "article" }
    ]

    Facts = [
       { Name = "Fact" }
    ]
}
```