package models

/**
 * models 的基类,提供初始化和一个工具类方法
 */


import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//TODO 有补包
)


//初始运行函数
func init(){
	//从配置文件中取出
	// dbhost := beego.AppConfig.String("dbhost")
	// dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	//生成 connection string
	// dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	dsn := dbuser + ":" + dbpassword + "@/" + dbname + "?charset=utf8"
	// orm.RegisterDataBase("default","mysql",dsn)		//orm注册数据库
	maxIdle := 30
	maxConn := 30
	//orm.RegisterDataBase("default", "mysql", "root:@/5sdj?charset=utf8", maxIdle, maxConn)
	orm.RegisterDataBase("default", "mysql", dsn, maxIdle, maxConn)
	orm.RegisterModel(
		new(User),
		new(BlogCategory),
		new(BlogPost),
		new(Article),
		new(Category),
		new(Joke),
		new(Tag),
		new(TagArticle),
		new(Photo),
		new(Movie),
	)					//orm注册模型


	//在程序中直接调用自动建表：

	// 数据库别名
	name := "default"

	// drop table 后再建表 true
	force := false	

	// 打印执行过程
	verbose := true

	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
	    fmt.Println(err)
	}

	
}


//加密函数
func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x",hash.Sum(nil))
}


//TODO 有补函数


//重新返回带前缀的表名的函数 前缀取自配置文件 str 表名
func TableName(str string) string {
	return fmt.Sprintf("%s%s",beego.AppConfig.String("dbprefix"),str)
}