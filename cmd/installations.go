package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func installations() *cobra.Command {
	installations := &cobra.Command{
		Use:     "installations",
		Aliases: []string{"inst", "i"},
		Short:   "lists the installations belonging to your github application",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO replace with golang github client/app packages
			client := &http.Client{}
			req, err := http.NewRequest("GET", "https://api.github.com/app/installations", nil)
			if err != nil {
				return err
			}
			req.Header.Add("Authorization", "Bearer "+viper.GetString(signedJwt))
			req.Header.Add("Accept", "application/vnd.github+json")
			resp, err := client.Do(req)
			if err != nil {
				return err
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil
			}

			var prettyJSON bytes.Buffer
			if err := json.Indent(&prettyJSON, body, "", "    "); err != nil {
				return err
			}

			fmt.Println(prettyJSON.String())
			return nil
		},
	}

	installations.PersistentFlags().String(signedJwt, "", "--jwt=YOUR_JWT, the JWT generated for your applications")
	installations.Flags().Int(perPage, 30, "--per-page=30, the number of results per page (max 100)")
	installations.Flags().Int(page, 1, "--page=1, page number of the results to fetch")

	installations.AddCommand(accessToken())
	installations.AddCommand(repositories())
	return installations
}

func accessToken() *cobra.Command {
	accessToken := &cobra.Command{
		Use:     "accessToken",
		Aliases: []string{"at"},
		Short:   "generate an access token for an installation of a github application",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := &http.Client{}
			req, err := http.NewRequest("POST", "https://api.github.com/app/installations/"+viper.GetString(installId)+"/access_tokens", nil)
			if err != nil {
				return err
			}
			req.Header.Add("Authorization", "Bearer "+viper.GetString(signedJwt))
			req.Header.Add("Accept", "application/vnd.github+json")
			resp, err := client.Do(req)
			if err != nil {
				return err
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil
			}

			var prettyJSON bytes.Buffer
			if err := json.Indent(&prettyJSON, body, "", "    "); err != nil {
				return err
			}
			fmt.Println(prettyJSON.String())

			return nil
		},
	}
	accessToken.Flags().String(installId, "", "--install-id=<INSTALL-ID>, the installation id of the install to generate a token for")

	return accessToken
}

func repositories() *cobra.Command {
	accessToken := &cobra.Command{
		Use:     "repositories",
		Aliases: []string{"repos"},
		Short:   "list the repositories for an installation of a github application",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := &http.Client{}
			req, err := http.NewRequest("GET", "https://api.github.com/installation/repositories", nil)
			if err != nil {
				return err
			}
			req.Header.Add("Authorization", "Bearer "+viper.GetString(token))
			req.Header.Add("Accept", "application/vnd.github+json")
			resp, err := client.Do(req)
			if err != nil {
				return err
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil
			}

			var prettyJSON bytes.Buffer
			if err := json.Indent(&prettyJSON, body, "", "    "); err != nil {
				return err
			}
			fmt.Println(prettyJSON.String())

			return nil
		},
	}
	accessToken.Flags().String(token, "", "--token=<ACCESS-TOKEN>, the Access Token for the installation")

	return accessToken
}
