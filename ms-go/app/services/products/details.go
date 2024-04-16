package products

import (
	"context"
	"ms-go/app/helpers"
	"ms-go/app/models"
	"ms-go/db"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func Details(data models.Product) (*models.Product, error) {

	if data.ID == 0 {
		return nil, &helpers.GenericError{Msg: "Missing params", Code: http.StatusUnprocessableEntity}
	}

	var product models.Product
	if err := db.Connection().FindOne(context.TODO(), bson.M{"id": data.ID}).Decode(&product); err != nil {
		return nil, &helpers.GenericError{Msg: "Product Not Found", Code: http.StatusNotFound}
	}

	defer db.Disconnect()

	return &product, nil
}
