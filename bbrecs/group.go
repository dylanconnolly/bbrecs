package bbrecs

type Group struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	InviteCode string `json:"inviteCode"`
}

type GroupUser struct {
	ID      uint
	UserID  uint
	GroupID uint
}
