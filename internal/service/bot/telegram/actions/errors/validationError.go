package errors

// ValidationError Ошибка валидации данных экшена
type ValidationError struct {
	Reason string
}

func (err ValidationError) Error() string {
	return err.Reason
}
