package domain

type Role string

type User struct {
	ID       string
	Name     string
	Password string
	Role     UserRole //admin, SuperModerator, ContentModerator, ForumModerator

	EncryptedPassword string
	Email             string
	OnlineChan        chan int
	IsOnline          bool

	UserListID int
	UserLists  *UserLists

	Articles []*Article
	Comments []*Comment

	Notifications []*Notification //new articles, news, replyes etc

	Chats []*Chat
}
type UserRole struct {
	UsersIDs string
	Role     Role
}
