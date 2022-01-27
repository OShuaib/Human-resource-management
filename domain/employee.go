package domain

type Employee struct {
	ID 		string		`json:"id,omitempty" bson:"_id,omitempty"`
	Name 	string		`json:"name" validate="required, min=3, max=50"`
	Salary	float64		`json:"salary"`
	Age 	float64		`json:"age"`
}