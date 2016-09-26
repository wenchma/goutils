# checksum - Calculate the MD5 and SHA1 hash

## Installation

Download the `checksum.go` code, then build the binary:

```
# go build -o /usr/local/bin/checksum checksum.go
```

## Usage

```
$ checksum 
Usage: checksum [md5|sha1] <file name>

$ checksum md5 file
+-----------+----------------------------------+
| Algorithm | Value                            |
+-----------+----------------------------------+
| sha1      | 41c99fbe91ab1aba5a1fbe3e5563d67d |
+-----------+----------------------------------+

$ checksum sha1 file
+-----------+------------------------------------------+
| Algorithm | Value                                    |
+-----------+------------------------------------------+
| sha1      | 73325962b630aae6c4742f5e1976dc3078a2d76f |
+-----------+------------------------------------------+
```

