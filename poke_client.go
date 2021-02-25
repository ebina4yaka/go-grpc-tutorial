package main

import (
	"context"
	"fmt"
	"github.com/ebina4yaka/go-grpc-tutorial/pokemon"
	"google.golang.org/grpc"
	"log"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := pokemon.NewPokemonServiceClient(conn)

	query := pokemon.PokeQuery{IndexNo: 120, FormType: 0}

	response, err := client.GetPokemon(context.Background(), &query)
	if err != nil {
		log.Fatalf("Error when calling GetPokemon: %s", err)
	}

	fmt.Printf("Pokemon No: %d\n", response.IndexNo)
	fmt.Printf("Pokemon Name: %s\n", response.Name)
	if response.FormType != 0 {
		fmt.Printf("Pokemon Form: %s\n", response.FormName)
	}
	if response.Type2 != "" {
		fmt.Printf("Pokemon Type 1: %s\n", response.Type1)
		fmt.Printf("Pokemon Type 2: %s\n", response.Type2)
	} else {
		fmt.Printf("Pokemon Type: %s\n", response.Type1)
	}
	if response.Ability2 != "" {
		fmt.Printf("Pokemon Ability 1: %s\n", response.Ability1)
		fmt.Printf("Pokemon Ability 2: %s\n", response.Ability2)
	} else {
		fmt.Printf("Pokemon Ability: %s\n", response.Ability1)
	}
	fmt.Printf("Pokemon Hidden Ability: %s\n", response.HiddenAbility)
	fmt.Printf("Pokemon Base stats HP: %d\n", response.Hp)
	fmt.Printf("Pokemon Base stats Attack: %d\n", response.Attack)
	fmt.Printf("Pokemon Base stats Defence: %d\n", response.Defence)
	fmt.Printf("Pokemon Base stats Special Attack: %d\n", response.SpecialAttack)
	fmt.Printf("Pokemon Base stats Special Defence: %d\n", response.SpecialDefence)
	fmt.Printf("Pokemon Base stats Speed: %d\n", response.Speed)
}
