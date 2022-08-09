package utils

import (
	"encoding/json"
)

func StructToJsonBytes(param interface{}) []byte {
	dataType, _ := json.Marshal(param)
	return dataType
}

func StructToJsonString(param interface{}) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}
func ByteToStruct(message []byte, abstract interface{}) (interface{}, error) {
	//msg := common_struct.ServerProviderAskMessage{}
	err := json.Unmarshal(message, &abstract)
	return abstract, err
}
