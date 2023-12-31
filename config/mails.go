package config

import "time"

const (
	VERIFICATION_EMAIL_SUBJECT          = "OTP For Verification | Interact"
	VERIFICATION_DELETE_SUBJECT         = "OTP For Deletion | Interact"
	VERIFICATION_DELETE_PROJECT_SUBJECT = "OTP For Project Deletion | Interact"
	VERIFICATION_LEAVE_ORG_SUBJECT      = "OTP For Leaving an Org | Interact"
	VERIFICATION_EMAIL_BODY             = "OTP: "
	VERIFICATION_OTP_EXPIRATION_TIME    = 10 * time.Minute

	EARLY_ACCESS_EMAIL_SUBJECT         = "Your EARLY ACCESS Token! | Interact"
	EARLY_ACCESS_EMAIL_BODY            = "Your token for early access is: "
	EARLY_ACCESS_TOKEN_EXPIRATION_TIME = EARLY_ACCESS_TOKEN_TTL
)
