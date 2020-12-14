#!/bin/bash

openssl req -x509 -config openssl.cnf -nodes -days 365 \
    -key private/privkey.pem -out private/fullchain.pem