[![Build Status](https://travis-ci.org/0intro/stellar-keypairs.svg?branch=master)](https://travis-ci.org/0intro/stellar-keypairs)

Stellar-KeyPairs
================

This tool generates Stellar key pairs.

Usage
-----

```
usage: stellar-keypairs [ -n nWorkers ] [ -p prefix | -s seed ]
```

Examples
--------

Generate a Stellar key pair from a random seed:

```
$ stellar-keypairs
Seed (secret key) SB45IAMBKTMCWK2DOESPZ3TVESQTW2Q24YZYCNGE6ENUA4LPY4BD6DMV
Public key GCMTXTBA2WJCIOWQUYKYV3ZENTMTC6A6FGDKRUMMVKUUPNEP34TJ6ZPK
```

Generate a Stellar key pair from the specified seed:

```
$ stellar-keypairs -seed SB45IAMBKTMCWK2DOESPZ3TVESQTW2Q24YZYCNGE6ENUA4LPY4BD6DMV
Seed (secret key) SB45IAMBKTMCWK2DOESPZ3TVESQTW2Q24YZYCNGE6ENUA4LPY4BD6DMV
Public key GCMTXTBA2WJCIOWQUYKYV3ZENTMTC6A6FGDKRUMMVKUUPNEP34TJ6ZPK
```

Generate a Stellar key pair with a public key beginning by the specified prefix (with 4 workers):

```
$ stellar-keypairs -prefix GBOB -n 4
Seed (secret key) SAN4F7ZOX7UNJX4GFHNXIGODARXPXPWCL3XRW4H2KKHW34PMJJRFD6BN
Public key GBOB5GANEX3ZZMOXKHWDYOJPU6D5S4HLE26C2BPYKBMBJ755H3AKURAN
```
