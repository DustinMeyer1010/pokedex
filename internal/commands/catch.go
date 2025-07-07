package commands

/*
//https://pokeapi.co/api/v2/pokemon/{id or name}/

func commandCatch(config *Config) error {

	if config.Arugments == "" {
		return fmt.Errorf("location missing\nUse map or mapb to get locations\nUsage: explore {location}")
	}

	res, _ := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", config.Arugments))

	if res.Status == "404 Not Found" {
		return fmt.Errorf("location not found")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", config.Arugments)

	var pokemon models.Pokemon

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &pokemon)

	if attemptCatch(pokemon.BaseExperience) {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		player.AddPokemon(pokemon)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

func attemptCatch(baseExperience int) bool {
	chance := 100 / float64(baseExperience)
	roll := rand.Float64()
	return roll < chance
}

*/
