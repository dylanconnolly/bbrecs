package bbrecs

type Group struct {
	ID         uint   `json:"uid"`
	Name       string `json:"name"`
	InviteCode string `json:"inviteCode"`
}

type GroupUser struct {
	ID      uint
	UserID  uint
	GroupID uint
}

type User struct {
	ID          uint   `json:"uid"`
	DisplayName string `json:"displayName"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `jsone:"phoneNumber"`
}

type Rec struct {
	Title string `json:"title"`
	// Tags enum #allow user creation?
	AuthorComment string `json:"authorComment"`
	// AuthorPhotos #grab image from google api
}

type Photo struct {
}
