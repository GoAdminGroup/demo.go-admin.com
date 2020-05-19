package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	ada "github.com/GoAdminGroup/go-admin/adapter/gin"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	_ "github.com/GoAdminGroup/themes/adminlte"
	_ "github.com/GoAdminGroup/themes/sword"

	"github.com/GoAdminGroup/components/echarts"
	"github.com/GoAdminGroup/demo_en/login"
	"github.com/GoAdminGroup/demo_en/pages"
	"github.com/GoAdminGroup/demo_en/tables"
	"github.com/GoAdminGroup/filemanager"
	adminContext "github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/librarian"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	eng := engine.Default()

	template.AddLoginComp(login.Get())
	template.AddComp(chartjs.NewChart())
	template.AddComp(echarts.NewChart())

	rootPath := "/data/www/go-admin-en"
	//rootPath = "."

	cfg := config.ReadFromJson(rootPath + "/config.json")
	cfg.CustomFootHtml = template.HTML(`<div style="display:none;">
    <script type="text/javascript" src="https://s9.cnzz.com/z_stat.php?id=1278156902&web_id=1278156902"></script>
	<!-- Global site tag (gtag.js) - Google Analytics -->
	<script async src="https://www.googletagmanager.com/gtag/js?id=UA-103003647-2"></script>
	<script>
	window.dataLayer = window.dataLayer || [];
	function gtag(){dataLayer.push(arguments);}
	gtag('js', new Date());
	gtag('config', 'UA-103003647-2');
	</script>
</div>`)
	cfg.CustomHeadHtml = template.HTML(`<link rel="icon" type="image/png" sizes="32x32" href="//quick.go-admin.cn/official/assets/imgs/icons.ico/favicon-32x32.png">
        <link rel="icon" type="image/png" sizes="96x96" href="//quick.go-admin.cn/official/assets/imgs/icons.ico/favicon-64x64.png">
        <link rel="icon" type="image/png" sizes="16x16" href="//quick.go-admin.cn/official/assets/imgs/icons.ico/favicon-16x16.png">`)

	cfg.Animation = config.PageAnimation{
		Type:     "fadeInUp",
		Duration: 0.9,
	}
	//cfg.HideConfigCenterEntrance = true
	cfg.AddUpdateProcessFn(func(values form.Values) (values2 form.Values, e error) {
		if values.Get("theme") == "adminlte" && values.Get("asset_url") == "//quick.go-admin.cn/demo/sword" {
			values.Add("asset_url", "//quick.go-admin.cn/demo")
		}
		if values.Get("theme") == "sword" && values.Get("asset_url") == "//quick.go-admin.cn/demo" {
			values.Add("asset_url", "//quick.go-admin.cn/demo/sword")
		}
		if values.Get("site_off") == "true" || values.Get("no_limit_login_ip") == "false" {
			return nil, errors.New("permission denied")
		}
		if values.Get("login_title") != "GoAdmin" {
			return nil, errors.New("permission denied")
		}
		if values.Get("custom_head_html") != string(cfg.CustomHeadHtml) {
			return nil, errors.New("permission denied")
		}
		if values.Get("custom_foot_html") != string(cfg.CustomFootHtml) {
			return nil, errors.New("permission denied")
		}
		if values.Get("footer_info") != "" || values.Get("login_logo") != string(cfg.LoginLogo) ||
			values.Get("logo") != string(cfg.Logo) || values.Get("mini_logo") != string(cfg.MiniLogo) {
			return nil, errors.New("permission denied")
		}
		return values, nil
	})

	if err := eng.AddConfig(cfg).
		AddGenerators(tables.Generators).
		AddGenerator("user", tables.GetUserTable).
		AddPlugins(filemanager.NewFileManagerWithConfig(filemanager.Config{
			Path:          "/data/www/go-admin-en/fm_example",
			AllowDelete:   false,
			AllowUpload:   true,
			AllowRename:   true,
			AllowDownload: true,
			AllowMove:     true,
		}), librarian.NewLibrarianWithConfig(librarian.Config{
			Path:      "/data/www/go-admin-en/fm_example/markdown",
			BuildMenu: false,
			Prefix:    "librarian",
		})).
		AddNavButtons("Website Info", "", action.PopUp("/website/info", "Website Info",
			func(ctx *adminContext.Context) (success bool, msg string, data interface{}) {
				return true, "ok", `<p>Created by <a href="https://github.com/chenhg5">cg33<a/></p>`
			})).
		AddNavButtons("Manager", "", action.Jump("/admin/info/manager")).
		Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", rootPath+"/uploads")

	// you can custom your pages like:

	r.GET("/admin", ada.Content(func(ctx *gin.Context) (panel types.Panel, e error) {
		if config.GetTheme() == "adminlte" {
			return pages.GetDashBoardContent(ctx)
		} else {
			return pages.GetDashBoard2Content(ctx)
		}
	}))
	r.GET("/admin/echarts", ada.Content(pages.GetDashBoard3Content))
	r.GET("/admin/form1", ada.Content(pages.GetForm1Content))
	r.GET("/admin/table", ada.Content(pages.GetTableContent))

	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/admin")
	})

	eng.Data("POST", "/admin/form/update", func(ctx *adminContext.Context) {
		fmt.Println("ctx.PostForm()", ctx.PostForm())
		ctx.PjaxUrl("/admin")
	})

	srv := &http.Server{
		Addr:    ":9033",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

	log.Print("closing database connection")
	eng.MysqlConnection().Close()
}
