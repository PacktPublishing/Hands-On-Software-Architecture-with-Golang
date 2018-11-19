

#create
curl -d '{"id":"xyz", "display_name":"HotelXyz", "star_rating":4, "no_rooms": 150}' -H "Content-Type: application/json"  -X POST 127.0.0.1:8080/v1/hotels

#get
curl 127.0.0.1:8080/v1/hotels/xyz

#get all 
curl 127.0.0.1:8080/v1/hotels

#update
curl -d '{"id":"xyz",  "star_rating":5}' -H "Content-Type: application/json"  -X PUT 127.0.0.1:8080/v1/hotels/xyz


#delete
curl -X DELETE 127.0.0.1:8080/v1/hotels/xyz