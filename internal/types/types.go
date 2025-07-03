package types

type Cluster struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Servers int    `json:"servers"`
}