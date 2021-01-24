package errors

// Equals - check if two errors are the same
func Equals(err1 error, err2 error) bool {
	if err1 == err2 {
		return true
	}

	if err1 != nil && err2 != nil {
		return err1.Error() == err2.Error()
	}

	return false
}

// EqualsExtended - check if two extended errors are the same
func EqualsExtended(err1 Error, err2 Error) bool {
	if !err1.Occur || !err2.Occur {
		return false
	}

	if (err1.Code == err2.Code) && (err1.Code != ErrCodeNone) {
		return true
	}

	return err1.String() == err2.String()
}
