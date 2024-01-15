package databases

import (
	"context"
	"github.com/muazhari/logi-backend-1/src/inners/models/value_objects"
	databaseDatastores "github.com/muazhari/logi-backend-1/src/outers/datastores/databases"
	"github.com/muazhari/logi-backend-1/src/outers/utilities"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type SettingDatabaseRepository struct {
	OneDatabaseDatastore *databaseDatastores.OneDatabaseDatastore
}

func NewSettingDatabaseRepository(oneDatabaseDatastore *databaseDatastores.OneDatabaseDatastore) *SettingDatabaseRepository {
	settingDatabaseRepository := &SettingDatabaseRepository{
		OneDatabaseDatastore: oneDatabaseDatastore,
	}
	return settingDatabaseRepository
}

func (settingDatabaseRepository *SettingDatabaseRepository) ReadOne() (output *value_objects.Setting, err error) {
	database := settingDatabaseRepository.OneDatabaseDatastore.Configuration.Database
	databaseQuery := settingDatabaseRepository.OneDatabaseDatastore.Client.Database(database)
	collectionQuery := databaseQuery.Collection("setting")

	findOneCtx, findOneCtxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer findOneCtxCancel()
	findOneQuery := collectionQuery.FindOne(findOneCtx, bson.M{})
	decodeErr := findOneQuery.Decode(&output)
	if decodeErr != nil {
		err = decodeErr
	}

	return output, err
}

func (settingDatabaseRepository *SettingDatabaseRepository) PatchOne(setting *value_objects.Setting) (err error) {
	database := settingDatabaseRepository.OneDatabaseDatastore.Configuration.Database
	databaseQuery := settingDatabaseRepository.OneDatabaseDatastore.Client.Database(database)
	collectionQuery := databaseQuery.Collection("setting")

	readSetting, readOneSettingErr := settingDatabaseRepository.ReadOne()
	if readOneSettingErr != nil {
		err = readOneSettingErr
	}

	replaceOne, replaceOneCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer replaceOneCancel()
	readSetting.Patch(setting)
	settingBsonDocument := utilities.ToBsonDocument(readSetting)
	_, replaceOneQueryErr := collectionQuery.ReplaceOne(replaceOne, nil, settingBsonDocument)
	if replaceOneQueryErr != nil {
		err = replaceOneQueryErr
	}

	return nil
}

func (settingDatabaseRepository *SettingDatabaseRepository) CreateOne(setting *value_objects.Setting) (err error) {
	database := settingDatabaseRepository.OneDatabaseDatastore.Configuration.Database
	databaseQuery := settingDatabaseRepository.OneDatabaseDatastore.Client.Database(database)
	collectionQuery := databaseQuery.Collection("setting")

	insertOneCtx, insertOneCtxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer insertOneCtxCancel()
	settingBsonDocument := utilities.ToBsonDocument(setting)
	_, insertOneQueryErr := collectionQuery.InsertOne(insertOneCtx, settingBsonDocument)
	if insertOneQueryErr != nil {
		err = insertOneQueryErr
	}

	return err
}
