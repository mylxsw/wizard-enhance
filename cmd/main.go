package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mylxsw/asteria/formatter"
	"github.com/mylxsw/asteria/level"
	"github.com/mylxsw/asteria/log"
	"github.com/mylxsw/asteria/writer"
	"github.com/mylxsw/container"
	"github.com/mylxsw/glacier/event"
	"github.com/mylxsw/glacier/infra"
	"github.com/mylxsw/glacier/starter/application"
	"github.com/mylxsw/wizard-enhance/internal/service"
	"github.com/urfave/cli"
	"github.com/urfave/cli/altsrc"

	"github.com/mylxsw/wizard-enhance/api"
	"github.com/mylxsw/wizard-enhance/config"
	localEvt "github.com/mylxsw/wizard-enhance/internal/event"
	"github.com/mylxsw/wizard-enhance/internal/scheduler"
)

var Version = "1.0"
var GitCommit = "5dbef13fb456f51a5d29464d"

func main() {
	app := application.Create(fmt.Sprintf("%s %s", Version, GitCommit))
	app.AddFlags(altsrc.NewStringFlag(cli.StringFlag{
		Name:   "listen",
		Usage:  "服务监听地址",
		Value:  "127.0.0.1:19921",
		EnvVar: "WIZARD_ENHANCE_LISTEN",
	}))
	app.AddFlags(altsrc.NewBoolFlag(cli.BoolFlag{
		Name:   "debug",
		Usage:  "是否使用调试模式，调试模式下，静态资源使用本地文件",
		EnvVar: "WIZARD_ENHANCE_DEBUG",
	}))
	app.AddFlags(altsrc.NewStringFlag(cli.StringFlag{
		Name:  "log_path",
		Usage: "日志文件输出目录（非文件名），默认为空，输出到标准输出",
	}))
	app.AddFlags(altsrc.NewStringFlag(cli.StringFlag{
		Name:   "secret",
		Usage:  "API 访问秘钥",
		EnvVar: "WIZARD_ENHANCE_SECRET",
	}))
	app.AddFlags(altsrc.NewStringFlag(cli.StringFlag{
		Name:   "db_conn_str",
		Usage:  "数据库连接字符串",
		EnvVar: "WIZARD_ENHANCE_DB_CONN",
		Value:  "root:@tcp(127.0.0.1:3306)/wizard?parseTime=true",
	}))
	app.AddFlags(altsrc.NewStringFlag(cli.StringFlag{
		Name:   "storage_path",
		Usage:  "Wizard 文件存储目录 storage 文件夹所在目录",
		EnvVar: "WIZARD_STORAGE_PATH",
		Value:  "/Users/mylxsw/codes/github/wizard/public",
	}))
	app.AddFlags(altsrc.NewStringFlag(cli.StringFlag{
		Name:   "gotenberg_server",
		Usage:  "PDF 转换服务器地址",
		EnvVar: "WIZARD_GOTENBERG_SERVER",
		Value:  "http://localhost:3000",
	}))

	app.BeforeServerStart(func(cc container.Container) error {
		stackWriter := writer.NewStackWriter()
		cc.MustResolve(func(c infra.FlagContext) {
			if !c.Bool("debug") {
				log.All().LogLevel(level.Info)
			}

			logPath := c.String("log_path")
			if logPath == "" {
				stackWriter.PushWithLevels(writer.NewStdoutWriter())
				return
			}

			log.All().LogFormatter(formatter.NewJSONWithTimeFormatter())
			stackWriter.PushWithLevels(writer.NewDefaultRotatingFileWriter(func(le level.Level, module string) string {
				return filepath.Join(logPath, fmt.Sprintf("%s-%s.log", time.Now().Format("20060102"), le.GetLevelName()))
			}))
		})

		stackWriter.PushWithLevels(
			NewErrorCollectorWriter(app.Container()),
			level.Error,
			level.Emergency,
			level.Critical,
		)
		log.All().LogWriter(stackWriter)

		return nil
	})

	app.Singleton(func(c infra.FlagContext) *config.Config {
		return &config.Config{
			Version:         Version,
			GitCommit:       GitCommit,
			Listen:          c.String("listen"),
			Debug:           c.Bool("debug"),
			LogPath:         c.String("log_path"),
			APISecret:       c.String("secret"),
			DBConnStr:       c.String("db_conn_str"),
			GotenbergServer: c.String("gotenberg_server"),
			StoragePath:     c.String("storage_path"),
		}
	})
	app.Singleton(func(conf *config.Config) (*sql.DB, error) {
		return sql.Open("mysql", conf.DBConnStr)
	})
	app.Singleton(func() *redis.Client {
		return redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		})
	})

	app.BeforeServerStop(func(cc infra.Resolver) error {
		return cc.Resolve(func(em event.Publisher) {
			em.Publish(localEvt.SystemUpDownEvent{
				Up:        false,
				CreatedAt: time.Now(),
			})
		})
	})

	app.Main(func(conf *config.Config, em event.Publisher) {
		if log.DebugEnabled() {
			log.WithFields(log.Fields{
				"config": conf,
			}).Debug("configuration")
		}

		em.Publish(localEvt.SystemUpDownEvent{
			Up:        true,
			CreatedAt: time.Now(),
		})
	})

	app.Provider(api.Provider{}, localEvt.Provider{}, scheduler.Provider{})
	app.Provider(service.Provider{})

	if err := app.Run(os.Args); err != nil {
		log.Errorf("exit with error: %s", err)
	}
}

type ErrorCollectorWriter struct {
	cc container.Container
}

func NewErrorCollectorWriter(cc container.Container) *ErrorCollectorWriter {
	return &ErrorCollectorWriter{cc: cc}
}

func (e *ErrorCollectorWriter) Write(le level.Level, module string, message string) error {
	// TODO  Error collector implementation
	return nil
}

func (e *ErrorCollectorWriter) ReOpen() error {
	return nil
}

func (e *ErrorCollectorWriter) Close() error {
	return nil
}
