package agency

import (
	"fmt"

	"github.com/dev-newus/GoAlinBackend/src/models/entities"
	"github.com/dev-newus/GoAlinBackend/src/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAll(c *fiber.Ctx, tenant *gorm.DB) error {
	dataAgency := []entities.Agency{}
	result := tenant.Find(&dataAgency)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	return utils.SendResponse(c, 200, dataAgency, "", map[string]string{})
}
