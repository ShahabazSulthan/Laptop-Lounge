package resCustomError

import (
	"errors"
)

var (
	IDParamsEmpty            = errors.New("ID parameter is empty")
	BindingConflict          = errors.New("data criteria not met or conflicting data")
	NotGetSellerIDInContext  = errors.New("failed to retrieve seller ID from context")
	NotGetUserIDInContext    = errors.New("failed to retrieve user ID from context")
	ErrConversionOfPage      = errors.New("error converting string to integer for page parameter")
	ErrConversionOfLimit     = errors.New("error converting string to integer for page limit parameter")
	ErrPagination            = errors.New("page must start from one")
	ErrPageLimit             = errors.New("page limit must be greater than one")
	ErrNegativeID            = errors.New("ID must be greater than one")
	ErrNoRowAffected         = errors.New("no data matching the specified criteria found in the database")
	ErrProductOrderCompleted = errors.New("product order is already completed")
	ErrAdminDashboard        = errors.New("encountered an issue in the admin dashboard")
)
