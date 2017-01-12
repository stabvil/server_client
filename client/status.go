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
	st := cli.App("lb", "client' app")

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
				defer res.Body.Close()
			}
		}
	})

	st.Command("save", "save in file --name --data[file.json]", func(cmd *cli.Cmd) {

		name_file := cmd.StringOpt("file", "", "Where ll save json data")
		name_json_file := cmd.StringOpt("json", "", "name of file where jsondata")

		cmd.Action = func() {
			//check
			if *name_file == "" {
				fmt.Println("Null file name")
				return
			}
			if *name_json_file == "" {
				fmt.Println("Null file name jsondata")
				return
			}

			//Open json file
			file_json, err := os.Open(*name_json_file)
			if err != nil {
				fmt.Println(err.Error())
			}
			defer file_json.Close()

			//read from json file
			json_str, err := ioutil.ReadAll(file_json)
			if err != nil {
				fmt.Println(err.Error())
			}

			err1 := ioutil.WriteFile(*name_file, json_str, 0666)
			if err1 != nil {
				fmt.Println(err.Error())
			}
		}
	})

	st.Command("get", "get data from file", func(cmd *cli.Cmd) {
		file := cmd.StringOpt("file", "", "name of file where data")

		cmd.Action = func() {
			if *file == "" {
				fmt.Println("Null file name")
				return
			}
			// open file
			read_file, err := os.Open(*file)
			if err != nil {
				fmt.Println(err.Error())
			}
			defer read_file.Close()

			str, err := ioutil.ReadAll(read_file)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(string(str))
		}

	})

	st.Run(os.Args)

}
