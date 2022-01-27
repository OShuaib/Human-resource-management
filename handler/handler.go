package handler

import (
	conn "github.com/OShuaib/Human-resource-management/Db"
	"github.com/OShuaib/Human-resource-management/domain"	
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func GetEmployee(c *fiber.Ctx) {

	query := bson.D{{}}
	cursor, err  := conn.Mg.Db.Collection("employees").Find(c.Context(),query)
	if err != nil {
		 c.Status(500).SendString(err.Error())
		 return
	}
	var employees []domain.Employee = make([]domain.Employee, 0)
	if err := cursor.All(c.Context(), &employees); err != nil {
		
		c.Status(500).SendString(err.Error())
		return
	}


	c.JSON(employees)
}

func CreateEmployee(c *fiber.Ctx)  {
	collection := conn.Mg.Db.Collection("employees")

	employee := new(domain.Employee)

	if err := c.BodyParser(employee); err != nil {
		 c.Status(400).SendString(err.Error())
		 return 
	}
	//employee.ID = ""

	insert, err := collection.InsertOne(c.Context(),employee)
	if err != nil {
		c.Status(400).SendString(err.Error())
	}

	filter := bson.D{{Key:"_ id", Value: insert.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	createdEmployee := &domain.Employee{}
	createdRecord.Decode(&createdEmployee)

	 c.Status(201).JSON(createdEmployee)
	 return
}

func UpdateEmployee(c *fiber.Ctx) {
	idParam := c.Params("id")

	employeeID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.SendStatus(400)
		return
	}

	employee := new(domain.Employee)

	if err := c.BodyParser(employee); err != nil {
		c.Status(400).SendString(err.Error())
		return
	}

	query := bson.D{{Key: "_id", Value: employeeID}}
	update := bson.D{
		{Key: "$set",
		Value: bson.D{
			{Key: "name", Value: employee.ID},
			{Key:"age", Value: employee.Age},
			{Key: "salary", Value: employee.Salary},
		},
	},
}

err = conn.Mg.Db.Collection("employees").FindOneAndUpdate(c.Context(), query, update).Err()

if err != nil {
	if err == mongo.ErrNoDocuments{
		c.SendStatus(400)
		return
	}
	c.SendStatus(500)
	return
}
  employee.ID = idParam

 c.Status(200).JSON(employee)
}

func DeleteEmployee(c *fiber.Ctx) {
	employeeID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		 c.SendStatus(400)
		 return
	}

	query := bson.D{{Key: "_id", Value: employeeID}}
	result, err := conn.Mg.Db.Collection("employees").DeleteOne(c.Context(),&query)

	if err != nil {
		c.SendStatus(500)
		return
	}

	if result.DeletedCount < 1 {
		c.SendStatus(404)
		return
	}

	c.Status(200).JSON("record deleted")

}