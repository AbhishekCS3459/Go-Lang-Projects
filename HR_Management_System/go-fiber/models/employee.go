package models

type Employee struct {
	ID     string  `json:"id" bson:"_id,omitempty"`
	Name   string  `json:"name" bson:"name"`
	Salary float64 `json:"salary" bson:"salary"`
	Age    float64 `json:"age" bson:"age"`
}
