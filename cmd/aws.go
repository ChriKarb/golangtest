package cmd

import (
	"fmt"
	aws_adapter "github.com/ckone4you/golangtest/pkg/awsIntegration"
	"github.com/spf13/cobra"
)

func NewAWSCommand(profile string, region string) *cobra.Command {

	var awsCmd = &cobra.Command{

		Use:   "awsIntegration",
		Short: "List S3 buckets",
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := aws_adapter.ListS3Buckets(profile, region)
			fmt.Printf("", result)
			return err
		},
	}
	awsCmd.Flags().StringP("awsIntegration-profile", "p", "default", "AWS profile to use")
	awsCmd.Flags().StringP("awsIntegration-region", "r", "eu-central-1", "AWS region to use")
	return awsCmd
}
