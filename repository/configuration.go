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
	getAppIdQuery                       = "SELECT id FROM application WHERE name = $1"
	getNamespaceIdAndActiveVersionQuery = "SELECT id, active_version FROM namespace WHERE app_id = $1 AND name = $2"
	getNamespaceIdAndLatestVersionQuery = "SELECT id, latest_version FROM namespace WHERE app_id = $1 AND name = $2"
	getConfigurationKeyValueQuery       = "SELECT key,value FROM configuration WHERE version = $1 AND namespace_id = $2"

	insertNewConfigurationQuery          = "INSERT INTO configuration VALUES ($1, $2, $3, $4)"                           // namespace_id, version, key, value
	insertHistoryQuery                   = "INSERT INTO history (user_id, namespace_id, predecessor_version, successor_version, created_at) VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP) RETURNING id" // user_id, namespace_id, predecessor_version, successor version
	insertConfigurationChangesQuery      = "INSERT INTO configuration_change VALUES ($1, $2, $3)"                        // history_id, key, new_value
	incrementNamespaceActiveVersionQuery = "UPDATE namespace SET active_version = $1, latest_version = $1 WHERE id = $2"
)

func (self ConfigRepository) GetConfiguration(appName string, namespaceName string, version string) []domain.Configuration {
	var cfg []domain.Configuration
	var activeVersion int
	var chosenVersion int
	var applicationId int
	var namespaceId int
	var rows *sql.Rows

	_ = self.db.QueryRow(getAppIdQuery, appName).Scan(&applicationId)
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
		cfg = append(cfg, domain.Configuration{NamespaceID: namespaceId, Version: chosenVersion, Key: key, Value: value})
	}
	return cfg
}

func (self ConfigRepository) InsertConfiguration(newConfigs domain.ConfigurationRequest) {
	var latestVersion int
	var activeVersion int
	var newVersion int
	var applicationId int
	var historyId int
	var namespaceId int

	_ = self.db.QueryRow(getAppIdQuery, newConfigs.AppName).Scan(&applicationId)
	_ = self.db.QueryRow(getNamespaceIdAndActiveVersionQuery, applicationId, newConfigs.Namespace).Scan(&namespaceId, &activeVersion)
	_ = self.db.QueryRow(getNamespaceIdAndLatestVersionQuery, applicationId, newConfigs.Namespace).Scan(&namespaceId, &latestVersion)

	newVersion = latestVersion + 1

	_ = self.db.QueryRow(insertHistoryQuery, 1, namespaceId, activeVersion, newVersion).Scan(&historyId)

	for _, config := range newConfigs.Data {
		key := config.Key
		value := config.Value
		self.db.Exec(insertNewConfigurationQuery, namespaceId, newVersion, key, value)
		self.db.Exec(insertConfigurationChangesQuery, historyId, key, value)
	}

	self.db.Exec(incrementNamespaceActiveVersionQuery, newVersion, namespaceId)
}

func NewConfigurationRepository() ConfigRepository {
	return ConfigRepository{
		db: appcontext.GetDB(),
	}
}
