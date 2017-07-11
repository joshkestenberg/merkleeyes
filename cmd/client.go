package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/types"
)

var c = client.NewHTTP("tcp://localhost:44657", "/websocket")

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get data by key",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("get expects 2 args: key, nonce")
		}

		type get struct {
			method string
			key    string
			nonce  string
		}

		getStruct := get{method: "get", key: args[0], nonce: args[1]}

		jsonTx, err := json.Marshal(getStruct)
		if err != nil {
			panic(err)
		}
		tx := types.Tx(jsonTx)
		result, err := c.BroadcastTxCommit(tx)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	},
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set new data or revise existing data",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			panic("set expects 3 args: key, value, nonce")
		}

		type set struct {
			method string
			key    string
			value  string
			nonce  string
		}

		setStruct := set{method: "set", key: args[0], value: args[1], nonce: args[2]}

		jsonTx, err := json.Marshal(setStruct)
		if err != nil {
			panic(err)
		}
		tx := types.Tx(jsonTx)
		result, err := c.BroadcastTxCommit(tx)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	},
}

var casCmd = &cobra.Command{
	Use:   "cas",
	Short: "Overwrite existing data with confirmation",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			panic("cas expects 4 args: key, value, comparevalue, nonce")
		}

		type cas struct {
			method       string
			key          string
			value        string
			comparevalue string
			nonce        string
		}

		casStruct := cas{method: "cas", key: args[0], value: args[1], comparevalue: args[2], nonce: args[3]}

		jsonTx, err := json.Marshal(casStruct)
		if err != nil {
			panic(err)
		}
		tx := types.Tx(jsonTx)
		result, err := c.BroadcastTxCommit(tx)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	},
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove data by key",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("remove expects 2 args: key, nonce")
		}

		type remove struct {
			method string
			key    string
			nonce  string
		}
		removeStruct := remove{method: "remove", key: args[0], nonce: args[1]}

		jsonTx, err := json.Marshal(removeStruct)
		if err != nil {
			panic(err)
		}
		tx := types.Tx(jsonTx)
		result, err := c.BroadcastTxCommit(tx)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
	RootCmd.AddCommand(setCmd)
	RootCmd.AddCommand(casCmd)
	RootCmd.AddCommand(removeCmd)
}
