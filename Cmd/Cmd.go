package Cmd

import (
	"database/sql"
	"github.com/WenkanHuang/gin_gorm/Config"
	"github.com/WenkanHuang/gin_gorm/Db"
	"github.com/WenkanHuang/gin_gorm/Model"
	"github.com/WenkanHuang/gin_gorm/router"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	cfgFile string
	logger  = &logrus.Logger{}
	rootCmd = &cobra.Command{}
)

func initConfig() {
	Config.MustInit(os.Stdout, cfgFile) // 配置初始化
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config/application.yaml", "config file (default is $HOME/.cobra.yaml) ")
	rootCmd.PersistentFlags().Bool("debug", true, "开启debug")
	viper.SetDefault("gin.mode", rootCmd.PersistentFlags().Lookup("debug"))
}

func Execute() error {
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		_, err := Db.Mysql(
			viper.GetString("db.hostname"),
			viper.GetInt("db.port"),
			viper.GetString("db.username"),
			viper.GetString("db.password"),
			viper.GetString("db.dbname"),
		)
		if err != nil {
			return err
		}

		errInMigration := Db.DB.AutoMigrate(&Model.User{}, &Model.Group{}, &Model.Todo{})
		if errInMigration != nil {
			logger.Println(errInMigration.Error())
			return errInMigration
		}
		d, _ := Db.DB.DB()
		defer func(d *sql.DB) {
			err := d.Close()
			if err != nil {
				logger.Println(err.Error())
			}
		}(d)

		r := router.SetupRouter()
		port := viper.GetString("server.port")
		errRun := r.Run(port)
		if errRun != nil {
			return errRun
		}
		logger.Println("port = *** =", port)
		return nil
	}

	return rootCmd.Execute()

}
