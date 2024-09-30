package cmd

import (
	"duckVault/internal"
	"encoding/base64"
	"github.com/spf13/cobra"
	"strings"
)

var filePath, namespace string

func NewCmdGenerate() *cobra.Command {
	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate files based on a given file like env..",
		Run: func(cmd *cobra.Command, args []string) {

			if !strings.Contains(filePath, ".env") {
				return
			}
			scannedValues := internal.ScanENVFile(filePath)
			for _, value := range scannedValues {
				value.TypeValue = base64.StdEncoding.EncodeToString([]byte(value.TypeValue))
				secrets := make(map[string]string)
				secrets[value.TypeKey] = value.TypeValue
				internal.WriteSecretYAMLFile(secrets, value.TypeKey, namespace, "Opaque")
			}
		},
	}

	generateCmd.Flags().StringVar(&filePath, "f", ".env", "filePath")
	generateCmd.Flags().StringVar(&namespace, "ns", "default", "Namespace")

	return generateCmd
}
