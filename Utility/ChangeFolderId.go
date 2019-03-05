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
	var comments []Comment
	db.Order("created_at asc").Find(&comments)
	glog.Info("Sum of Comments:", len(comments))
	for index, comment := range comments {
		//find record in table viewpoints
		var viewpoint Viewpoint
		if db.Where("id = ?", comment.ViewpointId).First(&viewpoint).RecordNotFound() {
			glog.Warningf("viewpointId %s was not found in viewpoints.continue.", comment.ViewpointId)
			//whether filed of viewpoint_id was updated to file_id
			if !db.Table("project_files").Where("id = ?", comment.ViewpointId).RecordNotFound() {
				glog.Infof("viewpointId %s was found in project_files %d", comment.ViewpointId, index+1)
				continue
			} else {
				glog.Infof("viewpointId %s was not found in project_files,exit.", comment.ViewpointId)
				return
			}
		}
		//rename & check attachments
		var attachment Attachment
		if !db.Where("parent_id = ?", comment.Id).First(&attachment).RecordNotFound() {
			// attachment existed & modify corresponding folder
			if err := Rename(viewpoint.ProjectId, comment.ViewpointId, viewpoint.ProjectFileId); err != nil {
				glog.Warning(path2.Join(CommentPath, viewpoint.ProjectId, comment.ViewpointId), " was not found.")
				//whether had alread been modified.
				newFilePath := path2.Join(CommentPath, viewpoint.ProjectId, viewpoint.ProjectFileId)
				if _, err := os.Open(newFilePath); err != nil && os.IsNotExist(err) {
					glog.Infof(" %s was not found %d ,exit.", newFilePath, index+1)
					return
				} else {
					//modirying in previous round
					glog.Infof(" %s was  found %d", newFilePath, index+1)
				}
			} else {
				glog.Infof("%s Modify Folder Name:%s -> %s", comment.Id, comment.ViewpointId, viewpoint.ProjectFileId)
			}
		}

		if db.Table("comments").Where("id = ?", comment.Id).Update("viewpoint_id", viewpoint.ProjectFileId).RecordNotFound() {
			glog.Fatal("fail to update viewpoint id in comments")
			return
		}
		glog.Infof("%s Update viewpoint_id %s -> %s", comment.Id, comment.ViewpointId, viewpoint.ProjectFileId)

		glog.Infof("%3d/%-3d", index+1, len(comments))
	}

}
func Rename(pid, viewpointid, fid string) error {
	oldpath := path2.Join(CommentPath, pid, viewpointid)
	newpath := path2.Join(CommentPath, pid, fid)
	err := os.Rename(oldpath, newpath)
	return err

}
