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
	// "fmt"
	"github.com/maxHyeon/goFileDelete/fileDelte"
	//"goFileDelete/fileDelete"
	"github.com/spf13/cobra"
)

// fileDeleteCmd represents the fileDelete command
var fileDeleteCmd = &cobra.Command{
	Use:   "fileDelete [File Absolute Path and Name]",
	Short: "fileDelete [File Absolute Path and Name]",
	Long: `fileDelete [File Absolute Path and Name]`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			panic("File Absolute Path needed")
		}else{
			fileDelete.FileDelete(args[0])
		}
		
	},
}

func init() {
	rootCmd.AddCommand(fileDeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fileDeleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fileDeleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
