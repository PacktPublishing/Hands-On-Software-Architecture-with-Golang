package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"math/rand"
	"net/http"
)

type Hotel struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
	City        string `json:"city"`
	NoRooms     int    `json:"noRooms"`
	StarRating  int    `json:"starRating"`
}

// define custom GraphQL ObjectType `hotelType` for our Golang struct `Hotel`
// Note that
// - the fields  map with the json tags for the fields in our struct
// - the field types match the field type in our struct
var hotelType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Hotel",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"displayName": &graphql.Field{
			Type: graphql.String,
		},
		"city": &graphql.Field{
			Type: graphql.String,
		},
		"noRooms": &graphql.Field{
			Type: graphql.Int,
		},
		"starRating": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

// define schema, with our rootQuery and rootMutation
var schema, schemaErr = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})

func main() {

	if schemaErr != nil {
		panic(schemaErr)
	}

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[in handler]", r.URL.Query())
		result := executeQuery(r.URL.Query()["query"][0], schema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Graphql server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

// root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createHotel": &graphql.Field{
			Type:        hotelType, // the return type for this field
			Description: "Create new hotel",
			Args: graphql.FieldConfigArgument{
				"displayName": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"city": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"noRooms": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"starRating": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// marshall and cast the argument value
				displayName, _ := params.Args["displayName"].(string)
				city, _ := params.Args["city"].(string)
				noRooms, _ := params.Args["noRooms"].(int)
				starRating, _ := params.Args["starRating"].(int)

				// create in 'DB'
				newHotel := Hotel{
					Id:          randomId(),
					DisplayName: displayName,
					City:        city,
					NoRooms:     noRooms,
					StarRating:  starRating,
				}
				hotels[newHotel.Id] = newHotel

				// return the new Hotel object
				return newHotel, nil
			},
		},
	},
})

//root query
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{

		"hotel": &graphql.Field{
			Type:        hotelType,
			Description: "Get a hotel with this id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, _ := params.Args["id"].(string)
				return hotels[id], nil
			},
		},
	},
})

// repository
var hotels map[string]Hotel

func init() {
	hotels = make(map[string]Hotel)
}

// Random ID Generator
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomId() string {
	b := make([]rune, 8)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
