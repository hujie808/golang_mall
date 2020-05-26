package routes

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"os"
	"strings"
	"web_iris/golang_mall/bootstrap"
	"web_iris/golang_mall/services"
	"web_iris/golang_mall/web/controllers"
	"web_iris/golang_mall/web/middlewares"
)

func Configure(b *bootstrap.Bootstrapper) {

	adminUserService := services.NewAdminUserService()
	mallAddressService := services.NewMallAddressService()
	mallbannerService := services.NewMallBannerService()
	mallLevelService := services.NewMallLevelService()
	mallWechatService := services.NewMallWechatService()
	mallCategoryService := services.NewMallCategoryService()
	mallCommodityService := services.NewMallCommodityService()
	mallskuService := services.NewMallSkuService()
	mallcouponsService := services.NewMallCouponsService()
	mallOrderService := services.NewMallOrderService()
	mallOrderInfoService := services.NewMallOrderInfoService()
	mallShoppingService := services.NewMallShoppingService()
	mallHistoryService := services.NewMallHistoryService()
	mallUserCouponsService := services.NewMallUserCouponsService()
	redLogService := services.NewRedLogService()
	retailLogService := services.NewRetailLogService()
	mallArticleService := services.NewMallArticleService()
	//验证二维码txt
	txt := b.Party("/")
	txt.Get("jubAYWPoqv.txt", func(ctx iris.Context) {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		txt_path := strings.Split(dir, "golang_mall")[0] + "golang_mall/web/public/jubAYWPoqv.txt"
		ctx.ServeFile(txt_path, false)
	})

	//后台设置签名
	adminGift := mvc.New(b.Party("/set_sign"))
	adminGift.Router.Use(identity.CrossDomainMiddle)
	adminGift.Register(
		adminUserService,
	)
	adminGift.Handle(new(controllers.SetSignatureController))
	//用户
	admin := mvc.New(b.Party("/admin_user"))
	admin.Router.Use(identity.CrossDomainMiddle, identity.SignMiddleware)
	admin.Register(
		adminUserService,
	)
	admin.Handle(new(controllers.AdminController))
	//轮播图
	banner := mvc.New(b.Party("/banner"))
	banner.Router.Use(identity.CrossDomainMiddle, identity.SignMiddleware)
	banner.Register(
		mallbannerService,
	)
	banner.Handle(new(controllers.MallBannerController))

	//用户级别
	mall_level := mvc.New(b.Party("/mall_level"))
	mall_level.Router.Use(identity.CrossDomainMiddle, identity.SignMiddleware)
	mall_level.Register(
		mallLevelService,
	)
	mall_level.Handle(new(controllers.MallLevelController))
	//wechat_user
	mallWechat := mvc.New(b.Party("/mall_wechat"))
	mallWechat.Router.Use(identity.CrossDomainMiddle, identity.SignMiddleware)
	mallWechat.Register(
		mallWechatService,
		mallLevelService,
	)
	mallWechat.Handle(new(controllers.MallWechatControlers))
	//分类
	mallcategory := mvc.New(b.Party("/mall_category"))
	mallcategory.Router.Use(identity.CrossDomainMiddle, identity.SignMiddleware)
	mallcategory.Register(
		mallCategoryService,
		mallCommodityService,
	)
	mallcategory.Handle(new(controllers.MallCategoryController))
	//商品
	mallcommodity := mvc.New(b.Party("/mall_commodity"))
	mallcommodity.Router.Use(identity.CrossDomainMiddle, identity.SignMiddleware)
	mallcommodity.Register(
		mallCategoryService,
		mallCommodityService,
		mallskuService,
	)
	mallcommodity.Handle(new(controllers.MallCommodityController))

	//优惠券
	mall_coupons := mvc.New(b.Party("/mall_coupons"))
	mall_coupons.Router.Use(identity.CrossDomainMiddle, identity.SignMiddleware)
	mall_coupons.Register(
		mallcouponsService,
		mallCommodityService,
	)
	mall_coupons.Handle(new(controllers.MallCouponsController))
	//订单
	mall_order := mvc.New(b.Party("/mall_order"))
	mall_order.Router.Use(identity.CrossDomainMiddle, identity.SignMiddleware)
	mall_order.Register(
		mallOrderService,
		mallOrderInfoService,
	)
	mall_order.Handle(new(controllers.MallOrderController))

	//红包展示
	redlog := mvc.New(b.Party("/red_log"))
	redlog.Router.Use(identity.CrossDomainMiddle, identity.SignMiddleware)
	redlog.Register(
		mallOrderService,
		mallOrderInfoService,
		redLogService,
		retailLogService,
	)
	redlog.Handle(new(controllers.MallRedLogController))
	//提成日志页面
	retail_log := mvc.New(b.Party("/retail_log"))
	retail_log.Router.Use(identity.CrossDomainMiddle, identity.SignMiddleware)
	retail_log.Register(
		mallOrderService,
		mallOrderInfoService,
		redLogService,
		retailLogService,
	)
	retail_log.Handle(new(controllers.MallRetailController))

	//文章
	mall_article := mvc.New(b.Party("/mall_article"))
	mall_article.Router.Use(identity.CrossDomainMiddle, identity.SignMiddleware)
	mall_article.Register(
		mallArticleService,
	)
	mall_article.Handle(new(controllers.MallArticleController))

	//微信

	wechatAuth := mvc.New(b.Party("/wechatauth")) //登录验证
	wechatAuth.Router.Use(identity.CrossDomainMiddle)
	wechatAuth.Register(
		mallWechatService,
	)
	wechatAuth.Handle(new(controllers.WechatAuth))

	//index
	mall := mvc.New(b.Party("/mall")) //登录验证
	mall.Router.Use(identity.CrossDomainMiddle, identity.OpenidMiddleware)
	mall.Register(
		mallbannerService,
		mallWechatService,
		mallCategoryService,
		mallCommodityService,
		mallskuService,
		mallcouponsService,
		mallOrderService,
		mallOrderInfoService,
		mallShoppingService,
		mallHistoryService,
		mallLevelService,
		mallUserCouponsService,
		mallAddressService,
		redLogService,
		retailLogService,
	)

	noOpenid := mvc.New(b.Party("/mall")) //登录验证
	noOpenid.Router.Use(identity.CrossDomainMiddle, identity.NoOpenidMiddleware)
	noOpenid.Register(
		mallbannerService,
		mallWechatService,
		mallCategoryService,
		mallCommodityService,
		mallskuService,
		mallcouponsService,
		mallOrderService,
		mallOrderInfoService,
		mallShoppingService,
		mallHistoryService,
		mallLevelService,
		mallUserCouponsService,
		mallAddressService,
		redLogService,
		retailLogService,
		mallArticleService,
	)

	//index
	index := noOpenid.Party("/index")
	index.Handle(new(controllers.WIndex))

	article := noOpenid.Party("/article")
	article.Handle(new(controllers.WArticle))
	//category
	category := mall.Party("/category")
	category.Handle(new(controllers.WCategory))
	//commodity
	commodity_list := mall.Party("/commodity_list")
	commodity_list.Handle(new(controllers.WCommodity))
	//shopping
	shopping := mall.Party("/shopping")
	shopping.Handle(new(controllers.WShopping))
	//my_page
	my_page := mall.Party("/my_page")
	my_page.Handle(new(controllers.WMyPage))
	//address
	address := mall.Party("/address")
	address.Handle(new(controllers.WAddress))

	//WUserCoupons
	usereCoupons := mall.Party("/user_coupons")
	usereCoupons.Handle(new(controllers.WUserCoupons))
	//WMyOrder
	wOrder := mall.Party("/order")
	wOrder.Handle(new(controllers.WMyOrder))

	//分销
	wRetail := mall.Party("/retail")
	wRetail.Handle(new(controllers.WRetail))

	//微信pay
	wechatPay := mvc.New(b.Party("/pay")) //登录验证
	wechatPay.Router.Use(identity.CrossDomainMiddle, identity.NoOpenidMiddleware)
	wechatPay.Register(
		mallWechatService,
		mallCommodityService,
		mallskuService,
		mallcouponsService,
		mallOrderService,
		mallOrderInfoService,
		mallShoppingService,
		mallLevelService,
		redLogService,
		retailLogService,
	)
	wechatPay.Handle(new(controllers.WechatPay))
	wechatPay.Router.Use(identity.CrossDomainMiddle)
	CheckPay := mvc.New(b.Party("/pay/check_wxpay/")) //登录验证
	CheckPay.Register(
		mallWechatService,
		mallCommodityService,
		mallskuService,
		mallcouponsService,
		mallOrderService,
		mallOrderInfoService,
		mallShoppingService,
		mallLevelService,
		redLogService,
		retailLogService,
	)
	CheckPay.Handle(new(controllers.CheckPay))

}
