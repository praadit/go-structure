package services

// import (
// 	"encoding/json"
// 	"net/http"

// 	"gitlab.com/m8851/pmo-echo-api/config"
// 	"gitlab.com/m8851/pmo-echo-api/constants"
// )

// type RecaptchaServiceInterface interface {
// 	Confirm(captcha string) bool
// }

// type RecaptchaService struct {
// 	external BaseExternalServiceInterface
// }

// func NewRecaptchaService() RecaptchaServiceInterface {
// 	external := NewBaseExternalService()
// 	return &RecaptchaService{
// 		external: external,
// 	}
// }

// func (service *RecaptchaService) Confirm(captcha string) bool {
// 	secretKey := config.AppConfig[config.RecaptchaKey]
// 	req, _ := http.NewRequest(http.MethodPost, constants.Recaptcha_Verify, nil)

// 	query := req.URL.Query()
// 	query.Add("secret", secretKey)
// 	query.Add("response", captcha)
// 	req.URL.RawQuery = query.Encode()

// 	var googleResponse map[string]interface{}
// 	content := service.external.SendRequest(req)

// 	err := json.Unmarshal(content, &googleResponse)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return googleResponse["success"].(bool)
// }
