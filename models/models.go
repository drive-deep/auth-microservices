package models

// GetAllModels returns a list of all models in the application
func GetAllModels() []interface{} {
	return []interface{}{
		(*User)(nil), // Add other models here as needed
	}
}
