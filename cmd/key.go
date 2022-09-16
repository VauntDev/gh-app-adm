package cmd

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func keyCmd() *cobra.Command {
	keyCmd := &cobra.Command{
		Use: "key",
	}
	keyCmd.AddCommand(bytes())
	return keyCmd
}

func bytes() *cobra.Command {
	bytesCmd := &cobra.Command{
		Use:   "base64",
		Short: "convert a private key file to base64 bytes",
		RunE: func(cmd *cobra.Command, args []string) error {

			if _, err := os.Stat(viper.GetString(privateKeyPath)); errors.Is(err, os.ErrNotExist) {
				return err
			}

			key, err := ioutil.ReadFile(viper.GetString(privateKeyPath))
			if err != nil {
				return err
			}

			fmt.Printf("Key bytes: %s", base64.StdEncoding.EncodeToString((key)))

			return nil
		},
	}
	bytesCmd.Flags().String(privateKeyPath, "", "--private-key-path, location of the private key file")
	return bytesCmd
}
