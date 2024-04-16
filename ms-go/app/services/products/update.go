package products

import (
	"context"
	"ms-go/app/helpers"
	"ms-go/app/models"
	"ms-go/db"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func Update(data models.Product, isAPI bool) (*models.Product, error) {

	if data.ID == 0 {
		return nil, &helpers.GenericError{Msg: "Missing parameters", Code: http.StatusUnprocessableEntity}
	}

	var product models.Product

	if err := db.Connection().FindOne(context.TODO(), bson.M{"id": data.ID}).Decode(&product); err != nil {
		return nil, &helpers.GenericError{Msg: "Product Not Found", Code: http.StatusNotFound}
	}

	data.UpdatedAt = time.Now()

	setUpdate(&data, &product)

	if err := db.Connection().FindOneAndUpdate(context.TODO(), bson.M{"id": data.ID}, bson.M{"$set": data}).Decode(&product); err != nil {
		return nil, &helpers.GenericError{Msg: err.Error(), Code: http.StatusInternalServerError}
	}

	if err := db.Connection().FindOne(context.TODO(), bson.M{"id": data.ID}).Decode(&product); err != nil {
		return nil, &helpers.GenericError{Msg: "Product Not Found", Code: http.StatusNotFound}
	}

	defer db.Disconnect()

	if isAPI {
	}

	return &product, nil
}

func setUpdate(new, old *models.Product) {
	if new.ID == 0 {
		new.ID = old.ID
	}

	if new.Name == "" {
		new.Name = old.Name
	}

	if new.Brand == "" {
		new.Brand = old.Brand
	}

	if new.Price == 0 {
		new.Price = old.Price
	}

	if new.Description == "" {
		new.Description = old.Description
	}

	new.CreatedAt = old.CreatedAt

	new.UpdatedAt = time.Now()
}
