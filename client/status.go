package main

import (
    "fmt"
	"os"
	"io/ioutil"
    "github.com/jawher/mow.cli"
	"net/http"
)

func main() {
    st := cli.App("st", "get server's status")

   
	
	
	st.Command("status", "get status from 127.0.0.1:8000/status", func(cmd *cli.Cmd) {
		
		cmd.Action = func() {
			
			res, err := http.Get("http://127.0.0.1:8000/status")
			if err != nil {
				fmt.Println(err.Error())
			}
			defer res.Body.Close()

			stat, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err.Error())
			}

			fmt.Printf("%s", stat)
		}
	})
	st.Run(os.Args)
				
				
}


	