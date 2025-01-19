package cmd

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/stainton/user-management/internal/db"
	"github.com/stainton/user-management/internal/handler"
)

func NewUserManager() *cobra.Command {
	var connectString string
	var dbString string
	var jwtKey string
	userManager := &cobra.Command{
		Use:   "user-manager",
		Short: "apiserver of user manager.",
		Long:  `apiserver of user manager.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			d := db.NewDB(dbString)
			if d == nil {
				os.Exit(1)
			}
			router := gin.Default()
			handler.RegisterUserManager(d, router)
			router.Run(connectString)
		},
	}
	userManager.Flags().StringVarP(&dbString, "dbString", "d", ":3306", "数据库连接地址")
	userManager.Flags().StringVarP(&connectString, "connectString", "c", ":8080", "服务端连接地址")
	userManager.Flags().StringVarP(&jwtKey, "jwtKey", "j", "", "jwt加密字段")
	userManager.MarkFlagRequired("jwtKey")
	return userManager
}
