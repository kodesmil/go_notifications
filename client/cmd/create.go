package cmd

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"

	pb "notifications/proto"
	"github.com/spf13/cobra"
	"time"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new notification",
	Long: `Create a new notification on the server through gRPC. 
	
	A blog post requires an AuthorId, Title and Content.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		author, err := cmd.Flags().GetString("user")
		title, err := cmd.Flags().GetString("title")
		content, err := cmd.Flags().GetString("content")
		timestampString, err := cmd.Flags().GetString("time")
		// convert string to Time.time object
		timestampGo, err := time.Parse(time.RFC3339, timestampString)
		// convert Time.time object to Proto format
		timestampProto, err := ptypes.TimestampProto(timestampGo)
		fmt.Print(timestampProto)
		if err != nil {
			return err
		}
		notification := &pb.Notification{
			UserId:  author,
			Title:   title,
			Content: content,
			Time:    timestampProto,
		}

		token := &pb.IDtoken{
			Token: "eyJhbGciOiJSUzI1NiIsImtpZCI6IjgzYTczOGUyMWI5MWNlMjRmNDM0ODBmZTZmZWU0MjU4Yzg0ZGI0YzUiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vYWJseS1oZWFsdGgiLCJhdWQiOiJhYmx5LWhlYWx0aCIsImF1dGhfdGltZSI6MTU4NTg2Mzg2OSwidXNlcl9pZCI6Imt2dFc4bDY3VnVNeE40VVdaVjhWZ3JseThIbTIiLCJzdWIiOiJrdnRXOGw2N1Z1TXhONFVXWlY4VmdybHk4SG0yIiwiaWF0IjoxNTg1ODYzODY5LCJleHAiOjE1ODU4Njc0NjksImVtYWlsIjoiaGVsbG9Aa29kZXNtaWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJmaXJlYmFzZSI6eyJpZGVudGl0aWVzIjp7ImVtYWlsIjpbImhlbGxvQGtvZGVzbWlsLmNvbSJdfSwic2lnbl9pbl9wcm92aWRlciI6InBhc3N3b3JkIn19.Bd1vmLWpZnd93ee4YLU_gkr4AUhvEE96IWXEThVxPuop1PmB1bt6iYj9OlxWNn2K_QIIWRfIJUNQ9Vc50hTkPBw3SdxwmzdOkQBLv0IYudRmop7N4hsLDBJZgLaILX-Rli52grOAi7vV91Ytz-eEIbN0Iq4uzbJg4uJKnGWxrEepFe8TG5aNcNu9YDiKYS87SXRrZGSMi4JpXBBjr9FNTd35uZxKVcuNTs5j6W7-aehqQMUcpDQ_nsIEp7Iap49Q4ZJ_eUIhu4HbnkFkJ2ieCeUOZ8AqKL9dPwZj644HXODQSMHVtYc5DhibxP7PDm3NWrYjvNRqK5wYe6PLDwUHEw",
		}

		response, err := client.NotificationCreate(
			context.TODO(),
			&pb.NotificationCreateRequest{
				Idtoken: token,
				Notification: notification,
			},
		)
		if err != nil {
			return err
		}
		fmt.Printf("Notification created: %s\n", response.Notification.Id)
		return nil
	},
}

func init() {
	createCmd.Flags().StringP("user", "u", "", "Add an user")
	createCmd.Flags().StringP("title", "t", "", "A title for the notification")
	createCmd.Flags().StringP("content", "c", "", "The content for the notification")
	createCmd.Flags().StringP("time", "z", "", "The time to notify")
	createCmd.MarkFlagRequired("user")
	createCmd.MarkFlagRequired("title")
	createCmd.MarkFlagRequired("content")
	createCmd.MarkFlagRequired("time")
	rootCmd.AddCommand(createCmd)
}
