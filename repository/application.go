package repository

import (
	"database/sql"
	"fmt"
	"github.com/go-squads/comet-backend/appcontext"
	"github.com/go-squads/comet-backend/domain"
	"log"
	"net/http"
)

type ApplicationRepository struct {
	db *sql.DB
}

func (self ApplicationRepository) validateApplicationName(appName string) bool {
	var rows *sql.Rows
	isAvailable := true

	rows, err = self.db.Query(selectApplicationName)

	for rows.Next() {
		var applicationName string
		err = rows.Scan(&applicationName)

		if applicationName == appName {
			isAvailable = false
		} else {
			isAvailable = true
		}
	}
	fmt.Println(isAvailable)
	return isAvailable
}

func (self ApplicationRepository) validateNamespaceName(namespaceName string, applicationId int) bool {
	var rows *sql.Rows
	namespaceAvailable := true

	rows, err = self.db.Query(checkNamespaceQuery, applicationId)

	for rows.Next() {
		var name string

		err = rows.Scan(&name)
		if name == namespaceName {
			namespaceAvailable = false
		} else {
			namespaceAvailable = true
		}
	}

	return namespaceAvailable
}

func (self ApplicationRepository) CreateApplication(newApp domain.CreateApplication) domain.Response {

	if self.validateApplicationName(newApp.ApplicationsName) == false {
		return domain.Response{Status: http.StatusBadRequest, Message: "Duplicate Application Name"}
	} else {
		var applicationId int
		_, err = self.db.Query(createNewApplicationQuery, newApp.ApplicationsName)
		if err != nil {
			log.Fatalf(err.Error())
		}

		fmt.Print(applicationId)
		return domain.Response{Status: http.StatusOK, Message: "Inserted New Application"}
	}

}

func (self ApplicationRepository) CreateNewNamespace(appName string, namespaceName domain.Namespace) domain.Response {
	var applicationId int
	if appName == "" {
		return domain.Response{Status: http.StatusBadRequest, Message: "App name null"}
	} else if namespaceName.Name == "" {
		return domain.Response{Status: http.StatusBadRequest, Message: "Namespace name null"}
	} else if appName == "" && namespaceName.Name == "" {
		return domain.Response{Status: http.StatusBadRequest, Message: "Namespace name and App name null"}
	} else {
		err = self.db.QueryRow(getAppIdQuery, appName).Scan(&applicationId)
		if err != nil {
			log.Fatalf(err.Error())
		}

		if self.validateNamespaceName(namespaceName.Name, applicationId) == false {
			return domain.Response{Status: http.StatusBadRequest, Message: "Namespace already taken"}
		} else {
			_, err = self.db.Query(insertNewNamespaceQuery, namespaceName.Name, applicationId, 1, 1)
			return domain.Response{Status: http.StatusOK, Message: "New Namespace Created"}
		}
	}
}

func NewApplicationRepository() ApplicationRepository {
	return ApplicationRepository{
		db: appcontext.GetDB(),
	}
}
