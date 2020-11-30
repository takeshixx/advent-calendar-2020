#!/usr/bin/env python3
import hashlib

with open('tokens') as f:
    tokens = f.read()
    
h = hashlib.sha256()

for token in tokens.split('\n'):
    h.update(token.encode())

print('password: ', h.hexdigest())