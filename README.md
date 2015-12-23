# no-ansi
:wrench: remove ansi codes

[![GoDoc](https://godoc.org/github.com/moul/no-ansi?status.svg)](https://godoc.org/github.com/moul/no-ansi)

## Usage

```console
$ cat test.txt | hexdump -C
00000000  41 41 41 0a 1b 5b 31 3b  33 34 6d 20 42 42 42 1b  |AAA..[1;34m BBB.|
00000010  5b 30 6d 0a 43 43 43 0a  1b 5b 33 35 6d 20 44 44  |[0m.CCC..[35m DD|
00000020  44 0a 45 45 45 0a 1b 5b  30 6d 0a 46 46 46 0a     |D.EEE..[0m.FFF.|
0000002f
$ cat test.txt | ./no-ansi | hexdump -C
00000000  41 41 41 0a 20 42 42 42  0a 43 43 43 0a 20 44 44  |AAA. BBB.CCC. DD|
00000010  44 0a 45 45 45 0a 0a 46  46 46 0a                 |D.EEE..FFF.|
0000001b
$ ./no-ansi cat test.txt | hexdump -C
00000000  41 41 41 0a 20 42 42 42  0a 43 43 43 0a 20 44 44  |AAA. BBB.CCC. DD|
00000010  44 0a 45 45 45 0a 0a 46  46 46 0a                 |D.EEE..FFF.|
0000001b
```

## Install

```sh
go get github.com/moul/no-ansi/cmd/no-ansi
```

## License

MIT
