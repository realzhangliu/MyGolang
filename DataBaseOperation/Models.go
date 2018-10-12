package DataBaseOperation

import (
	"dx/taishan/core/db"
	"dx/taishan/modules/user/models"
	"time"
)

type FileStruct struct {
	db.Model
	ContainerId string `json:"container_id" db:"container_id"`
	VersionId   string `json:"version_id" db:"version_id"`
	RawFileId   string `json:"raw_file_id" db:"raw_file_id"`
	Type        int    `json:"type" db:"type"`
	ParentPath  string `json:"parent_path" db:"parent_path"`
	RecycleId   string `json:"recycle_id" db:"recycle_id"`
	Status      int    `json:"status" db:"status"`
}
type Raw_file struct {
	db.Model
	ContainerId string `json:"container_id" db:"container_id"`
	FileId      string `json:"file_id" db:"file_id"`
	UploaderId  string `json:"uploader_id" db:"uploader_id"`
	VersionId   string `json:"version_id" db:"version_id"`
	Type        int    `json:"type" db:"type"`
	Size        int64  `json:"size" db:"size"`
	Name        string `json:"name" db:"name"`
	Thumbnail   string `json:"thumbnail" db:"thumbnail"`
	RecycleId   string `json:"recycle_id" db:"recycle_id"`
	Status      int    `json:"status" db:"status"`
}

var rf []Raw_file

type Model struct {
	Id        string     `db:"size=32" gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `db:"created_at,null" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at,null" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at,null" json:"deleted_at"`
}

type ContainerMember struct {
	Model
	ContainerId string      `json:"container_id" form:"container_id" gorm:"index"`
	OwnerId     string      `json:"owner_id" form:"owner_id" gorm:"index"`
	Status      int         `json:"status" form:"status"`
	MemberType  int         `json:"member_type" db:"member_type"`
	RecycleId   string      `json:"recycle_id" db:"recycle_id"`
	Owner       models.User `json:"owner"`
}

var CMS []ContainerMember

type UserProfile struct {
	db.Model
	UserId      string     `form:"user_id" binding:"exists" json:"user_id" db:"user_id" gorm:"user_id;size=32;require;unique;index"`
	Avatar      string     `form:"-" json:"avatar" db:"avatar"`
	Company     string     `form:"company" json:"company" db:"company"`
	Title       string     `form:"title" json:"title" db:"title"`
	Address1    string     `form:"address1" json:"address1" db:"address1"`
	Address2    string     `form:"address2" json:"address2" db:"address2"`
	Sex         string     `form:"sex" json:"sex" db:"sex"`
	IpAddress   string     `form:"ip_address" json:"ip_address"`
	Signature   string     `form:"signature" json:"signature" db:"signature"`
	Birthday    *time.Time `form:"birthday" json:"birthday" db:"birthday"`
	School      string     `form:"school" json:"school" db:"school"`
	HomeAddress string     `form:"home_address" json:"home_address" db:"home_address"`
	DetailLevel string     `form:"detail_level" json:"detail_level" db:"detail_level"`
}
type User struct {
	db.Model
	Name     string      `form:"name" binding:"exists,alphanum,min=4,max=255" json:"name"`
	Password string      `form:"password" binding:"exists,min=8,max=255" json:"-"`
	Email    string      `form:"email" binding:"email" json:"email"`
	Phone    string      `form:"phone" json:"phone"`
	Status   int         `form:"-" json:"status"`
	Profile  UserProfile `form:"-" gorm:"foreignkey:UserId" json:"profile"`
	//Company     Company      `form:"-" gorm:"foreignkey:UserId" json:"company"`
	//Roles       []UserRole   `json:"-" gorm:"many2many:user_user_roles"`
	//Oauth       models.Oauth `form:"-" gorm:"foreignkey:UserId" json:"oauth"`
	ActivatedAt time.Time
	CompanyId   string `form:"company_id" json:"company_id" db:"company_id"`
	AreaCode    string `form:"area_code" json:"area_code"`
	//Label       []UserLabel `form:"-" json:"label" gorm:"foreignkey:UserId"`

}
type UserForProfiles struct {
	//db.Model
	Name string `form:"name" binding:"exists,alphanum,min=4,max=255" json:"name"`
	//Password    string       `form:"password" binding:"exists,min=8,max=255" json:"-"`
	//Email       string       `form:"email" binding:"email" json:"email"`
	Phone  string `form:"phone" json:"phone"`
	Status int    `form:"-" json:"status"`
	//Profile     UserProfile  `form:"-" gorm:"foreignkey:UserId" json:"profile"`
	//Company     Company      `form:"-" gorm:"foreignkey:UserId" json:"company"`
	//Roles       []UserRole   `json:"-" gorm:"many2many:user_user_roles"`
	//Oauth       models.Oauth `form:"-" gorm:"foreignkey:UserId" json:"oauth"`
	//ActivatedAt time.Time
	//CompanyId   string      `form:"company_id" json:"company_id" db:"company_id"`
	//AreaCode    string      `form:"area_code" json:"area_code"`
	//Label       []UserLabel `form:"-" json:"label" gorm:"foreignkey:UserId"`
	//UserId      string     `form:"user_id" binding:"exists" json:"user_id" db:"user_id" gorm:"user_id;size=32;require;unique;index"`
	Avatar string `form:"-" json:"avatar" db:"avatar"`
	//Company     string     `form:"company" json:"company" db:"company"`
	//Title       string     `form:"title" json:"title" db:"title"`
	//Address1    string     `form:"address1" json:"address1" db:"address1"`
	//Address2    string     `form:"address2" json:"address2" db:"address2"`
	//Sex         string     `form:"sex" json:"sex" db:"sex"`
	//IpAddress   string     `form:"ip_address" json:"ip_address"`
	//Signature   string     `form:"signature" json:"signature" db:"signature"`
	//Birthday    *time.Time `form:"birthday" json:"birthday" db:"birthday"`
	//School      string     `form:"school" json:"school" db:"school"`
	//HomeAddress string     `form:"home_address" json:"home_address" db:"home_address"`
	//DetailLevel string     `form:"detail_level" json:"detail_level" db:"detail_level"`
}
type UserWebApp struct {
	db.Model
	UserId        string    `json:"user_id" form:"user_id" db:"user_id"`
	AppId         string    `json:"app_id" form:"app_id" db:"app_id"`
	DueTime       time.Time `json:"due_time" form:"due_time" db:"due_time"`
	SubscribeTime int       `json:"subscribe_time" form:"subscribe_time" db:"subscribe_time"`
	Status        int       `json:"status" form:"status" db:"status"`
}

var userWebApp UserWebApp
var userWebApps []UserWebApp

type WebApp struct {
	db.Model
	Name        string    `json:"name" form:"name" db:"name"`
	Description string    `json:"description" form:"description" db:"description"`
	Size        int64     `json:"size" form:"size" db:"size"`
	PublisherId string    `json:"publisher_id" form:"publisher_id" db:"publisher_id"`
	PublishDate time.Time `json:"pblish_date" form:'pblish_date' db:"pblish_date"`
	Version     string    `json:"version" form:"version" db:'version'`
	FreeDays    int       `json:"free_days" from:“free_days” db:"free_days"`
	Status      int       `json:"status" from:"status" db:"status"`
	AppKey      string    `json:"app_key" from:'app_key' db:"app_key"`
}

var webApp WebApp
var webApps []WebApp

type Container struct {
	db.Model
	Name        string            `json:"name" form:"name"gorm:"name"`
	Description string            `json:"description" form:"description"gorm:"description"`
	Cover       string            `json:"cover" form:"cover" gorm:"cover"`
	CreatorId   string            `json:"creator_id" form:"creator_id"`
	Status      int               `json:"status" form:"status"gorm:"status"`
	Permission  int               `json:"permission" form:"permission"gorm:"permission"`
	Members     []ContainerMember `json:"-"`
	Lvl         string            `gorm:"lvl"`
}

var containers []Container
