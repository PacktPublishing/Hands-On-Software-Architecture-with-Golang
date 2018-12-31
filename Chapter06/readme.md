## Messaging
### This package contains the code shown in the messaging chapter (6).

Note : each file is an independent go program, which can be run via 'go run'

### Kafka
The code expects Kafka to be running lon localhost:9092
You can do install Kafka on a mac using
```
brew cask install java
brew install kafka
```
and then start is using
```
zkServer start
kafka-server-start.sh /usr/local/etc/kafka/server.properties
```

### nsq
To install  and run nsq  please see [nsq quick start](https://nsq.io/overview/quick_start.html)
