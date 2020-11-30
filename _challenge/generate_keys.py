#!/usr/bin/env python
import random
import string
import hashlib
import json

def get_random_string(length=32):
    letters = string.ascii_lowercase + string.ascii_uppercase + string.digits
    return ''.join(random.choice(letters) for i in range(length))

def make_hash(keys):
    data = ''
    for i in keys.values():
        data += i
    print(data)
    h = hashlib.sha256(data.encode())
    return h.hexdigest()

keys = {}    
for i in range(24):
    day = 'day' + str(i+1).zfill(2)
    keys[day] = get_random_string()

with open('keys.json', 'w') as f:
    json.dump(keys, f, indent=4, sort_keys=True)

h = make_hash(keys)

with open('password', 'w') as f:
    f.write(h)