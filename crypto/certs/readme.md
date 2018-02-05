# gen certs

## rsa

```shell
# CA
openssl genrsa -out ca.key 2048
# Self-Signed Certificate
openssl req -new -x509 -key ca.key -out ca.crt -days 3600

# create server key and use ca's certificate sign server's certificate
openssl genrsa -out server.key 2048
openssl req -new -sha256 -key server.key -out server.csr -days 3600
openssl x509 -req -CA ca.crt -CAkey ca.key -CAcreateserial -in server.csr -out server.crt -extensions v3_req -extfile server.cnf

```

## ecc

```shell
# Generating an ECDSA Key
openssl ecparam -name prime256v1 -genkey -out ca.key
# Self-Signed Certificate
openssl req -new -x509 -key ca.key -out ca.crt -days 3600

# create server key and use ca's certificate sign server's certificate
openssl ecparam -name prime256v1 -genkey -out server.key
openssl req -new -sha256 -key server.key -out server.csr
openssl x509 -req -CA ca.crt -CAkey ca.key -CAcreateserial -in server.csr -out server.crt -extensions v3_req -extfile server.cnf

```
