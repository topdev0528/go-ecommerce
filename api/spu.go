package api

import (
	"fmt"
	"mall_go/internal/biz"

	"github.com/gofiber/fiber/v2"

	"mall_go/pkg/core"
	"mall_go/pkg/utils"
)

func SaleExplainFixed(c *fiber.Ctx) error {
	return nil
}
func Search(c *fiber.Ctx) error {
	q := c.Query("q")
	fmt.Println(q)
	return nil
}
func TagType(c *fiber.Ctx) error {
	types := c.Params("type")
	fmt.Println(types)
	return nil
}
func Detail(c *fiber.Ctx) error {

	id, err := utils.StringToUint(c.Params("id"))
	if err != nil {

		return err
	}
	data, err := biz.SpuById(id)
	if err != nil {

		return err
	}
	return core.SetData(c, data)

}
func Latest(c *fiber.Ctx) error {
	return nil
}
func SpuByCategory(c *fiber.Ctx) (err error) {
	id, err := utils.StringToUint(c.Params("id"))
	if err != nil {

		return
	}
	data, err := biz.SpuByCategory(id)
	if err != nil {

		return
	}
	return core.SetData(c, data)

}
