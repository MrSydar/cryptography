/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// trivialCmd represents the trivial command
var trivialCmd = &cobra.Command{
	Use:   "trivial",
	Short: "Run trivial cryptography method",
	Long:  `Run trivial cryptography method. You can create a secret from parts or vice versa.`,
	Run: func(cmd *cobra.Command, args []string) {
		method, _ := cmd.Flags().GetString("method")
		submethod, _ := cmd.Flags().GetString("sub-method")
		if method == "trivial" {
			if submethod == "cypher" {
				fmt.Println("Trivial method cypher")
			} else if submethod == "decypher" {
				fmt.Println("Trivial method decypher")
			} else {
				fmt.Println("No such sub method!")
			}
		} else if method == "shamir" {
			if submethod == "cypher" {
				fmt.Println("Shamir method cypher")
			} else if submethod == "decypher" {
				fmt.Println("Shamir method decypher")
			} else {
				fmt.Println("No such sub method!")
			}
		} else {
			fmt.Println("No such method!")
		}
	},
}

func init() {
	rootCmd.AddCommand(trivialCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//trivialCmd.PersistentFlags().String("method", "", "A help for foo")
	trivialCmd.PersistentFlags().StringP("secret", "s", "", "A secret to cypher")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	trivialCmd.Flags().String("method", "", "A method to run (trivial, shamir)")
	trivialCmd.MarkFlagRequired("method")

	trivialCmd.Flags().String("sub-method", "", "A sub method to run (cypher, decypher)")
	trivialCmd.MarkFlagRequired("sub-method")
}
