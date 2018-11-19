
#table

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

# create booking api
curl -d '{"customer_id": 999, "payment_identifier": 66, "sku":1, "room_id": 2, "check_in" : "2018-12-07" , "check_out" : "2018-12-08"}' -H "Content-Type: application/json"  -X POST 127.0.0.1:8080/v1/hotels/reservation/

