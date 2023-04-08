package mongo

import (
	"context"
	"fmt"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/infrastructure/dto"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

const (
	counterCollection = "counters"
	collection        = "customers"
)

type customerRepository struct {
	logger  logger.Logger
	db      *mongo.Database
	timeout time.Duration
}

func NewCustomerRepository(baseLogger logger.Logger, db *mongo.Database, timeout time.Duration) model.CustomerRepository {
	repo := &customerRepository{
		logger:  baseLogger.WithFields(logger.Fields{"logger": "mongo.CustomerRepository", "collection": collection}),
		db:      db,
		timeout: timeout,
	}
	repo.createIndex(context.Background())
	return repo
}

func (r *customerRepository) FindById(ctx context.Context, id int64) (*model.Customer, error) {
	log := r.logger.WithFields(logger.Fields{"method": "FindById", "id": id})
	filter := bson.D{{
		"_id", id,
	}}
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	var customer model.Customer
	if err := r.db.Collection(collection).FindOne(timeout, filter).Decode(&customer); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, model.CustomerNotFoundErr{Id: id}
		}
		log.WithFields(logger.Fields{"err": err}).Errorf(fmt.Sprintf("couldn't retrieve documents with filter %s", filter))
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) Create(ctx context.Context, customer model.Customer) (*model.Customer, error) {
	log := r.logger.WithFields(logger.Fields{"method": "Create"})
	//TODO: Optimize get next id with find email already exist with FindByEmail
	customerId, err := r.getNextId(ctx)
	if err != nil {
		return nil, err
	}
	customer.Id = customerId
	log = log.WithFields(logger.Fields{"customer": customer})
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	if _, err := r.db.Collection(collection).InsertOne(timeout, customer); err != nil {
		if weEr, ok := err.(mongo.WriteException); ok {
			if len(weEr.WriteErrors) > 0 && weEr.WriteErrors[0].Code == 11000 {
				for _, writeError := range weEr.WriteErrors {
					log.Debugf("WriteError: %s", writeError.Message)
				}
				log.Infof("customer already exist")
				return nil, model.CustomerAlreadyExistError{Email: customer.Email}
			}
		}
		log.WithFields(logger.Fields{"err": err}).Error("couldn't create customer")
		return nil, err
	}
	log.Info("customer created successfully")
	return nil, nil
}

func (r *customerRepository) Update(ctx context.Context, customer model.Customer) (*model.Customer, error) {
	log := r.logger.WithFields(logger.Fields{"method": "Update", "customerToUpdate": customer})
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	updateRes, err := r.db.Collection(collection).UpdateByID(timeout, customer.Id, bson.M{"$set": customer})
	if err != nil {
		log.WithFields(logger.Fields{"error": err}).Error("error when update customer")
		return nil, err
	}
	if updateRes.ModifiedCount == 0 {
		log.Errorf("couldn't update customer with id %v", customer.Id)
		return nil, model.CustomerCannotUpdate{Id: customer.Id}
	}
	log.Info("customer updated successfully")
	return &customer, nil
}

func (r *customerRepository) Delete(ctx context.Context, id int64) (bool, error) {
	log := r.logger.WithFields(logger.Fields{"method": "Delete", "customerId": id})
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	deleteRes, err := r.db.Collection(collection).DeleteOne(timeout, bson.M{"_id": id})
	if err != nil {
		log.WithFields(logger.Fields{"error": err}).Errorf("error when delete customer with id %v", id)
		return false, err
	}
	if deleteRes.DeletedCount == 0 {
		log.Infof("customer id %v cannot be deleted", id)
		return false, model.CustomerCannotDelete{Id: id}
	}
	log.Infof("customer with id %v deleted successfully", id)
	return true, nil
}

func (r *customerRepository) getNextId(ctx context.Context) (int64, error) {
	log := r.logger.WithFields(logger.Fields{"method": "getNextId"})
	opts := options.RunCmd().SetReadPreference(readpref.Primary())
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	command := bson.D{
		{"findAndModify", counterCollection},
		{"query", bson.D{{"_id", collection}}},
		{"update", bson.D{{"$inc", bson.D{{"seq", 1}}}}},
		{"new", true},
	}
	var res dto.NextIdResponse
	if err := r.db.RunCommand(timeout, command, opts).Decode(&res); err != nil {
		log.WithFields(logger.Fields{"error": err}).Errorf("get next id error with counter collection %s and _id %s", counterCollection, collection)
		return 0, err
	}
	log.Debugf("get next id successful with value %v", res.Value.Seq)
	return res.Value.Seq, nil

}

func (r *customerRepository) createIndex(ctx context.Context) {
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	index := mongo.IndexModel{
		Keys:    bson.D{{"id", 1}, {"email", 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := r.db.Collection(collection).Indexes().CreateOne(timeout, index)
	if err != nil {
		r.logger.WithFields(logger.Fields{"error": err}).Fatalf("could not create mongo index")
	} else {
		r.logger.Infof("mongo index created")
	}
}
