package structs

type apps struct {
	all []app
}

type app struct {
	vars map[string]string
}
