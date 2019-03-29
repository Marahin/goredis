# goredis

## Running

1. `cp .env.dist .env`  

2. `make run`

## Having fun

```
$ telnet localhost 6379

Trying ::1...
Connected to localhost.
Escape character is '^]'.
PING
PONG
SET a raz dwa trzy
OK
GET a


raz dwa trzy
```
