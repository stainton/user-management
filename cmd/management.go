package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func NewUserManager() *cobra.Command {
	var connectString string
	userManager := &cobra.Command{
		Use:   "user-manager",
		Short: "apiserver of user manager.",
		Long:  `apiserver of user manager.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			router := gin.Default()

			router.Run(connectString)
		},
	}
	userManager.Flags().StringVarP(&connectString, "connectString", "c", ":8080", "ip and port user-manager listen on")
	return userManager
}
