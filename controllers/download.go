package controllers
import (
	"fmt"
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
	FileName string
	FilePath string
	ModTime  string
}
var filePath string

func SetFilePath (filepath string) {
	filePath=filepath
}

func ListFileHandler(c *gin.Context){

	filenames :=getDirFileList(filePath,false)

	c.HTML(http.StatusOK, "file.html", gin.H{
		"filenames": filenames,
	})
}

func DownloadFileHandler(c *gin.Context)  {

	filename:=c.Param("id")
	len:=len(filename)
	filename=filename[1:len]
	log.Println("filename is ",filename)
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(filename)

	//filenames :=getDirFileList(filePath,false)

	//c.HTML(http.StatusOK, "file.html", gin.H{
	//	"filenames": filenames,
	//})
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
			filepropertylist[location].FileName=f.Name()
			filepropertylist[location].FilePath= "download/"+dirname+"/"+f.Name()
			filepropertylist[location].ModTime=f.ModTime().String()
			location++
		} else if (dirflag == false) && (f.IsDir() == false) {
			filepropertylist[location].FileName=f.Name()
			filepropertylist[location].FilePath= "download/"+dirname+"/"+f.Name()
			filepropertylist[location].ModTime=f.ModTime().String()
			location++
		} else {
			continue
		}
	}
	filepropertylist=filepropertylist[:location]



	return filepropertylist
}