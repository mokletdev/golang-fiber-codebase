package User

import (
	"context"
	"net/http"
	"time"

	"github.com/mokletdev/golang-fiber-codebase/bin/modules/User/models"
	"github.com/mokletdev/golang-fiber-codebase/utils/res"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUsers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var users []models.User
	defer cancel()

	results, err := userCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(res.Response{Success: false, Data: err.Error(), Message: "error", Code: http.StatusInternalServerError})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleUser models.User
		if err = results.Decode(&singleUser); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(res.Response{Success: false, Data: err.Error(), Message: "error", Code: http.StatusInternalServerError})
		}

		users = append(users, singleUser)
	}

	return c.Status(http.StatusOK).JSON(
		res.Response{Success: true, Message: "Users successfully fetched.", Data: users, Code: http.StatusOK},
	)
}
