package keybase_v1

type Status struct {
	Code   int      `codec:"code"`
	Name   string   `codec:"name"`
	Desc   string   `codec:"desc"`
	Fields []string `codec:"fields"`
}
type UID [16]byte
type LoginResBody struct {
	Uid UID `codec:"uid"`
}
type IsLoggedInRes struct {
	Body   *LoginResBody `codec:"body,omitempty"`
	Status Status        `codec:"status"`
}
type IsLoggedInArg struct {
}
type PasswordLoginArg struct {
	Password string `codec:"password"`
}
type PubkeyLoginArg struct {
}
type PasswordLoginRes struct {
	Body   *LoginResBody `codec:"body,omitempty"`
	Status Status        `codec:"status"`
}
type PubkeyLoginRes struct {
	Body   *LoginResBody `codec:"body,omitempty"`
	Status Status        `codec:"status"`
}
type LogoutRes struct {
	Status Status `codec:"status"`
}
type LogoutArg struct {
}
type SwitchUserRes struct {
	Status Status `codec:"status"`
}
type SwitchUserArg struct {
	Username string `codec:"username"`
}
type Login interface {
	IsLoggedIn(arg IsLoggedInArg) IsLoggedInRes
	PasswordLogin(arg PasswordLoginArg) PasswordLoginRes
	PubkeyLogin(arg PubkeyLoginArg) PubkeyLoginRes
	Logout(arg LogoutArg) LogoutRes
	SwitchUser(arg SwitchUserArg) SwitchUserRes
}
type CheckUsernameAvailableRes struct {
	Status Status `codec:"status"`
}
type CheckUsernameAvailableArg struct {
	Username string `codec:"username"`
}
type CheckEmailAvailableRes struct {
	Status Status `codec:"status"`
}
type CheckEmailAvailableArg struct {
	Email string `codec:"email"`
}
type SignupArg struct {
	Email      string `codec:"email"`
	InviteCode string `codec:"inviteCode"`
	Password   string `codec:"password"`
	Username   string `codec:"username"`
}
type InviteRequestArg struct {
	Email    string `codec:"email"`
	FullName string `codec:"fullName"`
	Notes    string `codec:"notes"`
}
type InviteRequestResBody struct {
	Code  string `codec:"code"`
	Place int    `codec:"place"`
}
type InviteRequestRes struct {
	Body   *InviteRequestResBody `codec:"body,omitempty"`
	Status Status                `codec:"status"`
}
type SignupResSuccess struct {
	Uid UID `codec:"uid"`
}
type SignupRes struct {
	Body   *SignupResSuccess `codec:"body,omitempty"`
	Status Status            `codec:"status"`
}
type Signup interface {
	CheckUsernameAvailable(arg CheckUsernameAvailableArg) CheckUsernameAvailableRes
	CheckEmailAvailable(arg CheckEmailAvailableArg) CheckEmailAvailableRes
	Signup(arg SignupArg) SignupRes
	InviteResuest(arg InviteRequestArg) InviteRequestRes
}
