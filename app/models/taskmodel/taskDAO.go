package taskmodel

import (
	"context"
	"strings"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var (
	COLLECTION_NAME = "tasks"
)

type PageQueryArgs struct {
	Keyword string `json:"keyword,omitempty" form:"keyword,omitempty" binding:"-"`
	Limit   int    `json:"limit,default=100" form:"limit,default=100" binding:"required,number"`
	Order   string `json:"order,default=desc" form:"order,default=desc" binding:"oneof=desc asc"`
	By      string `json:"by,default=updated_at" form:"by,default=updated_at" binding:"-"`
	Page    int    `json:"page,default=1" form:"page,default=1" binding:"number"`
}

type TaskDAO struct {
	collection *mongo.Collection
}

func NewDAO(client *mongo.Client) *TaskDAO {
	return &TaskDAO{
		collection: client.Database(viper.GetString("mongo.database")).Collection(COLLECTION_NAME),
	}
}

func NewDAOwithName(client *mongo.Client, dbName, collName string) *TaskDAO {
	return &TaskDAO{
		collection: client.Database(dbName).Collection(collName),
	}
}

func (dao *TaskDAO) Create(message *TaskDTO) (*mongo.InsertOneResult, error) {
	message.CreatedAt = time.Now()
	message.UpdatedAt = time.Now()
	result, err := dao.collection.InsertOne(context.Background(), message)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (dao *TaskDAO) GetByID(id string) (*TaskDTO, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}
	var task TaskDTO
	err = dao.collection.FindOne(context.Background(), filter).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (dao *TaskDAO) GetByQuery(filter interface{}, limit int64, skip int64, order string, by string) (*[]TaskDTO, error) {
	var tasks []TaskDTO
	ctx := context.Background()
	opts := options.Find().SetSkip(skip).SetLimit(limit)
	// Handle Order
	// https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/sort/
	order = strings.ToLower(order)
	if order == "desc" {
		opts = opts.SetSort(bson.M{by: -1})
	} else {
		opts = opts.SetSort(bson.M{by: 1})
	}

	cur, err := dao.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	if err = cur.All(ctx, &tasks); err != nil {
		panic(err)
	}

	return &tasks, nil
}

func (dao *TaskDAO) Update(id string, updateFields map[string]interface{}) (*mongo.UpdateResult, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}
	updateFields["updated_at"] = time.Now()
	update := bson.M{"$set": updateFields}
	result, err := dao.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (dao *TaskDAO) Delete(id string) (*mongo.DeleteResult, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}
	result, err := dao.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/count/
func (dao *TaskDAO) CountDocuments(filter interface{}) (int64, error) {
	ctx := context.Background()
	count, err := dao.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}
