package main

import (
	"context"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/ops-cn/go-devops/admin/app"
	"github.com/ops-cn/go-devops/common/logger"
)

// VERSION 版本号，可以通过编译的方式指定版本号：go build -ldflags "-X main.VERSION=x.x.x"
var VERSION = "6.3.1"

func main() {
	logger.SetVersion(VERSION)
	ctx := logger.NewTraceIDContext(context.Background(), "main")

	service := micro.NewService(
		micro.Name("ops-cn.admin"),
		micro.Flags(
			&cli.StringFlag{
				Name:     "conf",
				Aliases:  []string{"c"},
				Usage:    "配置文件(.json,.yaml,.toml)",
				Required: true,
			},
			//&cli.StringFlag{
			//	Name:     "model",
			//	Aliases:  []string{"m"},
			//	Usage:    "casbin的访问控制模型(.conf)",
			//	Required: true,
			//},
			//&cli.StringFlag{
			//	Name:  "menu",
			//	Usage: "初始化菜单数据配置文件(.yaml)",
			//},
			//&cli.StringFlag{
			//	Name:  "www",
			//	Usage: "静态站点目录",
			//},
		),
		micro.Address(":10087"),
	)
	service.Init(
		micro.Action(func(c *cli.Context) error {
			app.Init(
				ctx,
				app.SetConfigFile(c.String("conf")),
				app.SetModelFile(c.String("model")),
				//app.SetWWWDir(c.String("www")),
				//app.SetMenuFile(c.String("menu")),
				app.SetVersion(VERSION))
			return nil
		}),
	)
	app.SetHandler(service.Server())
	//err := adminPB.RegisterLoginHandler(service.Server(), new(impl.Login))
	//if err != nil {
	//	panic(err)
	//}
	if err := service.Run(); err != nil {
		panic(err)
	}

}

/*func newWebCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "web",
		Usage: "运行web服务",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conf",
				Aliases:  []string{"c"},
				Usage:    "配置文件(.json,.yaml,.toml)",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "model",
				Aliases:  []string{"m"},
				Usage:    "casbin的访问控制模型(.conf)",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "menu",
				Usage: "初始化菜单数据配置文件(.yaml)",
			},
			&cli.StringFlag{
				Name:  "www",
				Usage: "静态站点目录",
			},
		},
		Action: func(c *cli.Context) error {
			return app.Run(ctx,
				app.SetConfigFile(c.String("conf")),
				app.SetModelFile(c.String("model")),
				app.SetWWWDir(c.String("www")),
				app.SetMenuFile(c.String("menu")),
				app.SetVersion(VERSION))
		},
	}
}*/
