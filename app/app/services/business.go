package services

import "fmt"

type BusinessService struct {
}

type Params struct {
	A int
	B int
}

func (that *BusinessService) Get(params *Params, result *map[string]interface{}) error {
	*result = map[string]interface{}{
		"SUM": 1,
	}

	fmt.Printf("params %v,", params)
	return nil
}
