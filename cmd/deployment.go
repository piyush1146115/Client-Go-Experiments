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
	"github.com/piyush1146115/Client-Go-Experiments/api"
	"github.com/spf13/cobra"
)

var deploymentName string
var replicas int32

// deploymentCmd represents the deployment command
var deploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deployment called")
	},
}

var createDeployment = &cobra.Command{
	Use:   "createDeployment",
	Short: "This command is for creating deployment",
	Long:  "This command is used for creating deployment object using kubernetes API",
	Run: func(cmd *cobra.Command, args []string) {
		api.CreateDeployment()
	},
}

var getDeployment = &cobra.Command{
	Use:   "getDeployment",
	Short: "This command will list all deployment resources running",
	Long:  "This command will list all deployment resources that are running on",
	Run: func(cmd *cobra.Command, args []string) {
		api.GetDeployment()
	},
}

var deleteDeployment = &cobra.Command{
	Use:   "deleteDeployment",
	Short: "This command will delete a deployment",
	Long:  "This command will delete a deployment resource followed by a flag containing the name of the deployment to be deleted",
	Run: func(cmd *cobra.Command, args []string) {
		api.SetDeploymentName(deploymentName)
		api.DeleteDeployment()
	},
}

var updateDeployment = &cobra.Command{
	Use:   "updateDeployment",
	Short: "This command will update a deployment",
	Long:  "This command will update the number of replicas of a deployment object",
	Run: func(cmd *cobra.Command, args []string) {
		api.SetDeploymentName(deploymentName)
		api.SetReplicas(replicas)
		api.UpdateDeployment()
	},
}

func init() {
	rootCmd.AddCommand(deploymentCmd)
	rootCmd.AddCommand(createDeployment)
	rootCmd.AddCommand(getDeployment)
	rootCmd.AddCommand(deleteDeployment)
	rootCmd.AddCommand(updateDeployment)
	deleteDeployment.PersistentFlags().StringVarP(&deploymentName, "name", "n", "go-client-api-server", "This flag sets the name of the deployment to be deleted")
	updateDeployment.PersistentFlags().Int32VarP(&replicas, "replica", "r", 1, "This flag sets the number of replica in the update operation")
	updateDeployment.PersistentFlags().StringVarP(&deploymentName, "name", "n", "go-client-api-server", "This flag sets the name of the deployment to be updated")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deploymentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deploymentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
