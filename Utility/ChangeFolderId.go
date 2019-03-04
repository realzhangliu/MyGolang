package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"os"
	path2 "path"
)

var CommentPath string

const PATH102 = "/media/dx/Code/Data/taishan-data/data/comment_files"

func connectDB() *gorm.DB {
	var dbName string
	var user, host, pwd string
	var port int
	flag.StringVar(&dbName, "db", "taishan_dev", "-db taishan_dev2")
	flag.StringVar(&user, "u", "root", "-u root")
	flag.StringVar(&pwd, "p", "123", "-p 123")
	flag.IntVar(&port, "P", 3306, "-P 3306")
	flag.StringVar(&host, "h", "192.168.99.102", "-h localhost")
	flag.StringVar(&CommentPath, "d", "/media/dx/Code/Data/taishan-data/data/comment_files", "-d /media/dx/Code/Data/taishan-data/data/comment_files")
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
		user,
		pwd,
		host,
		port,
		dbName)
	flag.Set("logtostderr", "true")
	flag.Parse()
	glog.Info(dbUrl)

	var DB *gorm.DB
	var err error
	if DB, err = gorm.Open("mysql", dbUrl); err != nil {
		glog.Fatal(err)
	}

	//if _, err = os.Open(CommentPath); os.IsNotExist(err) {
	//	glog.Fatal(err)
	//}

	glog.Info(CommentPath)

	return DB
}

func main() {
	db := connectDB()

	//var ids []string
	var attachments []Attachment
	db.Find(&attachments)
	glog.Info("Sum of Attachmetns:", len(attachments))
	for _, attachment := range attachments {
		row, err := db.Model(&Comment{}).Where("id = ? ", attachment.ParentId).Rows()
		glog.Info(err)
		for row.Next() {
			var Id string
			var ViewpointId string
			row.Scan(&Id, &ViewpointId)
			var viewpoint Viewpoint
			if db.Where("id = ?", ViewpointId).First(&viewpoint).RecordNotFound() {
				glog.Warning("RecoreNotFound")
				continue
			}
			//rename
			if err := Rename(viewpoint.ProjectId, ViewpointId, viewpoint.ProjectFileId); err != nil {
				glog.Fatal(err)
				return
			}

			if db.Table("comments").Where("id = ?", Id).Update("viewpoint_id", viewpoint.ProjectFileId).RecordNotFound() {
				glog.Fatal("fail to update viewpoint id in comments")
				return
			}

			glog.Info("comments: ", "ID:", Id, " ", "ViewPointID:", ViewpointId, "->", viewpoint.ProjectFileId)
		}

	}

}
func Rename(pid, viewpointid, fid string) error {
	oldpath := path2.Join(CommentPath, pid, viewpointid)
	newpath := path2.Join(CommentPath, pid, fid)
	err := os.Rename(oldpath, newpath)
	return err

}
