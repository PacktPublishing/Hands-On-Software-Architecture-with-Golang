
#Create
curl -g 'http://localhost:8080/graphql?query=mutation+_{createHotel(displayName:"HotelX",city:"NY",noRooms:300,starRating:5){id}}'

#Get
curl -g 'http://localhost:8080/graphql?query={hotel(id:"XVlBzgba"){displayName,city,noRooms,starRating}}'
		


