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
go get github.com/Suenaa/agenda-golang
```

- run server
```
cd serverce
go run main.go
```

- use cli
```
go build agd.go
./agd [command]
```


**使用镜像**
- pull from docker hub
```
```

- use cli
```
```

# 测试
