package repository

import (
	"database/sql"
	"log"

	"github.com/go-squads/comet-backend/domain"
	"github.com/go-squads/comet-backend/appcontext"
)

var err error

type ConfigRepository struct{
	db *sql.DB
}

func (self ConfigRepository) GetConfiguration(appName string, namespaceName string) []domain.Configuration{
	var cfg []domain.Configuration
	var activeVersion int
	var applicationId int
	var namespaceId int

	_ = self.db.QueryRow("SELECT id FROM application WHERE name=$1",appName).Scan(&applicationId)
	_ = self.db.QueryRow("SELECT id,active_version FROM namespace WHERE app_id=$1 AND name=$2",applicationId,namespaceName).Scan(&namespaceId,&activeVersion)
	row, err := self.db.Query("SELECT key,value FROM configuration WHERE version=$1 AND namespace_id=$2",activeVersion,namespaceId)

	if err != nil{
		log.Fatalf(err.Error())
	}

	for row.Next(){
		var key string
		var value string

		err = row.Scan(&key,&value)
		cfg = append(cfg, domain.Configuration{namespaceId,activeVersion,key,value})
	}	
	return cfg
}

func NewConfigurationRepository() ConfigRepository{
	return ConfigRepository{
		db: appcontext.GetDB(),
	}
}
