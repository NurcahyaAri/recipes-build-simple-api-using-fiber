package handlers

import (
	"fmt"
	"recipes-fiber-basic-api/localdata"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ReadData(ctx *fiber.Ctx) error {
	data := localdata.InitData()
	return ctx.JSON(data.GetData())
}

func DeleteData(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	data := localdata.InitData()
	return ctx.JSON(data.DeleteData(id))
}

func InsertData(ctx *fiber.Ctx) error {
	data := localdata.InitData()
	user := new(localdata.UserModel)
	fmt.Println([]byte(ctx.Body()))
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	users := data.InsertData(localdata.UserModel{
		Name:   user.Name,
		Gender: user.Gender,
		Age:    user.Age,
	})

	return ctx.JSON(users)
}

func ReadDataById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	data := localdata.InitData()
	return ctx.JSON(data.GetDataById(id))
}

func PatchData(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	data := localdata.InitData()

	user := new(localdata.UserModel)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	users := data.UpdateDataById(id, localdata.UserModel{
		Name:   user.Name,
		Gender: user.Gender,
		Age:    user.Age,
	})

	return ctx.JSON(users)
}
