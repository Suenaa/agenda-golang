# agenda-golang
[![Build Status](https://travis-ci.org/Suenaa/agenda-golang.svg?branch=master)](https://travis-ci.org/Suenaa/agenda-golang)

# 使用说明
**agd command**

- regist       
　　`agd regist -u usernaeme -p password -e e-mail -t telephone`     

- login    
　　`agd login -u username -p password`

- logout    
　　`agd logout`

- list all users    
　　`agd lsu`

- delete current account         
　　`agd del -p password`

- create a meeting                      
　　`agd cm -t title -p participator1 -p participator2 -s start -e end`

- change the participants of a meeting                      
　　add: `agd ap -t title -p name`                                   
　　delete: `agd dp -t title -p name`

- list meetings during a period                    
　　`agd lsm -s start -e end`

- cancel a meeting                      
　　`agd cancel -t title`

- quit a meeing                   
　　`agd quit -t title`

- clear all meetings                  
　　`agd clear`


**使用命令行**
- go get
```
go get github.com/Suenaa/agenda-golang/cli/cmd
```

- run server
```
cd service
go run main.go
```

- use cli
```
cd cli
go build agd.go
./agd [command]
```


**使用镜像**
- pull from docker hub
```
$ sudo docker pull suenaa/agenda-golang
Using default tag: latest
latest: Pulling from suenaa/agenda-golang
f49cf87b52c1: Already exists 
7b491c575b06: Already exists 
b313b08bab3b: Already exists 
215a2061b8a4: Already exists 
04fa9dcc5f7d: Already exists 
044102b3b4a1: Already exists 
37e616f9e3fe: Already exists 
5031fa86d238: Pull complete 
d7ea7a7c9ce7: Pull complete 
32e5211ea8a7: Pull complete 
Digest: sha256:7e408efe4717d4af80e4b90efe2800457b5fb5b538c2dcf16a877ef582a6125d
Status: Downloaded newer image for suenaa/agenda-golang:latest

$ sudo docker images
REPOSITORY             TAG                 IMAGE ID            CREATED             SIZE
suenaa/agenda-golang   latest              fb135b17de0d        13 hours ago        780MB
hello-world            latest              f2a91732366c        3 weeks ago         1.85kB

```

- run server
```
$ sudo docker run -dit --name agenda-golang -v $HOME/Desktop/server:/data -p 8080:8080 suenaa/agenda-golang service
ae60d43ef433e7a7f8917b27b0fdaf45dbb7443cf141b663bb89a4427cb8d441

```
- use cli
```
$ sudo docker run --rm --net host -v $HOME/Desktop/client:/data suenaa/agenda-golang cli
A tool for managing meetings

Usage:
  agd [command]

Available Commands:
  ap          add participants to a meeting
  cancel      cancel a meeting
  clear       clear all meetings you create
  cm          create a meeting
  del         delete current account
  dp          delete participants in a meeting
  help        Help about any command
  login       log in
  logout      log out
  lsm         list all meetings during a period
  lsu         list all users
  quit        quit a meeting
  regist      regist a new user

Flags:
  -h, --help   help for agd

Use "agd [command] --help" for more information about a command.

```

**examples**
- register
```
$ sudo docker run --rm --net host -v $HOME/Desktop/client:/data suenaa/agenda-golang cli regist -u user1 -p 123 -e email1@mail.com -t 11111111
Success
$ sudo docker run --rm --net host -v $HOME/Desktop/client:/data suenaa/agenda-golang cli regist -u user2 -p 123 -e email2@mail.com -t 222222222222
Success
$ sudo docker run --rm --net host -v $HOME/Desktop/client:/data suenaa/agenda-golang cli regist -u user3 -p 123 -e email3@mail.com -t 33333333333333333
Success
```

- list all users
```
$ sudo docker run --rm --net host -v $HOME/Desktop/client:/data suenaa/agenda-golang cli lsu
[
  {
    "id": 1,
    "username": "user1",
    "password": "******",
    "email": "email1@mail.com",
    "phone": "11111111"
  },
  {
    "id": 2,
    "username": "user2",
    "password": "******",
    "email": "email2@mail.com",
    "phone": "222222222222"
  },
  {
    "id": 3,
    "username": "user3",
    "password": "******",
    "email": "email3@mail.com",
    "phone": "33333333333333333"
  }
]

```

- login
```
$ sudo docker run --rm --net host -v $HOME/Desktop/client:/data suenaa/agenda-golang cli login -u user1 -p 123
Success

```

- log out
```
$ sudo docker run --rm --net host -v $HOME/Desktop/client:/data suenaa/agenda-golang cli logout
Success
```

- delete current user
```
$ sudo docker run --rm --net host -v $HOME/Desktop/client:/data suenaa/agenda-golang cli del -p 123
Success
$ sudo docker run --rm --net host -v $HOME/Desktop/client:/data suenaa/agenda-golang cli lsu
[
  {
    "id": 2,
    "username": "user2",
    "password": "******",
    "email": "email2@mail.com",
    "phone": "222222222222"
  },
  {
    "id": 3,
    "username": "user3",
    "password": "******",
    "email": "email3@mail.com",
    "phone": "33333333333333333"
  }
]

```

- create a meeting
```
$ sudo docker run --rm --net host -v $HOME/Desktop/client:/data suenaa/agenda-golang cli cm -t title -p user1, user2 -s 2017-11-01T10:00 -e 2017-11-01T10:30
Success
```

- list meetings during a period
```
$ sudo docker run --rm --net host -v $HOME/Desktop/client:/data suenaa/agenda-golang cli lsm -s 2017-11-01T10:00 -e 2017-11-01T10:30
[
  {
    "id": 1,
    "title": "title",
    "sponsor": "user1",
    "participators": [
      "user2", "user3"
    ],
    "start": "2017-11-01 10:00",
    "end": "2017-11-01 10:30"
  }
]


```

- clear all meetings
```
$ sudo docker run --rm --net host -v $HOME/Desktop/client:/data suenaa/agenda-golang cli clear
Success
$ sudo docker run --rm --net host -v $HOME/Desktop/client:/data suenaa/agenda-golang cli lsm -s 2017-11-01T10:00 -e 2017-11-01T10:30
[]

```

# 测试
- ab TEST
```
$ ab -n 10000 -c 1000 "http://localhost:8080/user/login?username=user2&password=123"
This is ApacheBench, Version 2.3 <$Revision: 1528965 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/login?username=user2&password=123
Document Length:        3472 bytes

Concurrency Level:      1000
Time taken for tests:   2.029 seconds
Complete requests:      10000
Failed requests:        631
   (Connect: 0, Receive: 0, Length: 631, Exceptions: 0)
Non-2xx responses:      10000
Total transferred:      35880631 bytes
HTML transferred:       34720631 bytes
Requests per second:    4927.59 [#/sec] (mean)
Time per request:       202.939 [ms] (mean)
Time per request:       0.203 [ms] (mean, across all concurrent requests)
Transfer rate:          17266.11 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   22 142.2      0    1001
Processing:     0  159 221.5     99    1624
Waiting:        0  159 221.6     99    1624
Total:          0  181 294.5     99    2018

Percentage of the requests served within a certain time (ms)
  50%     99
  66%    165
  75%    169
  80%    171
  90%    288
  95%   1048
  98%   1214
  99%   1379
 100%   2018 (longest request)

```
