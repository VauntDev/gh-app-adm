package cmd

import "github.com/spf13/cobra"

func jwt() *cobra.Command {
	jwtCmd := &cobra.Command{
		Use: "jwt",
	}

	return jwtCmd
}

func generate() *cobra.Command {
	gen := &cobra.Command{
		Use: "gen",
	}

	return gen
}
