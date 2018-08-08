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

const (
	checkRoleBaseQuery = "SELECT role FROM users where token = $1"
	setUserRole        = "SET ROLE $1"
	ADMIN              = "admin"
	CLIENT             = "client"
)

func (self ApplicationRepository) setRoleBased(token string) {
	result, err := self.db.Exec("SET ROLE " + self.getUserRoleBased(token))
	if err != nil {
		log.Println(result)
	}
}

func (self ApplicationRepository) getUserRoleBased(token string) string {
	var userRole string
	err := self.db.QueryRow(checkRoleBaseQuery, token).Scan(&userRole)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(userRole)
	return userRole
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

func (self ApplicationRepository) CreateApplication(newApp domain.CreateApplication, token string) domain.Response {
	self.setRoleBased(token)

	if self.validateApplicationName(newApp.ApplicationsName) == false {
		return domain.Response{Status: http.StatusBadRequest, Message: "Duplicate Application Name"}
	} else {
		var applicationId int
		_, err = self.db.Query(createNewApplicationQuery, newApp.ApplicationsName)
		if err != nil {
			fmt.Println(err.Error() + " inserted application")
			return domain.Response{Status: http.StatusForbidden, Message: "Action Forbidden"}
		} else {

			fmt.Print(applicationId)
			return domain.Response{Status: http.StatusOK, Message: "Inserted New Application"}
		}
	}

}

func (self ApplicationRepository) CreateNewNamespace(appName string,token string, namespaceName domain.Namespace) domain.Response {
	var applicationId int

	self.setRoleBased(token)

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
			if err != nil {
				return domain.Response{Status: http.StatusForbidden, Message: "Action Forbidden"}
			}
			return domain.Response{Status: http.StatusOK, Message: "New Namespace Created"}
		}
	}
}

func (self ApplicationRepository) GetListOfNamespace(applicationId int, token string) []string {
	var list []string
	var row *sql.Rows

	self.setRoleBased(token)

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

func (self ApplicationRepository) GetApplicationNamespace(token string) []domain.ApplicationNamespace {
	var lsApplication []domain.ApplicationNamespace

	rows, err := self.db.Query(getListOfApplicationNamespaceQuery)
	if err != nil {
		log.Fatalf(err.Error())
	}

	for rows.Next() {
		var applicationName string
		var applicationId int

		err = rows.Scan(&applicationName, &applicationId)
		lsApplication = append(lsApplication, domain.ApplicationNamespace{ApplicationName: applicationName, Namespace: self.GetListOfNamespace(applicationId,token)})
	}

	fmt.Println(lsApplication)
	return lsApplication
}

func NewApplicationRepository() ApplicationRepository {
	return ApplicationRepository{
		db: appcontext.GetDB(),
	}
}
