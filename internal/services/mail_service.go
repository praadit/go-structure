package services

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"strings"

// 	"github.com/google/uuid"
// 	"github.com/sendgrid/rest"
// 	"github.com/sendgrid/sendgrid-go"
// 	"github.com/sendgrid/sendgrid-go/helpers/mail"
// 	"gitlab.com/m8851/pmo-echo-api/config"
// 	"gitlab.com/m8851/pmo-echo-api/constants"
// 	"gitlab.com/m8851/pmo-echo-api/handlers/dto"
// 	"gitlab.com/m8851/pmo-echo-api/utilities"
// )

// type MailServiceInterface interface {
// 	SendTemplateEmail(mailType constants.MailType, recipients []dto.RecipientDetail, templateID string, subject string, content string)
// 	GenerateRecipients(mailType constants.MailType, detail *dto.MailSchedulerDetailReq, redirectTarget string) (recipient dto.RecipientDetail)
// }

// type MailService struct {
// }

// func NewMailService() MailServiceInterface {
// 	return &MailService{}
// }

// func (service *MailService) initSendgridRequest() *mail.SGMailV3 {
// 	instance := mail.NewV3Mail()
// 	instance.SetFrom(mail.NewEmail(config.AppConfig[config.SendgridName], config.AppConfig[config.SendgridMail]))

// 	return instance
// }

// func (service *MailService) getSendgridRequest(mailRequest *mail.SGMailV3) rest.Request {
// 	request := sendgrid.GetRequest(config.AppConfig[config.SendgridKey], "/v3/mail/send", "https://api.sendgrid.com")
// 	request.Method = "POST"
// 	var Body = mail.GetRequestBody(mailRequest)
// 	request.Body = Body

// 	return request
// }

// func (service *MailService) SetTemplateData(p *mail.Personalization, mailType constants.MailType, recipient dto.RecipientDetail, subject string, content string) error {
// 	var errMail error = nil

// 	switch mailType {
// 	case constants.Mail_Verification:
// 		p.SetDynamicTemplateData("subject", subject)
// 		p.SetDynamicTemplateData("name", recipient.Name)
// 		p.SetDynamicTemplateData("email", recipient.Email)
// 		p.SetDynamicTemplateData("redirect_to", recipient.Url)
// 		p.SetDynamicTemplateData("expired_at", recipient.AdditionalData)
// 		break

// 	case constants.Mail_ResetPassword:
// 		p.SetDynamicTemplateData("subject", subject)
// 		p.SetDynamicTemplateData("email", recipient.Email)
// 		p.SetDynamicTemplateData("redirect_to", recipient.Url)
// 		p.SetDynamicTemplateData("expired_at", recipient.AdditionalData)
// 		break

// 	case constants.Mail_VerifyAndSetPassword:
// 		p.SetDynamicTemplateData("subject", subject)
// 		p.SetDynamicTemplateData("nama", recipient.Detail)
// 		p.SetDynamicTemplateData("email", recipient.Email)
// 		p.SetDynamicTemplateData("redirect_to", recipient.Url)
// 		p.SetDynamicTemplateData("expired_at", recipient.AdditionalData)
// 		break

// 	case constants.Mail_WelcomeMail:
// 		p.SetDynamicTemplateData("subject", subject)
// 		unsubsToken := GenerateUnsubscribeToken(recipient.Email)
// 		p.SetDynamicTemplateData("email", recipient.Email)
// 		p.SetDynamicTemplateData("unsubscribe_to", fmt.Sprintf("%s/api/unsubscribe?token=%s", config.AppConfig[config.Hostname], unsubsToken))
// 		break

// 	case constants.Mail_Notification:
// 		unsubsToken := GenerateUnsubscribeToken(recipient.Email)
// 		p.SetDynamicTemplateData("subject", subject)
// 		p.SetDynamicTemplateData("title", subject)
// 		p.SetDynamicTemplateData("content", content)
// 		p.SetDynamicTemplateData("email", recipient.Email)
// 		p.SetDynamicTemplateData("unsubscribe_to", fmt.Sprintf("%s/api/unsubscribe?token=%s", config.AppConfig[config.Hostname], unsubsToken))
// 		break

// 	case constants.Mail_ProfileNotif:
// 		unsubsToken := GenerateUnsubscribeToken(recipient.Email)
// 		p.SetDynamicTemplateData("subject", subject)
// 		p.SetDynamicTemplateData("first_name", recipient.Name)
// 		p.SetDynamicTemplateData("unsubscribe_to", fmt.Sprintf("%s/api/unsubscribe?token=%s", config.AppConfig[config.Hostname], unsubsToken))
// 		break

// 	case constants.Mail_RevPtApprove:
// 		p.SetDynamicTemplateData("subject", subject)
// 		p.SetDynamicTemplateData("nama", recipient.Name)
// 		break

// 	case constants.Mail_RevPtReject:
// 		p.SetDynamicTemplateData("subject", subject)
// 		p.SetDynamicTemplateData("nama", recipient.Name)
// 		break

// 	case constants.Mail_PkApproved:
// 		p.SetDynamicTemplateData("subject", subject)
// 		p.SetDynamicTemplateData("nama", recipient.Name)
// 		break

// 	case constants.Mail_PkRevisi:
// 		p.SetDynamicTemplateData("subject", subject)
// 		p.SetDynamicTemplateData("nama", recipient.Name)
// 		p.SetDynamicTemplateData("alasan_revisi", recipient.Detail)
// 		break

// 	case constants.Mail_RevPtAssignPropA:
// 		var detail dto.MailRevPtAssignPropA
// 		err := json.Unmarshal([]byte(recipient.Detail), &detail)
// 		if err != nil {
// 			errMail = err
// 			break
// 		}
// 		p.SetDynamicTemplateData("subject", subject)
// 		p.SetDynamicTemplateData("nama_reviewer_pt", recipient.Name)
// 		p.SetDynamicTemplateData("nama_perguruan_tinggi", detail.NamaPt)
// 		p.SetDynamicTemplateData("id_proposal_a", detail.IdPropA)
// 		break

// 	case constants.Mail_RevPtAssignPropB:
// 		var detail dto.MailRevPtAssignPropB
// 		err := json.Unmarshal([]byte(recipient.Detail), &detail)
// 		if err != nil {
// 			errMail = err
// 			break
// 		}
// 		p.SetDynamicTemplateData("subject", subject)
// 		p.SetDynamicTemplateData("nama_reviewer_pt", recipient.Name)
// 		p.SetDynamicTemplateData("nama_perguruan_tinggi", detail.NamaPt)
// 		p.SetDynamicTemplateData("nama_matakuliah", detail.NamaMatkul)
// 		p.SetDynamicTemplateData("id_proposal_b", detail.IdPropB)
// 		break

// 	case constants.Mail_Custom:
// 		p.SetDynamicTemplateData("subject", subject)
// 		p.SetDynamicTemplateData("content", content)
// 		break

// 	case constants.Mail_PtSendConnReq:
// 		var detail dto.MailPtSendConnReqData
// 		err := json.Unmarshal([]byte(recipient.Detail), &detail)
// 		if err != nil {
// 			errMail = err
// 			break
// 		}
// 		p.SetDynamicTemplateData("subject", subject)
// 		p.SetDynamicTemplateData("nama", recipient.Name)
// 		p.SetDynamicTemplateData("asal_pt", detail.AsalPt)
// 		// p.SetDynamicTemplateData("skema_kolaborasi", detail.Skema)
// 		p.SetDynamicTemplateData("url_connect_page", detail.Url)
// 		break

// 	case constants.Mail_PtSendColReq:
// 		var detail dto.MailPtSendColReqData
// 		err := json.Unmarshal([]byte(recipient.Detail), &detail)
// 		if err != nil {
// 			errMail = err
// 			break
// 		}
// 		p.SetDynamicTemplateData("subject", subject)
// 		p.SetDynamicTemplateData("nama", recipient.Name)
// 		p.SetDynamicTemplateData("asal_pt", detail.AsalPt)
// 		p.SetDynamicTemplateData("skema_kolaborasi", detail.Skema)
// 		p.SetDynamicTemplateData("url_collab_page", detail.Url)
// 		break

// 	case constants.Mail_PkAccConnReq:
// 		var detail dto.MailPkAppConnReqData
// 		err := json.Unmarshal([]byte(recipient.Detail), &detail)
// 		if err != nil {
// 			errMail = err
// 			break
// 		}
// 		p.SetDynamicTemplateData("subject", subject)
// 		p.SetDynamicTemplateData("nama_ketua_pelaksana", recipient.Name)
// 		p.SetDynamicTemplateData("asal_pt", detail.AsalPt)
// 		p.SetDynamicTemplateData("nama_praktisi", detail.NamaPk)
// 		p.SetDynamicTemplateData("url_profil_praktisi", detail.Url)
// 		break

// 	case constants.Mail_PkRejConnReq:
// 		var detail dto.MailPkRejConnReqData
// 		err := json.Unmarshal([]byte(recipient.Detail), &detail)
// 		if err != nil {
// 			errMail = err
// 			break
// 		}
// 		p.SetDynamicTemplateData("subject", subject)
// 		p.SetDynamicTemplateData("nama_ketua_pelaksana", recipient.Name)
// 		p.SetDynamicTemplateData("asal_pt", detail.AsalPt)
// 		p.SetDynamicTemplateData("nama_praktisi", detail.NamaPk)
// 		break

// 	case constants.Mail_PkAccColReq:
// 		var detail dto.MailPkAppColReqData
// 		err := json.Unmarshal([]byte(recipient.Detail), &detail)
// 		if err != nil {
// 			errMail = err
// 			break
// 		}
// 		p.SetDynamicTemplateData("subject", subject)
// 		p.SetDynamicTemplateData("nama_ketua_pelaksana", recipient.Name)
// 		p.SetDynamicTemplateData("asal_pt", detail.AsalPt)
// 		p.SetDynamicTemplateData("nama_praktisi", detail.NamaPk)
// 		p.SetDynamicTemplateData("url_profil_praktisi", detail.Url)
// 		break

// 	case constants.Mail_PkRejColReq:
// 		var detail dto.MailPkRejColReqData
// 		err := json.Unmarshal([]byte(recipient.Detail), &detail)
// 		if err != nil {
// 			errMail = err
// 			break
// 		}
// 		p.SetDynamicTemplateData("subject", subject)
// 		p.SetDynamicTemplateData("nama_ketua_pelaksana", recipient.Name)
// 		p.SetDynamicTemplateData("asal_pt", detail.AsalPt)
// 		p.SetDynamicTemplateData("nama_praktisi", detail.NamaPk)
// 		break

// 	case constants.Mail_PtPubPropA:
// 		var detail dto.MailPtPubPropAData
// 		err := json.Unmarshal([]byte(recipient.Detail), &detail)
// 		if err != nil {
// 			errMail = err
// 			break
// 		}
// 		p.SetDynamicTemplateData("subject", subject)
// 		p.SetDynamicTemplateData("nama_ketua_pelaksana", recipient.Name)
// 		p.SetDynamicTemplateData("asal_pt", detail.AsalPt)
// 		break

// 	case constants.Mail_PtPubPropB:
// 		var detail dto.MailPtPubPropBData
// 		err := json.Unmarshal([]byte(recipient.Detail), &detail)
// 		if err != nil {
// 			errMail = err
// 			break
// 		}
// 		p.SetDynamicTemplateData("subject", subject)
// 		p.SetDynamicTemplateData("nama_ketua_pelaksana", recipient.Name)
// 		p.SetDynamicTemplateData("asal_pt", detail.AsalPt)
// 		p.SetDynamicTemplateData("nama_matakuliah", detail.Matkul)
// 		break

// 	case constants.Mail_PraktisiManualApprove:
// 		var detail dto.MailPkManualApprove
// 		err := json.Unmarshal([]byte(recipient.Detail), &detail)
// 		if err != nil {
// 			errMail = err
// 			break
// 		}
// 		p.SetDynamicTemplateData("subject", subject)
// 		p.SetDynamicTemplateData("perguruan_tinggi", detail.NamaPt)
// 		p.SetDynamicTemplateData("mata_kuliah", detail.NamaMatkul)
// 		break

// 	case constants.Mail_PraktisiManualReject:
// 		var detail dto.MailPkManualReject
// 		err := json.Unmarshal([]byte(recipient.Detail), &detail)
// 		if err != nil {
// 			errMail = err
// 			break
// 		}
// 		p.SetDynamicTemplateData("subject", subject)
// 		p.SetDynamicTemplateData("perguruan_tinggi", detail.NamaPt)
// 		p.SetDynamicTemplateData("mata_kuliah", detail.NamaMatkul)
// 		break

// 	case constants.Mail_RevisiGeneral:
// 		var detail dto.MailGeneralRevisiData
// 		err := json.Unmarshal([]byte(recipient.Detail), &detail)
// 		if err != nil {
// 			errMail = err
// 			break
// 		}
// 		p.SetDynamicTemplateData("subject", subject)
// 		p.SetDynamicTemplateData("title", detail.Title)
// 		p.SetDynamicTemplateData("catatan", detail.Catatan)
// 		break

// 	}

// 	return errMail
// }

// func (service *MailService) SendTemplateEmail(mailType constants.MailType, recipients []dto.RecipientDetail, templateID string, subject string, content string) {
// 	log.Printf("sending mail : [%s] with template : [%s] to : [\n", strings.ToUpper(constants.MailTypeDescriptor[mailType]), templateID)

// 	for _, recipient := range recipients {
// 		mailRequest := service.initSendgridRequest()

// 		mailRequest.SetTemplateID(templateID)

// 		p := mail.NewPersonalization()
// 		var tos []*mail.Email
// 		tos = append(tos, mail.NewEmail(recipient.Name, recipient.Email))
// 		p.AddTos(tos...)
// 		p.Subject = subject

// 		err := service.SetTemplateData(p, mailType, recipient, subject, content)
// 		if err != nil {
// 			log.Printf("Skip sending to this recipient because content format invalid : %s (%v)", recipient.Email, err)
// 			continue
// 		}

// 		mailRequest.AddPersonalizations(p)

// 		formatedEmail := utilities.CensorFirst([]byte(recipient.Email), 6)

// 		request := service.getSendgridRequest(mailRequest)
// 		response, err := sendgrid.API(request)
// 		if err != nil {
// 			log.Printf("\t [FAILED] %s: %v\n", formatedEmail, err.Error())
// 		} else {
// 			log.Printf("\t [SUCCESS] %s: %d\n", formatedEmail, response.StatusCode)
// 		}
// 	}
// 	log.Printf("]\n")
// }

// func (service *MailService) GenerateRecipients(mailType constants.MailType, detail *dto.MailSchedulerDetailReq, redirectTarget string) (recipient dto.RecipientDetail) {
// 	if detail.IdAccount != uuid.Nil {
// 		recipient.IdAccount = detail.IdAccount
// 		recipient.Name = detail.Name
// 		recipient.Detail = detail.Detail
// 	}
// 	if detail.Role != 0 {
// 		recipient.Role = utilities.AddSpace(constants.UserRoleDescriptor[detail.Role])
// 	}
// 	if redirectTarget != "" {
// 		recipient.Url, recipient.AdditionalData = service.generateURL(mailType, detail, redirectTarget)
// 	}
// 	recipient.Email = detail.Email

// 	return recipient
// }

// func (usecase *MailService) generateURL(mailType constants.MailType, detail *dto.MailSchedulerDetailReq, targetUrl string) (url string, additional string) {
// 	var token string
// 	var expireAt string
// 	switch mailType {
// 	case constants.Mail_Verification:
// 		token, expireAt = GenerateVerifyEmailToken(detail.IdAccount, detail.Email)
// 		break
// 	case constants.Mail_ResetPassword:
// 		token, expireAt = GenerateResetPasswordToken(detail.Email)
// 		break
// 	case constants.Mail_VerifyAndSetPassword:
// 		token, expireAt = GenerateResetPasswordToken(detail.Email)
// 		break
// 	}
// 	url = fmt.Sprint(targetUrl + token)

// 	return url, expireAt
// }
