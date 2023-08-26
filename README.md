# B64

Learning about base64 by reading that [Redhat article](https://www.redhat.com/sysadmin/base64-encoding#:~:text=How%20Base64%20works,formats%20and%20require%20simple%20text.) wasn't enough so I had to re-implement it

> Please do not use this in a real project, it was not optimised to be fast, there are probably edge cases I did not catch either, this is solely to learn what goes on during Base64 encoding and decoding

## Usage

```bash
go build -o bin

./bin -e "Hello world"
> SGVsbG8gd29ybGQ= #output

./bin -d "SGVsbG8gd29ybGQ="
> Hello world #output
```

## Why not just read the implementation in other packages or languages?

No.
