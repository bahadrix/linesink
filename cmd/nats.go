package cmd

import (
	nats "github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// natsCmd represents the nats command
var natsCmd = &cobra.Command{
	Use:   "nats",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// natsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// natsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
