package model

//import (
	//"time"
	//"log"
	//"sso/data"
//)

type NewGPIUser struct {
	Lastname        string  `json:"lastname,omitempty"`
	Firstname string  `json:"firstname,omitempty"`
	Dob       string `json:"dob,omitempty"`
	Sex       string `json:"sex",omitempty`
	Email     string  `json:"email,omitempty"`
	Address string  `json:"address,omitempty"`
	City       string `json:"city,omitempty"`
	Username       string `json:"username,omitempty"`
	Password	string `json:"password,omitempty"`
	RPassword	string `json:"rpassword,omitempty"`
}

type GPIUser struct {
	ID	int `json:"id,omitempty"`
	Lastname        string  `json:"lname,omitempty"`
	Firstname string  `json:"fname,omitempty"`
	Dob       string `json:"dob,omitempty"`
	Sex       string `json:"gender,omitempty"`
	Email     string  `json:"myemail,omitempty"`
	Address string  `json:"address,omitempty"`
	City       string `json:"city,omitempty"`
	Username       string `json:"username,omitempty"`
	Password	string `json:"password,omitempty"`
}


// func (thisnewuser *NewGPIUser) CountServiceSubscribedTo() int {
// 	//defer Conn.Close()
// 	rows, _ := Conn.Query("select id, sname, description, dclass, iclass from allservices")

// 	var ServiceItems []NewGPIUser
// 	var ServiceItem NewGPIUser


// 	for rows.Next() {
// 		_ = rows.Scan(&ServiceItem.ID, &ServiceItem.Name, &ServiceItem.Description, &ServiceItem.DClass, &ServiceItem.IClass)

// 		ServiceItems = append(ServiceItems, ServiceItem)
// 	}

// 	return ServiceItems, nil
// }

// func (thisnewuser *NewGPIUser) SubscribeToServiceForUser(userid) bool {
// 	//defer Conn.Close()
// 	rows, _ := Conn.Query("select id, sname, description, dclass, iclass from allservices")

// 	var ServiceItems []NewGPIUser
// 	var ServiceItem NewGPIUser


// 	for rows.Next() {
// 		_ = rows.Scan(&ServiceItem.ID, &ServiceItem.Name, &ServiceItem.Description, &ServiceItem.DClass, &ServiceItem.IClass)

// 		ServiceItems = append(ServiceItems, ServiceItem)
// 	}

// 	return ServiceItems, nil
// }

// func (thisnewuser *NewGPIUser) UnsubscribeServiceForUser(userid) bool {
// 	//defer Conn.Close()
// 	rows, _ := Conn.Query("select id, sname, description, dclass, iclass from allservices")

// 	var ServiceItems []NewGPIUser
// 	var ServiceItem NewGPIUser


// 	for rows.Next() {
// 		_ = rows.Scan(&ServiceItem.ID, &ServiceItem.Name, &ServiceItem.Description, &ServiceItem.DClass, &ServiceItem.IClass)

// 		ServiceItems = append(ServiceItems, ServiceItem)
// 	}

// 	return ServiceItems, nil
// }

// type OAuthClient struct {
// 	ClientId     string
// 	ClientSecret string
// 	RedirectUri  string
// 	Scope        string
// }

// type OAuthToken struct {
// 	Token        string
// 	Expiration   string
// 	ClientId     string
// 	Username     string
// 	TokenType    string
// 	TokenScope   string
// 	RefreshToken string
// }

// func (token *OAuthToken) ReturnTokenExpirationInSeconds() int {
// 	log.Println("token expired date:" + token.Expiration)
// 	expTime, err := time.ParseInLocation("2018-01-02 15:04:05", token.Expiration, time.Local)
// 	if err != nil {
// 		log.Println("error while parse time")
// 		log.Println(err)
// 		return 0
// 	}
// 	eu := expTime.Unix()
// 	nu := time.Now().Unix()

// 	gap := eu - nu
// 	if gap <= 0 {
// 		return 0
// 	}
// 	return int(gap)
// }

// func (token *OAuthToken) IsTokenExpirated() bool {
// 	gap := token.ReturnTokenExpirationInSeconds()
// 	if gap == 0 {
// 		return true
// 	}
// 	return false
// }

// type TokenToReturn struct {
// 	Access_token  string        `json:"access_token"`
// 	Token_type    string        `json:"token_type"`
// 	Expires_in    int        `json:"expires_in"`
// 	Scope         string        `json:"scope"`
// 	User          string        `json:"user"`
// 	Refresh_Token string                `json:"refresh_token"`
// }
