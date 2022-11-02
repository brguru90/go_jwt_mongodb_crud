package my_modules

import (
	"learn_go/src/database"
	"time"
)

// Usage:
// 
// var test_session interface{}=nil
// 
// payload.GetSession(&test_session)
func (access_token AccessToken) GetSession(pointer_to_destination interface{}) error{
	return database.RedisPoolGetJSON(access_token.Token_id+"_session",pointer_to_destination)
}

// Usage:
// 
// payload.SetSession("test")
func (access_token AccessToken) SetSession(pointer_to_source interface{}) error{
	return database.RedisPoolSetJSON(access_token.Token_id+"_session",pointer_to_source,time.Duration(time.Now().UnixMilli())-time.Duration(access_token.Exp))
}
