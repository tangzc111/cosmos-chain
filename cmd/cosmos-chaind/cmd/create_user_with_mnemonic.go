package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"

	"cosmos-chain/x/core/types"
)

func newCreateUserWithMnemonicCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-user-with-mnemonic [username]",
		Short: "Generate a fresh wallet, create the user on-chain, and print the mnemonic/address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString("description")
			if err != nil {
				return err
			}

			mnemonic, addr, err := generateTempMnemonic(clientCtx)
			if err != nil {
				return err
			}

			msg := &types.MsgCreateUser{
				Creator:     clientCtx.GetFromAddress().String(),
				Index:       addr.String(),
				Address:     addr.String(),
				Username:    args[0],
				Description: description,
			}

			if err := tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg); err != nil {
				return err
			}

			_, err = fmt.Fprintf(cmd.OutOrStdout(), "Mnemonic: %s\nAddress: %s\n", mnemonic, addr.String())
			return err
		},
	}

	cmd.Flags().String("description", "", "Optional description to attach to the on-chain user record")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func generateTempMnemonic(clientCtx client.Context) (string, cosmostypes.AccAddress, error) {
	tmpKeyring := keyring.NewInMemory(clientCtx.Codec)
	uid := fmt.Sprintf("tmp-user-%d", time.Now().UnixNano())
	record, mnemonic, err := tmpKeyring.NewMnemonic(uid, keyring.English, cosmostypes.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	if err != nil {
		return "", nil, err
	}

	addr, err := record.GetAddress()
	if err != nil {
		return "", nil, err
	}

	return mnemonic, cosmostypes.AccAddress(addr), nil
}
