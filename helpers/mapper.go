package helpers

import "encoding/json"

//TODO use dynamic mapping

//Mapper can map data
func Mapper(object interface{}, model interface{}) (interface{}, error) {
	data, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, model); err != nil {
		return nil, err
	}
	return model, nil
}
