package constant

import "errors"

var ErrInsertDatabase error = errors.New("invalid add data in database")
var ErrEmptyInput error = errors.New("wishlist title, is_achieved cannot be empty")
var ErrUserExist error = errors.New("user already exist")
var ErrUserNotExist error = errors.New("user not exist")
var ErrUserWrongInput error = errors.New("wrong input email or password")
var ErrInternalServer error = errors.New("internal server error")
var ErrUserCreate error = errors.New("failed to create user")
