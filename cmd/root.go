package cmd

import (
	"TaskRESTfulExercise/services/httpserver"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile    string
	goversion  string
	buildstamp string
	githash    string
	port       int
)

var rootCmd = &cobra.Command{
	Use:   "Start",
	Short: "RESTful API'",
	Long:  `A RESTful API server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start a RESTful API server by 'http' or 'help' to get more infomation")
	},
}

var serverCmd = &cobra.Command{
	Use:   "http",
	Short: "start http server",
	Long:  `start http server with port`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("http server : %s\n", viper.GetString("app.name"))
		httpserver.StartGinServer(port)
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version",
	Long:  `show version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Git Commit Hash: %s\n", githash)
		fmt.Printf("UTC Build Time : %s\n", buildstamp)
		fmt.Printf("Golang Version : %s\n", goversion)
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.toml", "config file (default is $PWD/config.toml)")
	serverCmd.Flags().IntVar(&port, "port", 8080, "default is 8080")
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(versionCmd)
}

func initConfig() {

	viper.SetConfigFile(cfgFile)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
