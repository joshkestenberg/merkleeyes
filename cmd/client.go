package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/types"
)

var txCmd = &cobra.Command{
	Use:   "tx",
	Short: "Send a transaction using subcommands get, set, cas, remove",
}

type TxStruct struct {
	Method       string `json:"method"`
	Key          string `json:"key"`
	Value        string `json:"value"`
	CompareValue string `json:"compare_value"`
	Nonce        string `json:"nonce"`
}

func sendTx(txStruct TxStruct) {
	jsonTx, err := json.Marshal(txStruct)
	if err != nil {
		panic(err)
	}

	c := client.NewHTTP("tcp://localhost:46657", "/websocket")
	tx := types.Tx(jsonTx)
	result, err := c.BroadcastTxCommit(tx)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get data by key",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			panic("get expects 2 args: key, nonce")
		}

		txStruct := TxStruct{Method: "get", Key: args[0], Nonce: args[1]}
		sendTx(txStruct)
	},
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set new data or revise existing data",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			panic("set expects 3 args: key, value, nonce")
		}

		txStruct := TxStruct{Method: "set", Key: args[0], Value: args[1], Nonce: args[2]}
		sendTx(txStruct)
	},
}

var casCmd = &cobra.Command{
	Use:   "cas",
	Short: "Overwrite existing data with confirmation",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 4 {
			panic("cas expects 4 args: key, value, compare_value, nonce")
		}

		txStruct := TxStruct{Method: "cas", Key: args[0], Value: args[1], CompareValue: args[2], Nonce: args[3]}
		sendTx(txStruct)
	},
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove data by key",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			panic("remove expects 2 args: key, nonce")
		}

		txStruct := TxStruct{Method: "remove", Key: args[0], Nonce: args[1]}
		txStruct.Nonce = args[1]

		sendTx(txStruct)
	},
}

func init() {
	RootCmd.AddCommand(txCmd)
	txCmd.AddCommand(getCmd)
	txCmd.AddCommand(setCmd)
	txCmd.AddCommand(casCmd)
	txCmd.AddCommand(removeCmd)
}
