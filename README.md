# cmdMx

Command Multiplier is a simple utility that enables bulk modification of incrementing interfaces


| Flag | Description |
| ------ | ------ |
| -h | Show this help |
| -n | Number of times to repeat (def: 1) |
| -i | Increment the output (def: false) |
| -b | Increment by how many (def: 1) |
| -s | Starting replacement value (def: 1) |


# Indended Usage:
```sh
$ cmdMx -i -n 3 interface 0/%d\nshutdown
interface 0/1
shutdown
interface 0/2
shutdown
interface 0/3
shutdown
```

# Alternate Usage:
```sh
$ cmdMx -i -n 3 -b 2 interface 0/%d\nshutdown\ninterface 0/%d\nno shutdown
interface 0/1
shutdown
interface 0/3
no shutdown
interface 0/5
shutdown
interface 0/7
no shutdown
interface 0/9
shutdown
interface 0/11
no shutdown
```

# Because of the way strings are separated, this is also acceptable:
```sh
$ cmdMx -i -s 2 'this is a %d test' 'that is also a %d test'
this is a 2 test that is also a 3 test
this is a 4 test that is also a 5 test
```

# A byproduct of the string separation is where multiple matches are in the same string:
```sh
$ cmdMx -i -n 3 -b 2 'this is a %d test %d good' test this %d good %d test
this is a 1 test 1 good test this 3 good 5 test
this is a 7 test 7 good test this 9 good 11 test
this is a 13 test 13 good test this 15 good 17 test
```
