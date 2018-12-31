#  Case Study - Travel Website

## Infra Needed

### kafka
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

### mysql
The code expects MySQL to be installed locally - it connects to root:@tcp(127.0.0.1:3306)
You can install MySQL on a mac using
```
brew install mysql@5.7
```

Once you install mysql, you can create the schema using the following DDL
```
CREATE TABLE hotel_reservations (
  customer_id int,
  payment_identifier int,
  sku int,
  room_id int,
  check_in datetime,
  check_out datetime,
  status int(10),
  id varchar(100),
  PRIMARY KEY (id)
);
```

## API Calls

# create booking api
```
curl -d '{"customer_id": 999, "payment_identifier": 66, "sku":1, "room_id": 2, "check_in" : "2018-12-07" , "check_out" : "2018-12-08"}' -H "Content-Type: application/json"  -X POST 127.0.0.1:8080/v1/hotels/reservation/
```

