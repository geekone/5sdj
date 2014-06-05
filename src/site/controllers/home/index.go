package home

import(
	"fmt"
	// "strconv"
	"time"		
	 // "github.com/astaxie/beego/logs"
	 "github.com/astaxie/beego/cache"
	 _ "github.com/astaxie/beego/cache/memcache"
)

type IndexController struct {
	baseController
}

func (this *IndexController) Index(){
	//log := logs.NewLogger(10000)
	//log.EnableFuncCallDepth(true)				//显示行号和文件
	//log.SetLogger("console","")				//默认trace 
	//log.SetLogger("console",`{"level":1}`)	//从debug开始显示
	//log.SetLogger("console",`{"level":2}`)		//从info开始
	
	//log.SetLogger("file",`{"filename":"site.log"}`)	//定义文件

	//log.Trace("trace %s %s","param1","param2")
	//log.Debug("debug")
	//log.Info("info")
	//log.Warn("warning")
	//log.Error("error")
	//log.Critical("critical")

	var t string

	// bm,_:= cache.NewCache("memory",`{"interval":60}`)

	// // if bm.IsExist("astaxie") {
	// if bm.Get("astaxie") != nil{
		
	// 	v := bm.Get("astaxie")
	// 	t = v.(string)
	// }else{

	// 	v := time.Now().Format("2006-01-02 15:04:05")
	// 	t = string(v)
	// 	bm.Put("astaxie",t,10)
	// }

	//TODO File Cache 有问题
	// bm, err := cache.NewCache("file", `{"CachePath":"/cache","FileSuffix":".bin","DirectoryLevel":2,"EmbedExpiry":120}`)
	// if err != nil {
	// 	fmt.Println("init err")
	// }

	// if bm.IsExist("astaxie") {
	// 	v := bm.Get("astaxie")
	// 	t = v.(string)
	// }else{
	// 	v := time.Now().Format("2006-01-02 15:04:05")
	// 	t = string(v)
	// 	bm.Put("astaxie",t,10)
	// }

		
	//TODO memcache 如果关闭服务。怎么解决挂了
	bm,err := cache.NewCache("memcache",`{"conn":"127.0.0.1:11211"}`)

	if err != nil {
		fmt.Println("init err")
	}

	// if bm.IsExist("astaxie") {
	if bm.Get("astaxie") != nil{
		
		v := bm.Get("astaxie")
		t = v.(string)
	}else{

		v := time.Now().Format("2006-01-02 15:04:05")
		t = string(v)
		bm.Put("astaxie",t,10)
	}


	this.Data["time"] = t
	this.display()
}