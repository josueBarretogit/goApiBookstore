## Generate public and private key RSA

```bash
openssl genrsa -out demo.rsa
```

## Extract public key from privete key

```bash
openssl rsa -in demo.rsa -pubout > public.rsa.pub
```

