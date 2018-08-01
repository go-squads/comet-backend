package repository

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
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
	getNamespaceIdQuery                 = "SELECT id FROM namespace WHERE app_id = $1 AND name = $2"
	getNamespacesIdOnlyQuery            = "SELECT app_id FROM namespace GROUP BY app_id"
	getNamespaceIdAndActiveVersionQuery = "SELECT id, active_version FROM namespace WHERE app_id = $1 AND name = $2"
	getNamespaceIdAndLatestVersionQuery = "SELECT id, latest_version FROM namespace WHERE app_id = $1 AND name = $2"
	getLatestVersionNamespaceQuery      = "SELECT latest_version FROM namespace WHERE app_id = $1 AND name = $2"
	getConfigurationKeyValueQuery       = "SELECT key,value FROM configuration WHERE version = $1 AND namespace_id = $2"

	insertNewConfigurationQuery          = "INSERT INTO configuration VALUES ($1, $2, $3, $4)"                                                                                                       // namespace_id, version, key, value
	insertHistoryQuery                   = "INSERT INTO history (user_id, namespace_id, predecessor_version, successor_version, created_at) VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP) RETURNING id" // user_id, namespace_id, predecessor_version, successor version
	insertConfigurationChangesQuery      = "INSERT INTO configuration_change VALUES ($1, $2, $3)"                                                                                                    // history_id, key, new_value
	incrementNamespaceActiveVersionQuery = "UPDATE namespace SET active_version = $1, latest_version = $1 WHERE id = $2"
	showHistoryQuery                     = "SELECT u.username,n.name,predecessor_version,successor_version,key,new_value,h.created_at FROM history AS h INNER JOIN configuration_change as cfg ON h.id=cfg.history_id INNER JOIN namespace AS n ON h.namespace_id = n.id INNER JOIN users AS u ON h.user_id = u.id WHERE n.id = $1"
	fetchNamespaceQuery                  = "SELECT name FROM namespace WHERE app_id = $1"
	getListOfApplicationNamespaceQuery   = "SELECT app.name, app.id FROM application AS app INNER JOIN namespace AS n ON app.id = n.id"
	updateVersionBasedApplicationQuery   = "UPDATE namespace SET active_version = $1 WHERE app_id = $2 and name = $3"
)

func (self ConfigRepository) GetConfiguration(appName string, namespaceName string, version string) domain.ApplicationConfiguration {
	var appConfig domain.ApplicationConfiguration
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
		cfg = append(cfg, domain.Configuration{Key: key, Value: value})
	}
	appConfig = domain.ApplicationConfiguration{NamespaceID: namespaceId, Version: chosenVersion, Configurations: cfg}
	return appConfig
}

func (self ConfigRepository) InsertConfiguration(newConfigs domain.ConfigurationRequest) domain.Response {
	var latestVersion int
	var activeVersion int
	var newVersion int
	var applicationId int
	var historyId int
	var namespaceId int

	err = self.db.QueryRow(getAppIdQuery, newConfigs.AppName).Scan(&applicationId)
	if err != nil {
		return domain.FailedResponse(err)
	}

	err = self.db.QueryRow(getNamespaceIdAndActiveVersionQuery, applicationId, newConfigs.Namespace).Scan(&namespaceId, &activeVersion)
	if err != nil {
		return domain.FailedResponse(err)
	}

	err = self.db.QueryRow(getNamespaceIdAndLatestVersionQuery, applicationId, newConfigs.Namespace).Scan(&namespaceId, &latestVersion)
	if err != nil {
		return domain.FailedResponse(err)
	}

	newVersion = latestVersion + 1

	err = self.db.QueryRow(insertHistoryQuery, 1, namespaceId, activeVersion, newVersion).Scan(&historyId)
	if err != nil {
		return domain.FailedResponse(err)
	}

	for _, config := range newConfigs.Data {
		key := config.Key
		value := config.Value

		_, err = self.db.Exec(insertNewConfigurationQuery, namespaceId, newVersion, key, value)
		if err != nil {
			return domain.FailedResponse(err)
		}

		_, err = self.db.Exec(insertConfigurationChangesQuery, historyId, key, value)
		if err != nil {
			return domain.FailedResponse(err)
		}
	}

	_, err := self.db.Exec(incrementNamespaceActiveVersionQuery, newVersion, namespaceId)
	if err != nil {
		return domain.FailedResponse(err)
	} else {
		return domain.SuccessResponse()
	}
}

func (self ConfigRepository) ReadHistory(appName string, namespace string) []domain.ConfigurationHistory {
	var history []domain.ConfigurationHistory
	var applicationId int
	var namespaceId int
	var rows *sql.Rows

	err = self.db.QueryRow(getAppIdQuery, appName).Scan(&applicationId)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = self.db.QueryRow(getNamespaceIdQuery, applicationId, namespace).Scan(&namespaceId)
	if err != nil {
		log.Fatalf(err.Error())
	}

	rows, err = self.db.Query(showHistoryQuery, namespaceId)
	if err != nil {
		log.Fatalf(err.Error())
	}

	for rows.Next() {
		var username string
		var namespace string
		var predecessorVersion int
		var successorVersion int
		var key string
		var value string
		var createdTime string

		err = rows.Scan(&username, &namespace, &predecessorVersion, &successorVersion, &key, &value, &createdTime)
		history = append(history, domain.ConfigurationHistory{Username: username, Namespace: namespace, PredecessorVersion: predecessorVersion, SuccessorVersion: successorVersion, Key: key, Value: value, CreatedAt: createdTime})
	}
	return history
}

func (self ConfigRepository) GetListOfNamespace(applicationId int) []string {
	var list []string
	var row *sql.Rows

	row, err = self.db.Query(fetchNamespaceQuery, applicationId)
	if err != nil {
		log.Fatalf(err.Error())
	}

	for row.Next() {
		var name string

		err = row.Scan(&name)
		list = append(list, name)
	}
	return list
}

func (self ConfigRepository) getListApplicationId() []int {
	var lsNamespaceId []int
	var rows *sql.Rows

	rows, err := self.db.Query(getNamespacesIdOnlyQuery)
	if err != nil {
		log.Fatalf(err.Error())
	}

	for rows.Next() {
		var applicationId int

		err = rows.Scan(&applicationId)
		lsNamespaceId = append(lsNamespaceId, applicationId)
	}
	fmt.Println(lsNamespaceId)

	return lsNamespaceId
}

func (self ConfigRepository) GetApplicationNamespace() []domain.ApplicationNamespace {
	var lsApplication []domain.ApplicationNamespace

	rows, err := self.db.Query(getListOfApplicationNamespaceQuery)
	if err != nil {
		log.Fatalf(err.Error())
	}

	for rows.Next() {
		var applicationName string
		var applicationId int

		err = rows.Scan(&applicationName, &applicationId)
		lsApplication = append(lsApplication, domain.ApplicationNamespace{ApplicationName: applicationName, Namespace: self.GetListOfNamespace(applicationId)})
	}

	return lsApplication
}

func (self ConfigRepository) RollbackVersionNamespace(rollback domain.ConfigurationRollback) domain.Response {
	var activeVersion int
	var latestVersion int
	var applicationId int
	var namespaceId int
	var historyId int

	_ = self.db.QueryRow(getAppIdQuery, rollback.Appname).Scan(&applicationId)
	fmt.Println(rollback.Appname)
	_ = self.db.QueryRow(getNamespaceIdAndActiveVersionQuery, applicationId, rollback.NamespaceName).Scan(&namespaceId, &activeVersion)

	fmt.Println(rollback.NamespaceName)

	_ = self.db.QueryRow(getLatestVersionNamespaceQuery, applicationId, rollback.NamespaceName).Scan(&latestVersion)

	if rollback.Version > latestVersion {
		return domain.Response{Status: http.StatusBadRequest, Message: "Invalid version request"}
	}

	_, err = self.db.Exec(updateVersionBasedApplicationQuery, rollback.Version, applicationId, rollback.NamespaceName) //version, app_id, namespace_name

	if err != nil {
		log.Fatalf(err.Error())
	}

	if err != nil {
		return domain.FailedResponse(err)
	} else {
		err = self.db.QueryRow(insertHistoryQuery, 1, namespaceId, activeVersion, rollback.Version).Scan(&historyId)
		if err != nil {
			return domain.FailedResponse(err)
		}
		return domain.Response{Status: http.StatusOK, Message: "Updated"}
	}

}

func NewConfigurationRepository() ConfigRepository {
	return ConfigRepository{
		db: appcontext.GetDB(),
	}
}
