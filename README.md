# Rent parser

## Quick start with Docker
```bash
docker run -p 9080:9080 -d mrsuh/rent-parser
curl -X POST -d 'сдаю двушку за 30 тыс в месяц' 'http://127.0.0.1:9080/parse'
{"type":2,"price":30000}
```

### Types
* 0 - комната
* 1 - 1 комнатная квартира
* 2 - 2 комнатная квартира
* 3 - 3 комнатная квартира
* 4 - 4+ комнатная квартира


## Compilation from source code

### Generate configs and load [Yandex Tomita parser](https://github.com/yandex/tomita-parser)
```bash
sh bin/deploy.sh
```

### Configure
* config/config.yml
* tomita/{price, type}/config.proto

### Compile
```bash
sh bin/compile.sh
```

### Run
```bash
bin/server /config/config.yml
```

## Tests
```bash
time python tests/test.py --parser-url 'http://127.0.0.1:9080' --test-file 'go/src/rent-parser/tests/tests.yml' --process-count 4
=== RESULT ===
Total tests: 8000
Valid types: 7475 (93.44%)
Valid prices: 5302 (66.28%)

real    35m23.099s
user    0m42.048s
sys     0m2.529s
```