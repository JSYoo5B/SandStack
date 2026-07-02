package identity

import "errors"

var ErrUserNotFound = errors.New("user not found")

var ErrProjectNotFound = errors.New("project not found")

var ErrRoleNotFound = errors.New("role not found")

var ErrTokenNotFound = errors.New("token not found")

var ErrServiceNotFound = errors.New("service not found")

var ErrEndpointNotFound = errors.New("endpoint not found")
