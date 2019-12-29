package controllers
import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)
type FileItem struct {
	fileName string
	dirFlag  bool
	modTime  time.Time
}

type FileProperty struct {
	fileName string

	modTime  string
}
func FileDownHandler(c *gin.Context){
	filenames :=getDirFileList("datafile",false)
	
	c.HTML(http.StatusOK, "file.html", gin.H{
		"filenames": filenames,
	})   	
}



func getDirFileList(dirname string, dirflag bool)  []FileProperty {
	files,err := ioutil.ReadDir(dirname)
	filepropertylist := make([]FileProperty,1024)
	if err!=nil {
		log.Println("Read file error is ",err.Error())
		return nil 
	}
	location:=0
	for _,f:=range files {
		if (dirflag == true) && (f.IsDir()==true) {
			filepropertylist[location].fileName=f.Name()
			filepropertylist[location].modTime=f.ModTime().String()
			location++
		} else if (dirflag == false) && (f.IsDir() == false) {
			filepropertylist[location].fileName=f.Name()
			filepropertylist[location].modTime=f.ModTime().String()
			location++			
		} else {
			continue
		}
	}
	filepropertylist=filepropertylist[:location]
	

	
	return filepropertylist
}