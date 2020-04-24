package data

import (
	"sso/model"
	// "sso/util"
	// "strings"
	// "fmt"
	// "time"
)

func GetGPIService(serviceid string) *model.GPIService {
	//defer Conn.Close()
	rows, _ := Conn.Query("select id, sname, description, dclass, iclass from allservices limit 1")

	//var ServiceItems []GPIServices
	var ServiceItem model.GPIService

	//for rows.Next() {
	rows.Scan(&ServiceItem.ID, &ServiceItem.Name, &ServiceItem.Description, &ServiceItem.DClass, &ServiceItem.IClass)

		//ServiceItems = append(ServiceItems, ServiceItem)
	//}

	var ServiceItem_ = new(model.GPIService)

	ServiceItem_.ID = ServiceItem.ID
	ServiceItem_.Name = ServiceItem.Name
	ServiceItem_.Description = ServiceItem.Description
	ServiceItem_.DClass = ServiceItem.DClass
	ServiceItem_.IClass = ServiceItem.IClass

	return ServiceItem_

}

func GetSubscription(clientId string, user_id string) bool {
	//defer Conn.Close()
	row := Conn.QueryRow("SELECT id FROM GPI_Oauth_clients.subscription_history where user_id='"+user_id+"' and status = '1' and service_id in (select id from GPI_Oauth_clients.allservices where dclass= '"+clientId+"')");
	var client_id string
	// var token string
	// var expiration string
	// var username string
	// var token_type string
	// var token_scope string
	// var refresh_token string

	row.Scan(&client_id)

	if client_id != "" {
		
		return true
	}
	return false

}

// func GetAllGPIService() *[]model.GPIService {
// 	//defer Conn.Close()
// 	rows, _ := Conn.Query("select id, sname, description, dclass, iclass from allservices")

// 	var ServiceItems *[]model.GPIService
// 	var ServiceItem model.GPIService


// 	for rows.Next() {
// 		_ = rows.Scan(&ServiceItem.ID, &ServiceItem.Name, &ServiceItem.Description, &ServiceItem.DClass, &ServiceItem.IClass)

// 		ServiceItems = append(ServiceItems, ServiceItem)
// 	}

// 	return ServiceItems
// }