#!bin/python

import argparse
import yaml
import requests
import os
from multiprocessing import Pool

global parser_url

parser = argparse.ArgumentParser(description='Parser\'s test')
parser.add_argument('--parser-url', default='http://127.0.0.1:9080/parse', help='http://127.0.0.1:9080/parse')
parser.add_argument('--test-file', default='go/src/rent-parser/tests/tests.yml',
                    help='go/src/rent-parser/tests/tests.yml')
parser.add_argument('--process-count', default=4, help='4', type=int)

args = parser.parse_args()

parser_url = args.parser_url


class Result:
    def __init__(self):
        self.total = 0
        self.valid_type = 0
        self.valid_price = 0


def split_list(alist, wanted_parts=4):
    length = len(alist)
    return [ alist[i*length // wanted_parts: (i+1)*length // wanted_parts] for i in range(wanted_parts) ]


def test_function(tests):
    result = Result()
    pid = os.getpid()
    total_tests = len(tests)
    for test in tests:
        result.total += 1
        r = requests.post(parser_url, data=test['text'].encode('utf-8'))
        json_response = r.json()

        if json_response['type'] == test['type']:
            result.valid_type += 1

        if json_response['price'] in test['price']:
            result.valid_price += 1

        if result.total % 10 == 0:
            print('[process %d] %d out of %d tests completed' % (pid, result.total, total_tests))

    return result


tests = yaml.safe_load(open(args.test_file, 'r'))

pool = Pool(processes=args.process_count)
results = pool.map(test_function, split_list(tests, wanted_parts=args.process_count))

total = 0
valid_type = 0
valid_price = 0
for result in results:
    total += result.total
    valid_type += result.valid_type
    valid_price += result.valid_price

print('=== RESULT ===')
print('Total tests: %d' % total)
print('Valid types: %d (%.2f%%)' % (valid_type, valid_type*100/total))
print('Valid prices: %d (%.2f%%)' % (valid_price, valid_price*100/total))