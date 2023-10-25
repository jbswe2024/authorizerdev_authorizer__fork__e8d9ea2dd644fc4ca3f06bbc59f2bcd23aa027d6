// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AddEmailTemplateRequest struct {
	EventName string  `json:"event_name"`
	Subject   string  `json:"subject"`
	Template  string  `json:"template"`
	Design    *string `json:"design,omitempty"`
}

type AddWebhookRequest struct {
	EventName        string                 `json:"event_name"`
	EventDescription *string                `json:"event_description,omitempty"`
	Endpoint         string                 `json:"endpoint"`
	Enabled          bool                   `json:"enabled"`
	Headers          map[string]interface{} `json:"headers,omitempty"`
}

type AdminLoginInput struct {
	AdminSecret string `json:"admin_secret"`
}

type AdminSignupInput struct {
	AdminSecret string `json:"admin_secret"`
}

type AuthResponse struct {
	Message                   string  `json:"message"`
	ShouldShowEmailOtpScreen  *bool   `json:"should_show_email_otp_screen,omitempty"`
	ShouldShowMobileOtpScreen *bool   `json:"should_show_mobile_otp_screen,omitempty"`
	AccessToken               *string `json:"access_token,omitempty"`
	IDToken                   *string `json:"id_token,omitempty"`
	RefreshToken              *string `json:"refresh_token,omitempty"`
	ExpiresIn                 *int64  `json:"expires_in,omitempty"`
	User                      *User   `json:"user,omitempty"`
}

type DeleteEmailTemplateRequest struct {
	ID string `json:"id"`
}

type DeleteUserInput struct {
	Email string `json:"email"`
}

type EmailTemplate struct {
	ID        string `json:"id"`
	EventName string `json:"event_name"`
	Template  string `json:"template"`
	Design    string `json:"design"`
	Subject   string `json:"subject"`
	CreatedAt *int64 `json:"created_at,omitempty"`
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

type EmailTemplates struct {
	Pagination     *Pagination      `json:"pagination"`
	EmailTemplates []*EmailTemplate `json:"email_templates"`
}

type Env struct {
	AccessTokenExpiryTime            *string  `json:"ACCESS_TOKEN_EXPIRY_TIME,omitempty"`
	AdminSecret                      *string  `json:"ADMIN_SECRET,omitempty"`
	DatabaseName                     *string  `json:"DATABASE_NAME,omitempty"`
	DatabaseURL                      *string  `json:"DATABASE_URL,omitempty"`
	DatabaseType                     *string  `json:"DATABASE_TYPE,omitempty"`
	DatabaseUsername                 *string  `json:"DATABASE_USERNAME,omitempty"`
	DatabasePassword                 *string  `json:"DATABASE_PASSWORD,omitempty"`
	DatabaseHost                     *string  `json:"DATABASE_HOST,omitempty"`
	DatabasePort                     *string  `json:"DATABASE_PORT,omitempty"`
	ClientID                         string   `json:"CLIENT_ID"`
	ClientSecret                     string   `json:"CLIENT_SECRET"`
	CustomAccessTokenScript          *string  `json:"CUSTOM_ACCESS_TOKEN_SCRIPT,omitempty"`
	SMTPHost                         *string  `json:"SMTP_HOST,omitempty"`
	SMTPPort                         *string  `json:"SMTP_PORT,omitempty"`
	SMTPUsername                     *string  `json:"SMTP_USERNAME,omitempty"`
	SMTPPassword                     *string  `json:"SMTP_PASSWORD,omitempty"`
	SMTPLocalName                    *string  `json:"SMTP_LOCAL_NAME,omitempty"`
	SenderEmail                      *string  `json:"SENDER_EMAIL,omitempty"`
	SenderName                       *string  `json:"SENDER_NAME,omitempty"`
	JwtType                          *string  `json:"JWT_TYPE,omitempty"`
	JwtSecret                        *string  `json:"JWT_SECRET,omitempty"`
	JwtPrivateKey                    *string  `json:"JWT_PRIVATE_KEY,omitempty"`
	JwtPublicKey                     *string  `json:"JWT_PUBLIC_KEY,omitempty"`
	AllowedOrigins                   []string `json:"ALLOWED_ORIGINS,omitempty"`
	AppURL                           *string  `json:"APP_URL,omitempty"`
	RedisURL                         *string  `json:"REDIS_URL,omitempty"`
	ResetPasswordURL                 *string  `json:"RESET_PASSWORD_URL,omitempty"`
	DisableEmailVerification         bool     `json:"DISABLE_EMAIL_VERIFICATION"`
	DisableBasicAuthentication       bool     `json:"DISABLE_BASIC_AUTHENTICATION"`
	DisableMagicLinkLogin            bool     `json:"DISABLE_MAGIC_LINK_LOGIN"`
	DisableLoginPage                 bool     `json:"DISABLE_LOGIN_PAGE"`
	DisableSignUp                    bool     `json:"DISABLE_SIGN_UP"`
	DisableRedisForEnv               bool     `json:"DISABLE_REDIS_FOR_ENV"`
	DisableStrongPassword            bool     `json:"DISABLE_STRONG_PASSWORD"`
	DisableMultiFactorAuthentication bool     `json:"DISABLE_MULTI_FACTOR_AUTHENTICATION"`
	EnforceMultiFactorAuthentication bool     `json:"ENFORCE_MULTI_FACTOR_AUTHENTICATION"`
	Roles                            []string `json:"ROLES,omitempty"`
	ProtectedRoles                   []string `json:"PROTECTED_ROLES,omitempty"`
	DefaultRoles                     []string `json:"DEFAULT_ROLES,omitempty"`
	JwtRoleClaim                     *string  `json:"JWT_ROLE_CLAIM,omitempty"`
	GoogleClientID                   *string  `json:"GOOGLE_CLIENT_ID,omitempty"`
	GoogleClientSecret               *string  `json:"GOOGLE_CLIENT_SECRET,omitempty"`
	GithubClientID                   *string  `json:"GITHUB_CLIENT_ID,omitempty"`
	GithubClientSecret               *string  `json:"GITHUB_CLIENT_SECRET,omitempty"`
	FacebookClientID                 *string  `json:"FACEBOOK_CLIENT_ID,omitempty"`
	FacebookClientSecret             *string  `json:"FACEBOOK_CLIENT_SECRET,omitempty"`
	LinkedinClientID                 *string  `json:"LINKEDIN_CLIENT_ID,omitempty"`
	LinkedinClientSecret             *string  `json:"LINKEDIN_CLIENT_SECRET,omitempty"`
	AppleClientID                    *string  `json:"APPLE_CLIENT_ID,omitempty"`
	AppleClientSecret                *string  `json:"APPLE_CLIENT_SECRET,omitempty"`
	TwitterClientID                  *string  `json:"TWITTER_CLIENT_ID,omitempty"`
	TwitterClientSecret              *string  `json:"TWITTER_CLIENT_SECRET,omitempty"`
	MicrosoftClientID                *string  `json:"MICROSOFT_CLIENT_ID,omitempty"`
	MicrosoftClientSecret            *string  `json:"MICROSOFT_CLIENT_SECRET,omitempty"`
	MicrosoftActiveDirectoryTenantID *string  `json:"MICROSOFT_ACTIVE_DIRECTORY_TENANT_ID,omitempty"`
	OrganizationName                 *string  `json:"ORGANIZATION_NAME,omitempty"`
	OrganizationLogo                 *string  `json:"ORGANIZATION_LOGO,omitempty"`
	AppCookieSecure                  bool     `json:"APP_COOKIE_SECURE"`
	AdminCookieSecure                bool     `json:"ADMIN_COOKIE_SECURE"`
	DefaultAuthorizeResponseType     *string  `json:"DEFAULT_AUTHORIZE_RESPONSE_TYPE,omitempty"`
	DefaultAuthorizeResponseMode     *string  `json:"DEFAULT_AUTHORIZE_RESPONSE_MODE,omitempty"`
	DisablePlayground                bool     `json:"DISABLE_PLAYGROUND"`
}

type Error struct {
	Message string `json:"message"`
	Reason  string `json:"reason"`
}

type ForgotPasswordInput struct {
	Email       string  `json:"email"`
	State       *string `json:"state,omitempty"`
	RedirectURI *string `json:"redirect_uri,omitempty"`
}

type GenerateJWTKeysInput struct {
	Type string `json:"type"`
}

type GenerateJWTKeysResponse struct {
	Secret     *string `json:"secret,omitempty"`
	PublicKey  *string `json:"public_key,omitempty"`
	PrivateKey *string `json:"private_key,omitempty"`
}

type GetUserRequest struct {
	ID    *string `json:"id,omitempty"`
	Email *string `json:"email,omitempty"`
}

type InviteMemberInput struct {
	Emails      []string `json:"emails"`
	RedirectURI *string  `json:"redirect_uri,omitempty"`
}

type InviteMembersResponse struct {
	Message string  `json:"message"`
	Users   []*User `json:"Users"`
}

type ListWebhookLogRequest struct {
	Pagination *PaginationInput `json:"pagination,omitempty"`
	WebhookID  *string          `json:"webhook_id,omitempty"`
}

type LoginInput struct {
	Email       *string  `json:"email,omitempty"`
	PhoneNumber *string  `json:"phone_number,omitempty"`
	Password    string   `json:"password"`
	Roles       []string `json:"roles,omitempty"`
	Scope       []string `json:"scope,omitempty"`
	State       *string  `json:"state,omitempty"`
}

type MagicLinkLoginInput struct {
	Email       string   `json:"email"`
	Roles       []string `json:"roles,omitempty"`
	Scope       []string `json:"scope,omitempty"`
	State       *string  `json:"state,omitempty"`
	RedirectURI *string  `json:"redirect_uri,omitempty"`
}

type Meta struct {
	Version                      string `json:"version"`
	ClientID                     string `json:"client_id"`
	IsGoogleLoginEnabled         bool   `json:"is_google_login_enabled"`
	IsFacebookLoginEnabled       bool   `json:"is_facebook_login_enabled"`
	IsGithubLoginEnabled         bool   `json:"is_github_login_enabled"`
	IsLinkedinLoginEnabled       bool   `json:"is_linkedin_login_enabled"`
	IsAppleLoginEnabled          bool   `json:"is_apple_login_enabled"`
	IsTwitterLoginEnabled        bool   `json:"is_twitter_login_enabled"`
	IsMicrosoftLoginEnabled      bool   `json:"is_microsoft_login_enabled"`
	IsEmailVerificationEnabled   bool   `json:"is_email_verification_enabled"`
	IsBasicAuthenticationEnabled bool   `json:"is_basic_authentication_enabled"`
	IsMagicLinkLoginEnabled      bool   `json:"is_magic_link_login_enabled"`
	IsSignUpEnabled              bool   `json:"is_sign_up_enabled"`
	IsStrongPasswordEnabled      bool   `json:"is_strong_password_enabled"`
	IsMultiFactorAuthEnabled     bool   `json:"is_multi_factor_auth_enabled"`
}

type MobileLoginInput struct {
	PhoneNumber string   `json:"phone_number"`
	Password    string   `json:"password"`
	Roles       []string `json:"roles,omitempty"`
	Scope       []string `json:"scope,omitempty"`
	State       *string  `json:"state,omitempty"`
}

type MobileSignUpInput struct {
	Email                    *string                `json:"email,omitempty"`
	GivenName                *string                `json:"given_name,omitempty"`
	FamilyName               *string                `json:"family_name,omitempty"`
	MiddleName               *string                `json:"middle_name,omitempty"`
	Nickname                 *string                `json:"nickname,omitempty"`
	Gender                   *string                `json:"gender,omitempty"`
	Birthdate                *string                `json:"birthdate,omitempty"`
	PhoneNumber              string                 `json:"phone_number"`
	Picture                  *string                `json:"picture,omitempty"`
	Password                 string                 `json:"password"`
	ConfirmPassword          string                 `json:"confirm_password"`
	Roles                    []string               `json:"roles,omitempty"`
	Scope                    []string               `json:"scope,omitempty"`
	RedirectURI              *string                `json:"redirect_uri,omitempty"`
	IsMultiFactorAuthEnabled *bool                  `json:"is_multi_factor_auth_enabled,omitempty"`
	State                    *string                `json:"state,omitempty"`
	AppData                  map[string]interface{} `json:"app_data,omitempty"`
}

type OAuthRevokeInput struct {
	RefreshToken string `json:"refresh_token"`
}

type PaginatedInput struct {
	Pagination *PaginationInput `json:"pagination,omitempty"`
}

type Pagination struct {
	Limit  int64 `json:"limit"`
	Page   int64 `json:"page"`
	Offset int64 `json:"offset"`
	Total  int64 `json:"total"`
}

type PaginationInput struct {
	Limit *int64 `json:"limit,omitempty"`
	Page  *int64 `json:"page,omitempty"`
}

type ResendOTPRequest struct {
	Email       *string `json:"email,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	State       *string `json:"state,omitempty"`
}

type ResendVerifyEmailInput struct {
	Email      string  `json:"email"`
	Identifier string  `json:"identifier"`
	State      *string `json:"state,omitempty"`
}

type ResetPasswordInput struct {
	Token           string `json:"token"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type Response struct {
	Message string `json:"message"`
}

type SMSVerificationRequests struct {
	ID            string `json:"id"`
	Code          string `json:"code"`
	CodeExpiresAt int64  `json:"code_expires_at"`
	PhoneNumber   string `json:"phone_number"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     *int64 `json:"updated_at,omitempty"`
}

type SessionQueryInput struct {
	Roles []string `json:"roles,omitempty"`
	Scope []string `json:"scope,omitempty"`
}

type SignUpInput struct {
	Email                    *string                `json:"email,omitempty"`
	GivenName                *string                `json:"given_name,omitempty"`
	FamilyName               *string                `json:"family_name,omitempty"`
	MiddleName               *string                `json:"middle_name,omitempty"`
	Nickname                 *string                `json:"nickname,omitempty"`
	Gender                   *string                `json:"gender,omitempty"`
	Birthdate                *string                `json:"birthdate,omitempty"`
	PhoneNumber              *string                `json:"phone_number,omitempty"`
	Picture                  *string                `json:"picture,omitempty"`
	Password                 string                 `json:"password"`
	ConfirmPassword          string                 `json:"confirm_password"`
	Roles                    []string               `json:"roles,omitempty"`
	Scope                    []string               `json:"scope,omitempty"`
	RedirectURI              *string                `json:"redirect_uri,omitempty"`
	IsMultiFactorAuthEnabled *bool                  `json:"is_multi_factor_auth_enabled,omitempty"`
	State                    *string                `json:"state,omitempty"`
	AppData                  map[string]interface{} `json:"app_data,omitempty"`
}

type TestEndpointRequest struct {
	Endpoint         string                 `json:"endpoint"`
	EventName        string                 `json:"event_name"`
	EventDescription *string                `json:"event_description,omitempty"`
	Headers          map[string]interface{} `json:"headers,omitempty"`
}

type TestEndpointResponse struct {
	HTTPStatus *int64  `json:"http_status,omitempty"`
	Response   *string `json:"response,omitempty"`
}

type UpdateAccessInput struct {
	UserID string `json:"user_id"`
}

type UpdateEmailTemplateRequest struct {
	ID        string  `json:"id"`
	EventName *string `json:"event_name,omitempty"`
	Template  *string `json:"template,omitempty"`
	Subject   *string `json:"subject,omitempty"`
	Design    *string `json:"design,omitempty"`
}

type UpdateEnvInput struct {
	AccessTokenExpiryTime            *string  `json:"ACCESS_TOKEN_EXPIRY_TIME,omitempty"`
	AdminSecret                      *string  `json:"ADMIN_SECRET,omitempty"`
	CustomAccessTokenScript          *string  `json:"CUSTOM_ACCESS_TOKEN_SCRIPT,omitempty"`
	OldAdminSecret                   *string  `json:"OLD_ADMIN_SECRET,omitempty"`
	SMTPHost                         *string  `json:"SMTP_HOST,omitempty"`
	SMTPPort                         *string  `json:"SMTP_PORT,omitempty"`
	SMTPUsername                     *string  `json:"SMTP_USERNAME,omitempty"`
	SMTPPassword                     *string  `json:"SMTP_PASSWORD,omitempty"`
	SMTPLocalName                    *string  `json:"SMTP_LOCAL_NAME,omitempty"`
	SenderEmail                      *string  `json:"SENDER_EMAIL,omitempty"`
	SenderName                       *string  `json:"SENDER_NAME,omitempty"`
	JwtType                          *string  `json:"JWT_TYPE,omitempty"`
	JwtSecret                        *string  `json:"JWT_SECRET,omitempty"`
	JwtPrivateKey                    *string  `json:"JWT_PRIVATE_KEY,omitempty"`
	JwtPublicKey                     *string  `json:"JWT_PUBLIC_KEY,omitempty"`
	AllowedOrigins                   []string `json:"ALLOWED_ORIGINS,omitempty"`
	AppURL                           *string  `json:"APP_URL,omitempty"`
	ResetPasswordURL                 *string  `json:"RESET_PASSWORD_URL,omitempty"`
	AppCookieSecure                  *bool    `json:"APP_COOKIE_SECURE,omitempty"`
	AdminCookieSecure                *bool    `json:"ADMIN_COOKIE_SECURE,omitempty"`
	DisableEmailVerification         *bool    `json:"DISABLE_EMAIL_VERIFICATION,omitempty"`
	DisableBasicAuthentication       *bool    `json:"DISABLE_BASIC_AUTHENTICATION,omitempty"`
	DisableMagicLinkLogin            *bool    `json:"DISABLE_MAGIC_LINK_LOGIN,omitempty"`
	DisableLoginPage                 *bool    `json:"DISABLE_LOGIN_PAGE,omitempty"`
	DisableSignUp                    *bool    `json:"DISABLE_SIGN_UP,omitempty"`
	DisableRedisForEnv               *bool    `json:"DISABLE_REDIS_FOR_ENV,omitempty"`
	DisableStrongPassword            *bool    `json:"DISABLE_STRONG_PASSWORD,omitempty"`
	DisableMultiFactorAuthentication *bool    `json:"DISABLE_MULTI_FACTOR_AUTHENTICATION,omitempty"`
	EnforceMultiFactorAuthentication *bool    `json:"ENFORCE_MULTI_FACTOR_AUTHENTICATION,omitempty"`
	Roles                            []string `json:"ROLES,omitempty"`
	ProtectedRoles                   []string `json:"PROTECTED_ROLES,omitempty"`
	DefaultRoles                     []string `json:"DEFAULT_ROLES,omitempty"`
	JwtRoleClaim                     *string  `json:"JWT_ROLE_CLAIM,omitempty"`
	GoogleClientID                   *string  `json:"GOOGLE_CLIENT_ID,omitempty"`
	GoogleClientSecret               *string  `json:"GOOGLE_CLIENT_SECRET,omitempty"`
	GithubClientID                   *string  `json:"GITHUB_CLIENT_ID,omitempty"`
	GithubClientSecret               *string  `json:"GITHUB_CLIENT_SECRET,omitempty"`
	FacebookClientID                 *string  `json:"FACEBOOK_CLIENT_ID,omitempty"`
	FacebookClientSecret             *string  `json:"FACEBOOK_CLIENT_SECRET,omitempty"`
	LinkedinClientID                 *string  `json:"LINKEDIN_CLIENT_ID,omitempty"`
	LinkedinClientSecret             *string  `json:"LINKEDIN_CLIENT_SECRET,omitempty"`
	AppleClientID                    *string  `json:"APPLE_CLIENT_ID,omitempty"`
	AppleClientSecret                *string  `json:"APPLE_CLIENT_SECRET,omitempty"`
	TwitterClientID                  *string  `json:"TWITTER_CLIENT_ID,omitempty"`
	TwitterClientSecret              *string  `json:"TWITTER_CLIENT_SECRET,omitempty"`
	MicrosoftClientID                *string  `json:"MICROSOFT_CLIENT_ID,omitempty"`
	MicrosoftClientSecret            *string  `json:"MICROSOFT_CLIENT_SECRET,omitempty"`
	MicrosoftActiveDirectoryTenantID *string  `json:"MICROSOFT_ACTIVE_DIRECTORY_TENANT_ID,omitempty"`
	OrganizationName                 *string  `json:"ORGANIZATION_NAME,omitempty"`
	OrganizationLogo                 *string  `json:"ORGANIZATION_LOGO,omitempty"`
	DefaultAuthorizeResponseType     *string  `json:"DEFAULT_AUTHORIZE_RESPONSE_TYPE,omitempty"`
	DefaultAuthorizeResponseMode     *string  `json:"DEFAULT_AUTHORIZE_RESPONSE_MODE,omitempty"`
	DisablePlayground                *bool    `json:"DISABLE_PLAYGROUND,omitempty"`
}

type UpdateProfileInput struct {
	OldPassword              *string                `json:"old_password,omitempty"`
	NewPassword              *string                `json:"new_password,omitempty"`
	ConfirmNewPassword       *string                `json:"confirm_new_password,omitempty"`
	Email                    *string                `json:"email,omitempty"`
	GivenName                *string                `json:"given_name,omitempty"`
	FamilyName               *string                `json:"family_name,omitempty"`
	MiddleName               *string                `json:"middle_name,omitempty"`
	Nickname                 *string                `json:"nickname,omitempty"`
	Gender                   *string                `json:"gender,omitempty"`
	Birthdate                *string                `json:"birthdate,omitempty"`
	PhoneNumber              *string                `json:"phone_number,omitempty"`
	Picture                  *string                `json:"picture,omitempty"`
	IsMultiFactorAuthEnabled *bool                  `json:"is_multi_factor_auth_enabled,omitempty"`
	AppData                  map[string]interface{} `json:"app_data,omitempty"`
}

type UpdateUserInput struct {
	ID                       string                 `json:"id"`
	Email                    *string                `json:"email,omitempty"`
	EmailVerified            *bool                  `json:"email_verified,omitempty"`
	GivenName                *string                `json:"given_name,omitempty"`
	FamilyName               *string                `json:"family_name,omitempty"`
	MiddleName               *string                `json:"middle_name,omitempty"`
	Nickname                 *string                `json:"nickname,omitempty"`
	Gender                   *string                `json:"gender,omitempty"`
	Birthdate                *string                `json:"birthdate,omitempty"`
	PhoneNumber              *string                `json:"phone_number,omitempty"`
	Picture                  *string                `json:"picture,omitempty"`
	Roles                    []*string              `json:"roles,omitempty"`
	IsMultiFactorAuthEnabled *bool                  `json:"is_multi_factor_auth_enabled,omitempty"`
	AppData                  map[string]interface{} `json:"app_data,omitempty"`
}

type UpdateWebhookRequest struct {
	ID               string                 `json:"id"`
	EventName        *string                `json:"event_name,omitempty"`
	EventDescription *string                `json:"event_description,omitempty"`
	Endpoint         *string                `json:"endpoint,omitempty"`
	Enabled          *bool                  `json:"enabled,omitempty"`
	Headers          map[string]interface{} `json:"headers,omitempty"`
}

type User struct {
	ID                       string                 `json:"id"`
	Email                    *string                `json:"email,omitempty"`
	EmailVerified            bool                   `json:"email_verified"`
	SignupMethods            string                 `json:"signup_methods"`
	GivenName                *string                `json:"given_name,omitempty"`
	FamilyName               *string                `json:"family_name,omitempty"`
	MiddleName               *string                `json:"middle_name,omitempty"`
	Nickname                 *string                `json:"nickname,omitempty"`
	PreferredUsername        *string                `json:"preferred_username,omitempty"`
	Gender                   *string                `json:"gender,omitempty"`
	Birthdate                *string                `json:"birthdate,omitempty"`
	PhoneNumber              *string                `json:"phone_number,omitempty"`
	PhoneNumberVerified      *bool                  `json:"phone_number_verified,omitempty"`
	Picture                  *string                `json:"picture,omitempty"`
	Roles                    []string               `json:"roles"`
	CreatedAt                *int64                 `json:"created_at,omitempty"`
	UpdatedAt                *int64                 `json:"updated_at,omitempty"`
	RevokedTimestamp         *int64                 `json:"revoked_timestamp,omitempty"`
	IsMultiFactorAuthEnabled *bool                  `json:"is_multi_factor_auth_enabled,omitempty"`
	AppData                  map[string]interface{} `json:"app_data,omitempty"`
}

type Users struct {
	Pagination *Pagination `json:"pagination"`
	Users      []*User     `json:"users"`
}

type ValidateJWTTokenInput struct {
	TokenType string   `json:"token_type"`
	Token     string   `json:"token"`
	Roles     []string `json:"roles,omitempty"`
}

type ValidateJWTTokenResponse struct {
	IsValid bool                   `json:"is_valid"`
	Claims  map[string]interface{} `json:"claims,omitempty"`
}

type ValidateSessionInput struct {
	Cookie string   `json:"cookie"`
	Roles  []string `json:"roles,omitempty"`
}

type ValidateSessionResponse struct {
	IsValid bool  `json:"is_valid"`
	User    *User `json:"user"`
}

type VerificationRequest struct {
	ID          string  `json:"id"`
	Identifier  *string `json:"identifier,omitempty"`
	Token       *string `json:"token,omitempty"`
	Email       *string `json:"email,omitempty"`
	Expires     *int64  `json:"expires,omitempty"`
	CreatedAt   *int64  `json:"created_at,omitempty"`
	UpdatedAt   *int64  `json:"updated_at,omitempty"`
	Nonce       *string `json:"nonce,omitempty"`
	RedirectURI *string `json:"redirect_uri,omitempty"`
}

type VerificationRequests struct {
	Pagination           *Pagination            `json:"pagination"`
	VerificationRequests []*VerificationRequest `json:"verification_requests"`
}

type VerifyEmailInput struct {
	Token string  `json:"token"`
	State *string `json:"state,omitempty"`
}

type VerifyOTPRequest struct {
	Email       *string `json:"email,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	Otp         string  `json:"otp"`
	State       *string `json:"state,omitempty"`
}

type Webhook struct {
	ID               string                 `json:"id"`
	EventName        *string                `json:"event_name,omitempty"`
	EventDescription *string                `json:"event_description,omitempty"`
	Endpoint         *string                `json:"endpoint,omitempty"`
	Enabled          *bool                  `json:"enabled,omitempty"`
	Headers          map[string]interface{} `json:"headers,omitempty"`
	CreatedAt        *int64                 `json:"created_at,omitempty"`
	UpdatedAt        *int64                 `json:"updated_at,omitempty"`
}

type WebhookLog struct {
	ID         string  `json:"id"`
	HTTPStatus *int64  `json:"http_status,omitempty"`
	Response   *string `json:"response,omitempty"`
	Request    *string `json:"request,omitempty"`
	WebhookID  *string `json:"webhook_id,omitempty"`
	CreatedAt  *int64  `json:"created_at,omitempty"`
	UpdatedAt  *int64  `json:"updated_at,omitempty"`
}

type WebhookLogs struct {
	Pagination  *Pagination   `json:"pagination"`
	WebhookLogs []*WebhookLog `json:"webhook_logs"`
}

type WebhookRequest struct {
	ID string `json:"id"`
}

type Webhooks struct {
	Pagination *Pagination `json:"pagination"`
	Webhooks   []*Webhook  `json:"webhooks"`
}
