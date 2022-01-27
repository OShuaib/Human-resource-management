package handler

import (
	conn "github.com/OShuaib/Human-resource-management/Db"
	"github.com/OShuaib/Human-resource-management/domain"	
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func GetEmployee() func(c *fiber.Ctx) error {

	query := bson.D{{}}
	cursor, err  := conn.Mg.Db.Collection("employees").Find(c.Context(),query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	var employees []domain.Employee = make([]domain.Employee, 0)
	if err := cursor.All(c.Context(), &employees); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(employees)
}

func CreateEmployee() func(c *fiber.Ctx) error {
	collection := conn.Mg.Db.Collection("employees")

	employee := new(domain.Employee)

	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	employee.ID = ""

	insert, err := collection.InsertOne(c.Context(),employee)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	filter := bson.D{{Key:"_ id", Value: insert.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	createdEmployee := &domain.Employee{}
	createdRecord.Decode(createdEmployee)

	return c.Status(201).JSON(createdEmployee)
}

func UpdateEmployee() func(c *fiber.Ctx) error {
	idParam := c.Params("id")

	employeeID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.SendStatus(400)
	}

	employee := new(domain.Employee)

	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).SendString(err.Error())
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
		return c.SendStatus(400)
	}
	return c.SendStatus(500)
}
employee.ID = idParam

return c.Status(200).JSON(employee)
}

func DeleteEmployee() func(c *fiber.Ctx) error {
	employeeID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.SendStatus(400)
	}

	query := bson.D{{Key: "_id", Value: employeeID}}
	result, err := conn.mg.Db.Collection("employees").DeleteOne(c.Context(),&query)

	if err != nil {
		return c.SendStatus(500)
	}

	if result.DeletedCount < 1 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON("record deleted")

}