package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pquerna/otp/totp"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "mfa",
	Long: `Easy way to obtain your mfa digi without mobile device.
You can set key by environment variable MFA_SECRET_KEY or cli flags`,
	Run: func(cmd *cobra.Command, args []string) {
		if key == "" {
			log.Fatal("Invalid key")
		}

		digi, err := totp.GenerateCode(key, time.Now())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(digi)
	},
}

var (
	key string
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&key, "key", "k", os.Getenv("MFA_SECRET_KEY"), "Specify the secret key which provided by your platform.")
}
