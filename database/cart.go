package database

import (
	"context"
	"errors"
	"log"

	"github.com/DieFleitas/ecommerce-akhil/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var(
	ErrCantFindProduct = errors.New("Can't find the product")
	ErrCantDecodeProducts = errors.New("Can't find the product")
	ErrUserIdIsNotValid = errors.New("This user is not valid")
	ErrCantUpdateUser = errors.New("Cannot add this product to the cart")
	ErrCantRemoveItemCart = errors.New("Cannot remove this item from the cart")
	ErrCantGetItem = errors.New("Was unable to get the item from the cart")
	ErrCantBuyCartItem = errors.New("Cannot update the purchase")
)

func AddProductToCart(ctx context.Context, prodCollection, userCollection *mongo.Collection, productID primitive.ObjectID, userID string) error {
	searchfromdb, err := prodCollection.Find(ctx, bson.M{"_id": productID})
	if err != nil {
		log.Println(err)
		return ErrCantFindProduct
	}
	var productCart []models.ProductUser
	err = searchfromdb.All(ctx, &productCart)
	if err != nil {
		log.Println(err)
		return ErrCantDecodeProducts
	}

	id, err :=primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return ErrUserIdIsNotValid
	}

	filter := bson.D{primitive.E{Key:"_id", Value: id}}
	update := bson.D{{Key:"$push", Value: bson.D{primitive.E{Key:"usercart", Value: bson.D{{Key:"$each", Value: productCart}}}}}}

	_, err = userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return ErrCantUpdateUser
	}
}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}