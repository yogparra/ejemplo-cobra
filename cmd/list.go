package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the paths to scan",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		nombreArchivo := "/home/gmo/Escritorio/uchile/cobra/ejemplo-cobra/endpoints.txt"
		bytesLeidos, err := ioutil.ReadFile(nombreArchivo)
		if err != nil {
			log.Print(err)
		} else {
			contenido := string(bytesLeidos)
			fmt.Println(contenido)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	//listCmd.Flags().BoolP("list", "l", false, "list the paths to scan")
}
