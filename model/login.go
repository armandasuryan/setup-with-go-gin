package model

type (
	Login struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	VerifyOTP struct {
		Username string `json:"username" validate:"required"`
		OTPCode  int    `json:"otp" validate:"required"`
	}

	OTP struct {
		OTP      int    `json:"otp"`
		TimeLeft string `json:"time_left"`
	}

	ResponseLogin struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		RoleName string `json:"role_name"`
		Name     string `json:"employee_name"`
		Email    string `json:"email"`
		Token    string `json:"token"`
	}

	UserProfile struct {
		ID    int    `json:"id"`
		Name  string `json:"employee_name"`
		Email string `json:"email"`
	}
)
