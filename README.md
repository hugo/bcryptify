# bcryptify

Takes a password and cost and returns the bcrypt hash. Useful for e.g. seeding
databases

## Usage
```
$ go install https://github.com/thisishugo/bcryptify
$ bcryptify -p PASSWORD [-c COST]
$ bcryptify -p swordfish
2016/03/10 18:27:23 [bcryptify] generated hash
$2a$10$qh57MNoJYbGorZgOE2uMaeydrhCprazHMcIKbmOibm7RZT6YRJLMe
```
