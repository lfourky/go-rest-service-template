package usecase

const (
	ErrTypeDatabase = "database"
	ErrTypeDomain   = "domain"
)

// Error is a wrapper for possible usecase errors.
type Error struct {
	Message string
	Code    int
	Type    string
}

// Error returns the error message.
func (e *Error) Error() string {
	return e.Message
}

// Domain error codes.
const (
	ErrCodeDomainInternal         = 5010000
	ErrCodeDomainUnableToSendMail = iota + ErrCodeDomainInternal
	ErrCodeDomainInvalidUUID
)

// Domain errors.
var (
	ErrDomainInternal = &Error{
		Message: "internal error",
		Code:    ErrCodeDomainInternal,
		Type:    ErrTypeDomain,
	}

	ErrDomainMailSendingFailed = &Error{
		Message: "unable to send mail",
		Code:    ErrCodeDomainUnableToSendMail,
		Type:    ErrTypeDomain,
	}

	ErrDomainInvalidUUID = &Error{
		Message: "invalid UUID",
		Code:    ErrCodeDomainInvalidUUID,
		Type:    ErrTypeDomain,
	}
)

// Database error codes.
const (
	ErrCodeDatabaseInternal     = 5020000
	ErrCodeDatabaseItemNotFound = iota + ErrCodeDatabaseInternal
	ErrCodeDatabaseItemCreationFailed
	ErrCodeDatabaseUserNotFound
	ErrCodeDatabaseUserAlreadyExists
)

// Database errors.
var (
	ErrDatabaseInternal = &Error{
		Message: "internal database error",
		Code:    ErrCodeDatabaseInternal,
		Type:    ErrTypeDatabase,
	}

	ErrDatabaseItemNotFound = &Error{
		Message: "item not found",
		Code:    ErrCodeDatabaseItemNotFound,
		Type:    ErrTypeDatabase,
	}

	ErrDatabaseItemCreationFailed = &Error{
		Message: "item creation failed",
		Code:    ErrCodeDatabaseItemCreationFailed,
		Type:    ErrTypeDatabase,
	}

	ErrDatabaseUserNotFound = &Error{
		Message: "user not found",
		Code:    ErrCodeDatabaseUserNotFound,
		Type:    ErrTypeDatabase,
	}

	ErrDatabaseUserAlreadyExists = &Error{
		Message: "user already exists",
		Code:    ErrCodeDatabaseUserAlreadyExists,
		Type:    ErrTypeDatabase,
	}
)
