package entity

func Factory(name string) IDuck {
	switch name {
	case "duck":
		return &Duck{Color: "White", Age: 3}
	case "goose":
		return &Goose{Color: "Black"}
	default:
		panic("No such animal")
	}
}
