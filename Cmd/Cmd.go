package Cmd

import (
	"database/sql"
	"github.com/WenkanHuang/gin_gorm/Config"
	"github.com/WenkanHuang/gin_gorm/Db"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"net/http"
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config/dev.yaml", "config file (default is $HOME/.cobra.yaml)")
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

		Db.DB.AutoMigrate(&model.User{})
		d, _ := Db.DB.DB()
		defer func(d *sql.DB) {
			err := d.Close()
			if err != nil {
				log.Println(err.Error())
			}
		}(d)

		r := router.SetupRouter()
		r.Run()

		port := viper.GetString("port")
		log.Println("port = *** =", port)
		return http.ListenAndServe(port, nil) // listen and serve
	}

	return rootCmd.Execute()

}
