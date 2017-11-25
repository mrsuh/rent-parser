# Rent parser

[![Build Status](https://travis-ci.org/mrsuh/rent-parser.svg?branch=master)](https://travis-ci.org/mrsuh/rent-parser)

## Installation
```
sh bin/deploy.sh
set parameters in a config/config.yml
set parameters in a tomita/{price, type}/config.proto
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
curl -X POST -d 'сдаю двушку за 30 тыс в месяц' 'http://localhost:/parse'
{"type":2,"price":30000}
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
tomita.conf.price: '/path/to/tomita/price/config.proto'
```

tomita/{area, contact, price, type}/config.proto
```proto
encoding "utf8";

TTextMinerConfig {

    Dictionary = "/path/to/tomita/type/dict.gzt";

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