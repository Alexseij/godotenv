package autoload

/*
	You can just read the .env file on import just by doing

		import _ "github.com/joho/godotenv/autoload"

	And bob's your mother's brother
*/

import "github.com/Alexseij/godotenv/v1"

func init() {
	godotenv.Load()
}
