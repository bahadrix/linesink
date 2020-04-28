package cmd

import (
	nats "github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// natsCmd represents the nats command
var natsCmd = &cobra.Command{
	Use:   "nats",
	Short: "NATS sink",
	Long:  `Publish lines to NATS endpoint`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		subject, _ := cmd.Flags().GetString("subject")

		nc, err := nats.Connect(url)
		if err != nil {
			log.Fatal(err)
		}

		lines, err := readPipedLines(10)

		if err != nil {
			log.Fatal(err)
		}

		for line := range lines {
			err = nc.Publish(subject, line)
			if err != nil {
				log.Errorf("publish failed: %s", err.Error())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(natsCmd)

	natsCmd.Flags().String("url", "", "NATS Url")
	natsCmd.Flags().StringP("subject", "s", "", "Subject to write data into")

	_ = natsCmd.MarkFlagRequired("url")
	_ = natsCmd.MarkFlagRequired("subject")
}
