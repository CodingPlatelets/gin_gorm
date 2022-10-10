package cmd

import (
	"database/sql"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/WenkanHuang/gin_gorm/config"
	"github.com/WenkanHuang/gin_gorm/db"
	"github.com/WenkanHuang/gin_gorm/model"
	"github.com/WenkanHuang/gin_gorm/router"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{}
)

func initConfig() {
	config.MustInit(os.Stdout, cfgFile) // 配置初始化
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config/application.yaml", "config file (default is $HOME/.cobra.yaml) ")
	rootCmd.PersistentFlags().Bool("debug", true, "start debug")
	viper.SetDefault("gin.mode", rootCmd.PersistentFlags().Lookup("debug"))
}

func Execute() error {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat:           "2006-01-02 15:04:05",
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		ForceColors:               true,
	})

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	//log.SetLevel(log.WarnLevel)

	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		_, err := db.Mysql(
			viper.GetString("db.hostname"),
			viper.GetInt("db.port"),
			viper.GetString("db.username"),
			viper.GetString("db.password"),
			viper.GetString("db.dbname"),
		)
		if err != nil {
			return err
		}

		errInMigration := db.DB.AutoMigrate(&model.User{}, &model.Group{}, &model.Todo{})
		if errInMigration != nil {
			log.Infof(errInMigration.Error())
			return errInMigration
		}
		d, _ := db.DB.DB()
		defer func(d *sql.DB) {
			err := d.Close()
			if err != nil {
				log.Infof(err.Error())
			}
		}(d)

		r := router.SetupRouter()
		port := viper.GetString("server.port")
		errRun := r.Run(port)
		if errRun != nil {
			return errRun
		}
		log.Println("port = *** =", port)
		return nil
	}

	return rootCmd.Execute()

}
