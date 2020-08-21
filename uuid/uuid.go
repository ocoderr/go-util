package uuid

import "github.com/google/uuid"

func GenUUID() string  {
	random, _ := uuid.NewRandom()
	return random.String()
}
