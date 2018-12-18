package models

import (
	"database/sql"

	"github.com/astaxie/beego"
	_ "github.com/mattn/go-sqlite3"
)

// DataEvent structure for data send to web interface
type DataEvent struct { // structure for data send to web interface
	PackageNumber int // time stamp
	Address       string
	HexString     string
	ASCIIString   string
	TimeStamp     string
}

// RecMessage structure for data received from web
type RecMessage struct {
	Name     string
	Port     int
	Action   string
	Protocol string
	Data     string
	SaveFlag int
}

type saveMessage struct {
	PackageNumber int
	TimeStamp     string
	ASCIIString   string
	HexString     string
}

var endFlag chan int
var saveData chan saveMessage

// InitSaveChan initialize channel variable
func InitSaveChan() {
	endFlag = make(chan int)
	saveData = make(chan saveMessage)
}

// TranSaveChan  save data to channel variable
func TranSaveChan(data DataEvent) {
	var msg saveMessage
	msg.PackageNumber = data.PackageNumber
	msg.TimeStamp = data.TimeStamp
	msg.ASCIIString = data.ASCIIString
	msg.HexString = data.HexString
	saveData <- msg
}

//EndSaveTask  set channel to end save task
func EndSaveTask() {
	endFlag <- 1

}

//CreateDataSaveTask create saving data task
func CreateDataSaveTask(filename string) error {

	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		beego.Error(err)
		return err
	}
	defer db.Close()
	sqlExec := `create table DATA (packageNumber integer not null primary key,TimeStamp text
		 ,ASCIIString text,HexString text)`
	_, err = db.Exec(sqlExec)
	if err != nil {
		beego.Error(err)
		return err
	}

	stmt, err := db.Prepare(`insert into data(packageNumber,TimeStamp,ASCIIString,HexString)
	               values(?,?,?,?)`)
	if err != nil {
		beego.Error(err)
		return err
	}
	defer stmt.Close()
	for {
		select {
		case data := <-saveData:
			_, err = stmt.Exec(data.PackageNumber, data.TimeStamp, data.ASCIIString, data.HexString)
			if err != nil {
				beego.Error(err)
				return err
			}

		case <-endFlag:
			return nil
		}
	}
}
