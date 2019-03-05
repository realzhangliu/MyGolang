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

	if _, err = os.Open(CommentPath); os.IsNotExist(err) {
		glog.Fatal(err)
	}

	glog.Info(CommentPath)

	return DB
}

func main() {
	db := connectDB()

	//var ids []string
	var attachments []Attachment
	db.Find(&attachments)
	var comments []Comment
	db.Find(&comments)
	glog.Info("Sum of Attachmetns:", len(attachments))
	for index, attachment := range attachments {
		row, err := db.Model(&Comment{}).Where("id = ? and deleted_at is null", attachment.ParentId).Select("id,viewpoint_id").Rows()
		if err != nil {
			glog.Info(err)
		}
		for row.Next() {
			var Id string
			var ViewpointId string
			err := row.Scan(&Id, &ViewpointId)
			if err != nil {
				glog.Info(err)
			}
			var viewpoint Viewpoint
			if db.Where("id = ?", ViewpointId).First(&viewpoint).RecordNotFound() {
				glog.Warning(ViewpointId, " Viewpoint RecoreNotFound continue.")
				if !db.Table("project_files").Where("id = ?", ViewpointId).RecordNotFound() {
					glog.Infof("%s was found in project_files", ViewpointId)
				}
				continue
			}
			//rename
			if err := Rename(viewpoint.ProjectId, ViewpointId, viewpoint.ProjectFileId); err != nil {
				glog.Warning(path2.Join(CommentPath, viewpoint.ProjectId, ViewpointId), "  not found.")
				newFilePath := path2.Join(CommentPath, viewpoint.ProjectId, viewpoint.ProjectFileId)
				if _, err := os.Open(newFilePath); err != nil && os.IsNotExist(err) {
					glog.Info(newFilePath, " not found.")
					return
				} else {
					//modirying in previous round
					glog.Info(newFilePath, " was found.")

				}
			}
			glog.Infof("%s Modify DIR Name:%s -> %s", Id, ViewpointId, viewpoint.ProjectFileId)

			if db.Table("comments").Where("id = ?", Id).Update("viewpoint_id", viewpoint.ProjectFileId).RecordNotFound() {
				glog.Fatal("fail to update viewpoint id in comments")
				return
			}
			glog.Infof("%s Update viewpoint_id %s -> %s", Id, ViewpointId, viewpoint.ProjectFileId)

		}
		glog.Infof("%3d/%-3d", index+1, len(attachments))
	}

}
func Rename(pid, viewpointid, fid string) error {
	oldpath := path2.Join(CommentPath, pid, viewpointid)
	newpath := path2.Join(CommentPath, pid, fid)
	err := os.Rename(oldpath, newpath)
	return err

}
