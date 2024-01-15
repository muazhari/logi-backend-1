package managements

import (
	"github.com/muazhari/logi-backend-1/src/inners/models/value_objects"
	databaseRepositories "github.com/muazhari/logi-backend-1/src/outers/repositories/databases"
)

type SettingManagement struct {
	SettingDatabaseRepository *databaseRepositories.SettingDatabaseRepository
}

func NewSettingManagement(settingRepository *databaseRepositories.SettingDatabaseRepository) *SettingManagement {
	settingManagement := &SettingManagement{
		SettingDatabaseRepository: settingRepository,
	}
	return settingManagement
}

func (settingManagement *SettingManagement) ReadOne() (output *value_objects.Setting, err error) {

	readOne, readOneErr := settingManagement.SettingDatabaseRepository.ReadOne()
	if readOneErr != nil {
		err = readOneErr
	}

	output = readOne

	return output, err
}

func (SettingManagement *SettingManagement) PatchOne(setting *value_objects.Setting) (err error) {

	patchOneErr := SettingManagement.SettingDatabaseRepository.PatchOne(setting)
	if patchOneErr != nil {
		err = patchOneErr
	}

	return err
}
