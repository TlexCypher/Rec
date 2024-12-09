package api

import "github.com/labstack/echo/v4"

func RegisterInventory(ctx echo.Context,
	productName string, productCode string,
	remainingQuantity int,
	remarks *[]string) error {
	return nil
}
