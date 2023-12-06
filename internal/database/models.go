package database

type Swap struct {
	ID              string
	SourceChain     string
	DestinationChain string
	Amount          float64
	Status          string
	//TODO: Add other fields as needed
}
