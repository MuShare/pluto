package request

type MailRegister struct {
	Mail     string `json:"mail"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (mr MailRegister) Validation() bool {
	if mr.Mail == "" || mr.Password == "" || mr.Name == "" {
		return false
	}
	return true
}

type MailLogin struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
	DeviceID string `json:"device_id"`
	AppID    string `json:"app_id"`
}

func (ml MailLogin) Validation() bool {
	if ml.Mail == "" || ml.Password == "" {
		return false
	}

	if ml.DeviceID == "" || ml.AppID == "" {
		return false
	}

	return true
}
