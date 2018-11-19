package main

import (
	"errors"
	"fmt"
	"github.com/mediocregopher/radix.v2/redis"
	"strconv"
)

type Hotel struct {
	Id         string
	Name       string
	City       string
	StarRating int
	Likes      int
}

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// save an object in redis
	taj := Hotel{"taj", "Hotel Taj", "Mumbai", 4, 0}
	fmt.Println(taj)
	setHotel(conn, &taj)

	// retrieve an object from redis
	hotel, err := getHotel(conn, "taj")
	fmt.Println(hotel, err)

	//simulate some likes
	// 3 likes for Taj
	err = incrementLikes(conn, "taj")
	if err != nil {
		fmt.Println("incrementLikes error ", err)
	}
	incrementLikes(conn, "taj")
	incrementLikes(conn, "taj")

	// some other hotels to test the sorted set
	leela := Hotel{"leela", "Leeka K", "Bangalore", 4, 0}
	setHotel(conn, &leela)
	// 2 likes for leeka
	incrementLikes(conn, "leela")
	incrementLikes(conn, "leela")

	notGood := Hotel{"notSoGood", "HotelXX", "Bangalore", 4, 0}
	setHotel(conn, &notGood)
	// only 1 like
	incrementLikes(conn, "notSoGood")

	top3, err := top3LikedHotels(conn)
	fmt.Println(top3, err)

}

func getHotel(conn *redis.Client, id string) (*Hotel, error) {
	reply, err := conn.Cmd("HGETALL", "hotels:"+id).Map()
	if err != nil {
		return nil, err
	}

	h := new(Hotel)
	h.Id = id
	h.Name = reply["name"]
	h.City = reply["city"]
	if h.Likes, err = strconv.Atoi(reply["likes"]); err != nil {
		fmt.Println("likes err", err)
		return nil, err
	}
	if h.StarRating, err = strconv.Atoi(reply["rating"]); err != nil {
		fmt.Println("ratings err", err)
		return nil, err
	}

	return h, nil
}

func setHotel(conn *redis.Client, h *Hotel) error {
	resp := conn.Cmd("HMSET",
		"hotels:"+h.Id,
		"name", h.Name,
		"city", h.City,
		"likes", h.Likes,
		"rating", h.StarRating)
	if resp.Err != nil {
		fmt.Println("save err", resp.Err)
		return resp.Err
	}

	return nil
}

func incrementLikes(conn *redis.Client, id string) error {

	//  Sanity check to ensure that the hotel exists!
	exists, err := conn.Cmd("EXISTS", "hotels:"+id).Int()
	if err != nil || exists == 0 {
		return errors.New("no such hotel")
	}

	// Use the MULTI command to inform Redis that we are starting a new
	// transaction.
	err = conn.Cmd("MULTI").Err
	if err != nil {
		return err
	}

	// Increment the number of likes  for the hotel. in the album hash by 1.
	// Because we have initiated a  MULTI command, this HINCRBY command is queued NOT executed.
	// We still check the reply's Err field  to check if there was an error for the queing
	err = conn.Cmd("HINCRBY", "hotels:"+id, "likes", 1).Err
	if err != nil {
		return err
	}

	// Now we increment the leaderboard sorted set
	err = conn.Cmd("ZINCRBY", "likes", 1, id).Err
	if err != nil {
		return err
	}

	// Execute both commands in our transaction atomically.
	// EXEC returns the replies from both commands as an array
	err = conn.Cmd("EXEC").Err
	if err != nil {
		return err
	}
	return nil
}

func top3LikedHotels(conn *redis.Client) ([]string, error) {
	// Use the ZREVRANGE command to fetch the hotels from likes sorted set
	// with  the highest score first
	// The start and stop values are zero-based indexes, so we use 0 and 2
	// respectively to limit the reply to the top three.

	reply, err := conn.Cmd("ZREVRANGE", "likes", 0, 2).List()
	if err != nil {
		return nil, err
	}

	return reply, nil

}
