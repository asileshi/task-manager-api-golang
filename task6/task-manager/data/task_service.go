package data

import (
	"context"

	"reflect"
	"task-manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetTasks() ([]models.Task, error) {
    var tasks []models.Task
    cursor, err := taskCollection.Find(context.TODO(), bson.D{{}})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())
    err = cursor.All(context.Background(), &tasks)
    if err != nil {
        return nil, err
    }
    return tasks, nil
}

func IsZeroValue(value interface{}) bool {
	return reflect.ValueOf(value).IsZero()
}

func GetTaskByID(id string) (models.Task, error){
    var task models.Task
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil{
        return task, err
    }

    filter := bson.D{{"_id",objectID}}

    err = taskCollection.FindOne(context.TODO(), filter).Decode(&task)
    if err != nil{
        return models.Task{}, err
    }

    return task, nil

}

func CreateTask(task models.Task) (models.Task, error){
	result, err := taskCollection.InsertOne(context.TODO(), task)
    if err != nil {
        return models.Task{}, err
    }

    
    //insertedTaskID := result.InsertedID.(primitive.ObjectID)

    var insertedTask models.Task
    err = taskCollection.FindOne(context.TODO(), bson.M{"_id": result.InsertedID}).Decode(&insertedTask)
    if err != nil {
        return models.Task{}, err
    }

    return insertedTask, nil
}
func UpdateTask(id string, updatedTask models.Task) (models.Task, error) {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return models.Task{}, err
    }

    filter := bson.D{{"_id", objectID}}

    updateFields := bson.M{}

    v := reflect.ValueOf(updatedTask)
	typeOfTask := v.Type()

	for i := 0; i < v.NumField(); i++ {
        field := typeOfTask.Field(i).Name
        value := v.Field(i).Interface()
        
        if !IsZeroValue(value){
            updateFields[field] = value
        }
	}

    update := bson.D{{"$set", updateFields}}

    _, err = taskCollection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        return models.Task{}, err
    }

    // Fetch the updated document from the database
    var result models.Task
    err = taskCollection.FindOne(context.TODO(), filter).Decode(&result)
    if err != nil {
        return models.Task{}, err
    }

    return result, nil
}

func DeleteTask(id string) error{
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil{
        return err
    }

    filter := bson.D{{"_id",objectID}}

    _, err = taskCollection.DeleteOne(context.TODO(), filter)
    if err != nil{
        return err
    }

    return nil
}