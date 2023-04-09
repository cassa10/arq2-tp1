package mongo

import (
	"context"
	"fmt"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const productCollection = "products"

type productRepository struct {
	logger  logger.Logger
	db      *mongo.Database
	timeout time.Duration
}

func NewProductRepository(baseLogger logger.Logger, db *mongo.Database, timeout time.Duration) model.ProductRepository {
	repo := &productRepository{
		logger:  baseLogger.WithFields(logger.Fields{"logger": "mongo.ProductRepository", "productCollection": productCollection}),
		db:      db,
		timeout: timeout,
	}
	repo.createIndex(context.Background())
	return repo
}

func (r *productRepository) FindById(ctx context.Context, id int64) (*model.Product, error) {
	log := r.logger.WithFields(logger.Fields{"method": "FindById", "id": id})
	filter := bson.M{"_id": id}
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	var product model.Product
	if err := r.db.Collection(productCollection).FindOne(timeout, filter).Decode(&product); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.ProductNotFoundErr{Id: id}
		}
		log.WithFields(logger.Fields{"err": err}).Errorf(fmt.Sprintf("couldn't retrieve documents with filter %s", filter))
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Update(ctx context.Context, product model.Product) (bool, error) {
	log := r.logger.WithFields(logger.Fields{"method": "Update", "productToUpdate": product})
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	updateRes, err := r.db.Collection(productCollection).UpdateByID(timeout, product.Id, bson.M{"$set": product})
	if err != nil {
		log.WithFields(logger.Fields{"error": err}).Error("error when update product")
		return false, err
	}
	if updateRes.ModifiedCount == 0 {
		log.Errorf("couldn't update product with id %v", product.Id)
		return false, exception.ProductCannotUpdate{Id: product.Id}
	}
	log.Info("product updated successfully")
	return true, nil
}

func (r *productRepository) Delete(ctx context.Context, id int64) (bool, error) {
	log := r.logger.WithFields(logger.Fields{"method": "Delete", "productId": id})
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	deleteRes, err := r.db.Collection(productCollection).DeleteOne(timeout, bson.M{"_id": id})
	if err != nil {
		log.WithFields(logger.Fields{"error": err}).Errorf("error when delete product with id %v", id)
		return false, err
	}
	if deleteRes.DeletedCount == 0 {
		log.Infof("product id %v cannot be deleted", id)
		return false, exception.ProductCannotDelete{Id: id}
	}
	log.Infof("product with id %v deleted successfully", id)
	return true, nil
}

func (r *productRepository) Create(ctx context.Context, product model.Product) (int64, error) {
	log := r.logger.WithFields(logger.Fields{"method": "Create"})
	productId, err := getNextId(ctx, log, r.db, r.timeout, productCollection)
	if err != nil {
		return 0, err
	}
	product.Id = productId

	log = log.WithFields(logger.Fields{"product": product})
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	if _, err := r.db.Collection(productCollection).InsertOne(timeout, product); err != nil {
		log.WithFields(logger.Fields{"error": err}).Error("couldn't create product")
		return 0, err
	}
	log.Info("product created successfully")
	return product.Id, nil
}

func (r *productRepository) FindAllBySellerId(ctx context.Context, sellerId int64) ([]model.Product, error) {
	log := r.logger.WithFields(logger.Fields{"method": "FindById", "sellerId": sellerId})
	filter := bson.M{"sellerId": sellerId}
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()

	cur, err := r.db.Collection(productCollection).Find(timeout, filter)
	if err != nil {
		log.WithFields(logger.Fields{"error": err}).Errorf(fmt.Sprintf("something went wrong when find with filter %s", filter))
		return nil, err
	}
	products := make([]model.Product, 0)
	for cur.Next(timeout) {
		var product model.Product
		if err := cur.Decode(&product); err != nil {
			log.WithFields(logger.Fields{"error": err}).Errorf(fmt.Sprintf("something went wrong when want decode product"))
			return products, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *productRepository) Search(ctx context.Context, filters model.ProductSearchFilter) ([]model.Product, error) {
	panic("not implemented yet")
}

func (r *productRepository) createIndex(ctx context.Context) {
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{{"sellerId", 1}},
		},
		{
			Keys: bson.D{
				{"price", 1},
				{"category", 1},
				{"stock", -1},
			},
		},
	}
	_, err := r.db.Collection(productCollection).Indexes().CreateMany(timeout, indexes)
	if err != nil {
		r.logger.WithFields(logger.Fields{"error": err}).Fatalf("could not create mongo index")
	} else {
		r.logger.Infof("mongo index created")
	}
}
