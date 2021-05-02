package main

import (
	"fmt"
	"github.com/ku/cryptact-etherscan/csv"
	etherscan "github.com/nanmu42/etherscan-api"
	"github.com/spf13/cobra"
	"os"
)

var etherCmd = &cobra.Command{
	Use:   "ether",
	Short: "scan the transactions by etherscan",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := etherscan.New(etherscan.Mainnet, opts.apiKey)
		_ = client
		address := etherOpts.address
		start := 0
		end := 0xffffffff
		page := 1
		offset := 0
		txs, err := client.NormalTxByAddress(address, &start, &end, page, offset, true)
		if err != nil {
			return err
		}

		c := csv.New(os.Stdout)
		for _, tx := range txs {
			c.Add(NewCryptactLog(etherOpts.source, &tx))
		}
		c.Flush()
		return nil
	},
}

type RootOptions struct {
	apiKey string
}
type EtherOptions struct {
	source  string
	address string
}

var client *etherscan.Client
var opts = &RootOptions{}
var etherOpts = &EtherOptions{}

func main() {
	var rootCmd = &cobra.Command{Use: "cryptac"}

	rootCmd.AddCommand(etherCmd)
	rootCmd.PersistentFlags().StringVar(&opts.apiKey, "key", "", "etherscan api key")
	rootCmd.MarkFlagRequired("key")
	rootCmd.MarkPersistentFlagRequired("key")

	etherCmd.PersistentFlags().StringVar(&etherOpts.source, "source", "", "cryptact source name")
	etherCmd.MarkPersistentFlagRequired("source")

	etherCmd.PersistentFlags().StringVar(&etherOpts.address, "address", "", "address to scan")
	etherCmd.MarkPersistentFlagRequired("address")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
