package helpers

func IsUpdateRequired(newModel interface{}, oldModel interface{}) interface{} {
	if newModel == oldModel {
		return oldModel
	}
	return newModel
}
