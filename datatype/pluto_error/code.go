package pluto_error

import "net/http"

var (
	ServerError = NewPlutoError(http.StatusInternalServerError, 1000, "Server Error", nil)
	BadRequest  = NewPlutoError(http.StatusBadRequest, 1001, "Bad Request", nil)
	Forbidden   = NewPlutoError(http.StatusForbidden, 1002, "Permission Forbidden", nil)

	MailIsAlreadyRegister = NewPlutoError(http.StatusForbidden, 2001, "Mail is already been registered", nil)
	MailNotExist          = NewPlutoError(http.StatusForbidden, 2002, "Mail does not exist", nil)
	MailIsNotVerified     = NewPlutoError(http.StatusForbidden, 2003, "Mail is not verified", nil)
	MailAlreadyVerified   = NewPlutoError(http.StatusBadRequest, 2004, "Mail is already verified", nil)

	InvalidPassword      = NewPlutoError(http.StatusForbidden, 3001, "Invalid Password", nil)
	InvalidRefreshToken  = NewPlutoError(http.StatusForbidden, 3002, "Invalid Refresh Token", nil)
	InvalidJWTToekn      = NewPlutoError(http.StatusForbidden, 3003, "Invalid JWT Token", nil)
	InvalidGoogleIDToken = NewPlutoError(http.StatusForbidden, 3004, "Invalid Google ID Token", nil)
	InvalidWechatCode    = NewPlutoError(http.StatusForbidden, 3005, "Invalid Wechat Code", nil)
	InvalidAvatarFormat  = NewPlutoError(http.StatusBadRequest, 3006, "Invalid Avatar Format", nil)
	InvalidAppleIDToken  = NewPlutoError(http.StatusForbidden, 3007, "Invalid Apple ID Token", nil)

	JWTTokenExpired = NewPlutoError(http.StatusForbidden, 3008, "JWT Token Expired", nil)

	ScopeExists         = NewPlutoError(http.StatusForbidden, 4001, "Scope already exists", nil)
	ScopeNotExist       = NewPlutoError(http.StatusForbidden, 4002, "Scope not exist", nil)
	ScopeAttached       = NewPlutoError(http.StatusForbidden, 4003, "Scope already attached", nil)
	ApplicationNotExist = NewPlutoError(http.StatusForbidden, 4004, "Application does not exist", nil)
	RoleNotExist        = NewPlutoError(http.StatusForbidden, 4005, "Role does not exist", nil)
	NotPlutoAdmin       = NewPlutoError(http.StatusForbidden, 4006, "Not the pluto admin", nil)
)
