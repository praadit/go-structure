package constants

const (
	SUCCESS = "Success"

	// System Error Code
	ERR_SYSTEM_UnknownError     = "#0001" // Error yg gk ke recover
	ERR_SYSTEM_NotFound         = "#0002" // Error endpoint Not Found
	ERR_SYSTEM_Unimplemented    = "#0003" // Error endpoint Not Found
	ERR_SYSTEM_InvalidConfig    = "#0004" // Error endpoint Not Found
	ERR_SYSTEM_UnderMaintenance = "#0005"

	// Request Error Code
	ERR_REQ_RequestNotValid     = "#1001" // Error ketika request gak valid di validasi
	ERR_REQ_FailedToBindRequest = "#1002" // Error ketika request data gk bener
	ERR_REQ_AutenticationFailed = "#1003"
	ERR_REQ_AccessForbidden     = "#1004"
	ERR_REQ_RecaptchaFailed     = "#1005"
	ERR_REQ_DoubleRequest       = "#1006"

	// Database Error Code
	ERR_DB_FailedToGetData    = "#2001"
	ERR_DB_FailedToSaveData   = "#2002"
	ERR_DB_FailedToDeleteData = "#2003"
	ERR_DB_NoDataChanges      = "#2004"
	ERR_DB_DataNotFound       = "#2005"
	ERR_DB_DuplicateUnique    = "#2006"
	ERR_DB_InvalidRelation    = "#2007"

	// Expected Error
	ERR_EXP_DataStateAlreadyVerified      = "#3000" // ketika praktisi ingin publish data yang udah verified
	ERR_EXP_DataStateAlreadyPublished     = "#3001" // ketika praktisi ingin mengubah/publish data yang state nya bukan 1 atau 2
	ERR_EXP_DataStateAlreadyChanged       = "#3002" // ketika data state praktisi sama dengan request
	ERR_EXP_DataStateAlreadyRejected      = "#3003" // ketika praktisi ingin publish data yang rejected
	ERR_EXP_DataStateNotVerified          = "#3004" // ketika praktisi ingin update data yang belum verified
	ERR_EXP_DataStateNotPublished         = "#3005"
	ERR_EXP_DataStateAlreadyResponded     = "#3006"
	ERR_EXP_EmailAlreadyRegistered        = "#3100" // requested email has been registered for another account
	ERR_EXP_EmailNotVerified              = "#3101" // email is not verified yet
	ERR_EXP_VerifyEmailTokenInvalid       = "#3201" // when verify email token is invalid
	ERR_EXP_ConnectionRequestAlreadyExist = "#3301" // when verify email token is invalid
	ERR_EXP_KeahlianAlreadyExist          = "#3401" // when requested keahlian already exist

	ERR_PDDIKTI_ReturnBadRequest    = "#4001"
	ERR_PDDIKTI_ReturnNotFound      = "#4002"
	ERR_PDDIKTI_ReturnInternalError = "#4003"
	ERR_PDDIKTI_ReturnBadGateway    = "#4004"
	ERR_PDDIKTI_ReturnTimeout       = "#4005"
	ERR_PDDIKTI_NIDNNotValid        = "#4101"
)
