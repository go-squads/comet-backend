package repository

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/go-squads/comet-backend/appcontext"
	"github.com/go-squads/comet-backend/domain"
)

var err error

type ConfigRepository struct {
	db *sql.DB
}

const (
	getAppNameQuery                     = "SELECT id FROM application WHERE name=$1"
	getNamespaceIdAndActiveVersionQuery = "SELECT id,active_version FROM namespace WHERE app_id=$1 AND name=$2"
	getConfigurationKeyValueQuery       = "SELECT key,value FROM configuration WHERE version=$1 AND namespace_id=$2"
)

func (self ConfigRepository) GetConfiguration(appName string, namespaceName string, version string) []domain.Configuration {
	var cfg []domain.Configuration
	var activeVersion int
	var chosenVersion int
	var applicationId int
	var namespaceId int
	var rows *sql.Rows

	_ = self.db.QueryRow(getAppNameQuery, appName).Scan(&applicationId)
	_ = self.db.QueryRow(getNamespaceIdAndActiveVersionQuery, applicationId, namespaceName).Scan(&namespaceId, &activeVersion)

	if version != "" {
		versionInt, _ := strconv.Atoi(version)
		chosenVersion = versionInt
	} else {
		chosenVersion = activeVersion
	}

	rows, err = self.db.Query(getConfigurationKeyValueQuery, chosenVersion, namespaceId)

	if err != nil {
		log.Fatalf(err.Error())
	}

	for rows.Next() {
		var key string
		var value string

		err = rows.Scan(&key, &value)
		cfg = append(cfg, domain.Configuration{namespaceId, chosenVersion, key, value})
	}
	return cfg
}

func NewConfigurationRepository() ConfigRepository {
	return ConfigRepository{
		db: appcontext.GetDB(),
	}
}
