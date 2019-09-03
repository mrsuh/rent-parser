# Tests

```bash
python3 -m venv venv
. venv/bin/activate
pip install -r requirements.txt
python test.py --parser-url 'http://127.0.0.1:9080' --test-file 'go/src/rent-parser/tests/tests.yml' --process-count 4
```