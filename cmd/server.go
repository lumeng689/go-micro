package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"github.com/lumeng689/go-micro/server"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "world server api",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Recover error : %v", err)
			}
		}()
		server.Run()
	},
}

func init() {
	serverCmd.Flags().StringVarP(&server.ServerPort, "port", "p", "50051", "server port")
	serverCmd.Flags().BoolVar(&server.SecureServer, "secure", false, "use secure to protect server")
	serverCmd.Flags().StringVarP(&server.CertPemPath, "cert-pem", "", "./conf/certs/server.pem", "cert-pem path")
	serverCmd.Flags().StringVarP(&server.CertKeyPath, "cert-key", "", "./conf/certs/server-key.pem", "cert-key path")
	serverCmd.Flags().StringVarP(&server.CertServerName, "cert-server-name", "", "127.0.0.1", "server's hostname")
	serverCmd.Flags().StringVarP(&server.SwaggerDir, "swagger-dir", "", "proto", "path to the directory which contains swagger definitions")

	// 添加到指令队列中
	rootCmd.AddCommand(serverCmd)
}
