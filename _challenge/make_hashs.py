#!/usr/bin/env python3
import sys
import json
import hashlib

if len(sys.argv) < 2:
    print("Please provde keys.json file")
    sys.exit(1)

data, hashs = {}, {}

with open(sys.argv[1], "r") as f:
    data = json.load(f)

for k, v in data.items():
    hashs[k] = hashlib.sha256(v.encode()).hexdigest()

with open("hashs.json", "w") as f:
    json.dump(hashs, f, indent=4, sort_keys=True)
