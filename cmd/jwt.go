package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func jwtCmd() *cobra.Command {
	jwtCmd := &cobra.Command{
		Use: "jwt",
	}
	jwtCmd.AddCommand(generate())
	return jwtCmd
}

func generate() *cobra.Command {
	gen := &cobra.Command{
		Use:   "gen",
		Short: "generate a jwt token",
		RunE: func(cmd *cobra.Command, args []string) error {

			if _, err := os.Stat(viper.GetString(privateKeyPath)); errors.Is(err, os.ErrNotExist) {
				return err
			}

			key, err := ioutil.ReadFile(viper.GetString(privateKeyPath))
			if err != nil {
				return err
			}

			rsaKey, err := jwt.ParseRSAPrivateKeyFromPEM(key)
			if err != nil {
				return err
			}

			claims := &jwt.StandardClaims{
				IssuedAt:  time.Now().Unix() - 60,
				ExpiresAt: time.Now().Unix() + (10 * 60),
				Issuer:    viper.GetString(appId),
			}

			token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
			ss, err := token.SignedString(rsaKey)
			if err != nil {
				return err
			}

			fmt.Println("Signed Token:", ss)

			return nil
		},
	}
	gen.Flags().String(privateKeyPath, "", "--private-key-path, location of the private key file")
	gen.Flags().String(appId, "", "--app-id, app id of github app")
	return gen
}
