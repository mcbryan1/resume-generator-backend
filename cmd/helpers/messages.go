package helpers

const (
	// Response codes
	SuccessRespCode             = "000"
	ErrorRespCode               = "001"
	NotFoundRespCode            = "002"
	InvalidInputRespCode        = "003"
	UnauthorizedRespCode        = "004"
	ForbiddenRespCode           = "005"
	ConflictRespCode            = "006"
	InternalServerErrorRespCode = "007"
	ServiceUnavailableRespCode  = "008"
	GatewayTimeoutRespCode      = "009"

	// Response messages
	SuccessRespDesc                = "Success"
	ErrorRespDesc                  = "Error"
	NotFoundRespDesc               = "Not Found"
	InvalidInputRespDesc           = "Invalid Input"
	UnauthorizedRespDesc           = "Unauthorized"
	ForbiddenRespDesc              = "Forbidden"
	ConflictRespDesc               = "Conflict"
	InternalServerErrorRespDesc    = "Internal Server Error"
	InvalidRequestDataRespDesc     = "Invalid request data"
	InvalidPhoneOrPasswordRespDesc = "Invalid phone number or password"
	InvalidTokenRespDesc           = "Invalid token"
	ExpiredTokenRespDesc           = "Expired token"
	InvalidCredentialsRespDesc     = "Invalid credentials"
	InvalidTokenSignatureRespDesc  = "Invalid token signature"
	ValidationErrorRespDesc        = "Validation error"
)
