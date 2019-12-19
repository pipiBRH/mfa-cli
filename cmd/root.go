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
	Use:   "mfa",
	Short: "Easy way to obtain your mfa digi without mobile device.",
	Run: func(cmd *cobra.Command, args []string) {
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
	rootCmd.Flags().StringVarP(&key, "key", "k", "", "Specify the secret key which provided by your platform.")
}
