package cmd

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var path string

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan the routes",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if path != "" {

			fmt.Printf("Testing connection with %s\n", path)
			timeout := time.Duration(1 * time.Second)
			_, err := net.DialTimeout("tcp", path, timeout)
			if err != nil {
				fmt.Printf(" --> %s%s\n", path, ", Error.")
				log.Print(err)
			} else {
				fmt.Printf(" --> %s%s\n", path, ", Correct. ")
			}

		} else {

			readFile, err := os.Open("/home/gmo/Escritorio/uchile/cobra/ejemplo-cobra/endpoints.txt")
			if err != nil {
				log.Fatal(err)
			}
			fileScanner := bufio.NewScanner(readFile)
			fileScanner.Split(bufio.ScanLines)

			var lines []string
			for fileScanner.Scan() {
				lines = append(lines, fileScanner.Text())
			}

			readFile.Close()
			for _, line := range lines {

				fmt.Printf("Testing connection with %s\n", line)
				timeout := time.Duration(1 * time.Second)
				_, err := net.DialTimeout("tcp", line, timeout)
				if err != nil {
					fmt.Printf(" --> %s%s\n", line, ", Error.")
					log.Print(err)
				} else {
					fmt.Printf(" --> %s%s\n", line, ", Correct. ")
				}

			}
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
	//scanCmd.Flags().BoolP("scan", "s", false, "scan the routes")
	scanCmd.PersistentFlags().StringVar(&path, "parameter", "", "path to scan")
}
