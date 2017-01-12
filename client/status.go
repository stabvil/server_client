package main

import (
	"encoding/json"
	"fmt"
	"github.com/jawher/mow.cli"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	st := cli.App("st", "get server's status")

	st.Command("status", "get status from 127.0.0.1:8000/status", func(cmd *cli.Cmd) {

		cmd.Action = func() {

			res, err := http.Get("http://127.0.0.1:8000/status")
			if err != nil {
				status_fail, _ := json.Marshal([]string{"server", "not running"})
				fmt.Println(string(status_fail))
			} else {
				stat, err := ioutil.ReadAll(res.Body)
				if err != nil {
					fmt.Println(err.Error())
				}
				fmt.Println(string(stat))
			}

		}
	})
	st.Run(os.Args)

}
