package gitlab

import "time"

// gitlab user info
// access_level master所对应的access_level为40，developer的权限为30，将access_level从40改成30就实现了降权
type GitlabUser struct {
	Id               uint      `json:"id"`                 //用户id
	Name             string    `json:"name"`               //用户昵称
	Username         string    `json:"username"`           //用户登录名称
	State            string    `json:"state"`              //用户状态,active、blocked
	AvatarUrl        string    `json:"avatar_url"`         //用户头像地址
	WebUrl           string    `json:"web_url"`            //用户主页地址
	CreateAt         time.Time `json:"create_at"`          //用户创建时间
	LastSignInAt     time.Time `json:"last_sign_in_at"`    //用户最近一次登录时间
	Email            string    `json:"email"`              //用户邮箱
	CanCreateGroup   bool      `json:"can_create_group"`   //用户是否具备权限创建群组
	CanCreateProject bool      `json:"can_create_project"` //用户是否具备权限创建项目
	IsAdmin          bool      `json:"is_admin"`           //用户是否为管理员

	AccessLevel uint `json:"access_level"` //用户所处权限
}

//是否已经被锁定
func (t *GitlabUser) IsStateBlocked() bool {
	return t.State == "blocked"
}