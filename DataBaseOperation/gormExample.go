package DataBaseOperation

import (
	"dx/taishan/core/db"
	"dx/taishan/modules/user/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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
type RawFile struct {
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
	Name        string       `form:"name" binding:"exists,alphanum,min=4,max=255" json:"name"`
	//Password    string       `form:"password" binding:"exists,min=8,max=255" json:"-"`
	//Email       string       `form:"email" binding:"email" json:"email"`
	Phone       string       `form:"phone" json:"phone"`
	Status      int          `form:"-" json:"status"`
	//Profile     UserProfile  `form:"-" gorm:"foreignkey:UserId" json:"profile"`
	//Company     Company      `form:"-" gorm:"foreignkey:UserId" json:"company"`
	//Roles       []UserRole   `json:"-" gorm:"many2many:user_user_roles"`
	//Oauth       models.Oauth `form:"-" gorm:"foreignkey:UserId" json:"oauth"`
	//ActivatedAt time.Time
	//CompanyId   string      `form:"company_id" json:"company_id" db:"company_id"`
	//AreaCode    string      `form:"area_code" json:"area_code"`
	//Label       []UserLabel `form:"-" json:"label" gorm:"foreignkey:UserId"`
	//UserId      string     `form:"user_id" binding:"exists" json:"user_id" db:"user_id" gorm:"user_id;size=32;require;unique;index"`
	Avatar      string     `form:"-" json:"avatar" db:"avatar"`
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

func RungOrm() {

	//var CMdata []ContainerMember
	//var RAWfiles []RawFile
	userId := "2fd8584564ad47798ac4d23cf4a03ea1"
	db, err := gorm.Open("mysql", "root:123@/taishan_dev?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	//db.Where("owner_id = ? AND member_type=? ", userId, 1).Find(&CMdata)
	//fmt.Println(CMdata)

	//v, _ := db.Table("container_members").Select("id").Where("owner_id = ? AND member_type=? ", userId, 1).Rows()

	//var ids []struc
	//var ids []string
	//v, _ := db.Table("container_members").Select("id").Where("owner_id = ? AND member_type=? ", userId, 1).Rows()
	//v, err := db.Table("container_members").Select("id").Where("owner_id = ? AND member_type=? ", userId, 1).Rows()
	//if v.Next() {
	//	var id string
	//	v.Scan(&id)
	//	ids = append(ids, id)
	//	fmt.Println(id)
	//}

	//ids=[]string{"c4898f0f42d844579e395762d0099c0b","f24445cba3d64bf4bb70f16a20628835"}

	//v,_=db.Table("raw_files").Select("id").Where("uploader_id=?", ids).Rows()
	//if v.Next(){
	//
	//}
	//fmt.Println(ids)
	//db.Table("raw_files").Where("uploader_id = ?", "1dc6e02381114f26aba01fc89868200a").Find(&RAWfiles)

	//db.Table("raw_files").Where("uploader_id=?")
	//var files []FileStruct
	//db.Table("files").Joins("JOIN raw_files on raw_files.id = files.raw_file_id").Where("raw_files.uploader_id in (?)",ids).Find(&files)
	//db.Table("files").Where("raw_file_id =? ",).Find(&fss)

	//fmt.Println(files)

	//db.Table("raw_files").Where("uploader_id = ?", ids).Find(&RAWfiles)
	//fmt.Println(RAWfiles)

	//var profiles []UserProfile

	//v, err := db.Table("users").Select("id").Where("id in (?)", db.Table("container_members").Select("owner_id").Where("container_id in (?) and member_type = ?", db.Table("containers").Select("id").Where("creator_id = ?", userId).QueryExpr(), 2).QueryExpr()).Rows()
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//var ids []string
	//if v.Next() {
	//	var id string
	//	v.Scan(&id)
	//	ids = append(ids, id)
	//}
	//fmt.Println(ids)

	var result []UserForProfiles
	//var result2 []UserProfile
	//var result3 []User
	err = db.Table("users").Select("*").Joins("JOIN user_profiles on user_profiles.user_id=users.id").Where("users.id in (?)", db.Table("container_members").Select("owner_id").Where("container_id in (?) and member_type = ?", db.Table("containers").Select("id").Where("creator_id = ?", userId).QueryExpr(), 2).QueryExpr()).Find(&result).Error

	//err = db.Table("users").Select("*").Joins("left join user_profiles on user_profiles.user_id=users.id").Where("users.id in (?)", "dbd09cf3d711455d8c69421225d5eb2b").Find(&result).Error
	//err = db.Table("user_profiles").Where("user_id in (?)", "dbd09cf3d711455d8c69421225d5eb2b").Find(&result2).Error
	fmt.Println(result)

	//p := fmt.Println
	//p(result[0].Name)
	//p(result[0].Id)
	//p(result[0].Avatar)
	//p(result[0].Email)
	//p(result[0].CompanyId)
	//p(result[0].Status)
}

func FuncOne(ip *int) {

}
