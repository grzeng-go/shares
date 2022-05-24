package main

import (
	"fmt"
	"os"
	"shares/internal/config"
	"shares/internal/routers"

	"github.com/gin-gonic/gin"
	"github.com/gmsec/goplugins/plugin"
	"github.com/xxjwxc/public/mydoc/myswagger"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/server"
)

// CallBack service call backe
func CallBack() {
	mylog.SetLog(mylog.GetDefaultZap())

	// swagger
	myswagger.SetHost("https://localhost:" + config.GetPort())
	myswagger.SetBasePath("shares")
	myswagger.SetSchemes(true, false)
	// -----end --
	// event.UPPPP()

	// mylog.Error(config.GetMaxCapacity())
	// reg := registry.NewDNSNamingRegistry()
	// reg := etcdv3.NewEtcdv3NamingRegistry(clientv3.Config{
	// 	Endpoints:   config.GetEtcdInfo().Addrs,
	// 	DialTimeout: time.Second * time.Duration(config.GetEtcdInfo().Timeout),
	// })
	// grpc 相关 初始化服务
	// service := micro.NewService(
	// 	micro.WithName("gmsec.srv.shares"),
	// 	// micro.WithRegisterTTL(time.Second*30),      //指定服务注册时间
	// 	micro.WithRegisterInterval(time.Second*15), //让服务在指定时间内重新注册
	// 	micro.WithRegistryNaming(reg),
	// )
	// ----------- end

	// gin restful 相关
	router := gin.Default()
	router.Use(routers.Cors())
	v1 := router.Group("/shares/api/v1")
	routers.OnInitRoot(nil, v1) // 自定义初始化
	// ------ end

	plg, b := plugin.RunHTTP(plugin.WithGin(router),
		// plugin.WithMicro(service),
		plugin.WithAddr(":"+config.GetPort()))

	if b == nil {
		plg.Wait()
	}
	fmt.Println("done")
}

/*func init() {
	core.Dao.GetDBr().DB.AutoMigrate(&model.AnalyCdTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.AnalyDbszxTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.AnalyDwflTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.AnalyFlTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.AnalyHpTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.AnalyHyTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.AnalyLhbTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.AnalyTkTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.AnalyUpTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.APITbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.BillRefundTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.BillTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.BsRapidlyTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.GroupListTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.GroupTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.GroupWatchTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.HyDailyTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.HyInfoTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.HyUpTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.LhbDailyTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.LogTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.MaCdTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.MaFlTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.MaLhbTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.MaUpTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.MsgRapidlyTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.MsgTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.MyselfTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.NoPegTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.NoTdxTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.NoTradingDayTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.SharesDailyTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.SharesInfoTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.SharesWatchTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.TdxDailyTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.TencentTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.TkTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.WxMsgTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.WxUserinfo{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.XqMsgDailyActiveTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.XqMsgDailyTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.XqMsgTbl{})
	core.Dao.GetDBr().DB.AutoMigrate(&model.ZljlrDailyTbl{})
}*/

func main() {
	if config.GetIsDev() || len(os.Args) == 0 {
		CallBack()
	} else {
		server.On(config.GetServiceConfig()).Start(CallBack)
	}
}
