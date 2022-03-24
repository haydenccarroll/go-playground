/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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

// pemdasCmd represents the pemdas command
var pemdasCmd = &cobra.Command{
	Use:   "pemdas \"expression string\"",
	Short: "This is a PEMDAS calculator created in Go.",
	Long: `To use, do pemdas "3+4*(-12+5). This returns the result of the expression.

There is an optional flag -i or --implicit-mul, which turns on implicit multiplication. 
Implicit multiplication is off by default.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(RecursivePemdas(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(pemdasCmd)
	pemdasCmd.Flags().BoolP("implicit-mul", "i", false, "True if implicit multiplication is above multiplication")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pemdasCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pemdasCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
