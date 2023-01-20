package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "trace the ip",
	Long: `trace the ip address`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				fmt.Println(ip)
				showData()
			}
		} else {
			fmt.Println("Please provide an IP address to trace")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}

type Ip struct {
	IP			string `json::"ip"`
	City 		string `json::"city"`
	Region 		string `json::"region"`
	Country 	string `json::"country"`
	Loc 		string `json::"loc"`
	Timezone 	string `json::"timezone"`
	Postal 		string `json::"postal"`
}

func showData(){
	url := "http://ipinfo.io/1.1.1.1/geo"
	responseByte := getData(url)

	data := Ip{}

	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		log.Println("Unable to unmarshal the response")
	}

	fmt.Println("DATA FOUND :")
	fmt.Println(data)
}

func getData(url string) []byte{
	response, err := http.Get(url)

	if err != nil{
		log.Println("Unable to get response")
	}

	responseByte, err := ioutil.ReadAll(response.Body)

	if err != nil{
		log.Println("Unable to read the response")
	}

	return responseByte
}
