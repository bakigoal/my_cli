/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "This command will get the desired Gopher",
	Long:  `This get command will call GitHub respository in order to return the desired Gopher.`,
	Run:   getCmdHandler,
}

func getCmdHandler(_ *cobra.Command, args []string) {
	var gopherName = "dr-who"

	if len(args) >= 1 && args[0] != "" {
		gopherName = args[0]
	}

	// Get the data
	URL := "https://github.com/scraly/gophers/raw/main/" + gopherName + ".png"
	fmt.Println("Try to get '" + gopherName + "' Gopher...")
	response, err := http.Get(URL)
	logErrorIfExists(err)
	defer logErrorIfExists(response.Body.Close())

	if response.StatusCode == 200 {
		saveImage(gopherName, response.Body)
	} else {
		fmt.Println("Error: " + gopherName + " not exists! :-(")
	}
}

func saveImage(fileName string, body io.ReadCloser) {
	// Create the file
	mkDirIfNotExist("img")
	out, err := os.Create("img/" + fileName + ".png")
	logErrorIfExists(err)
	defer logErrorIfExists(out.Close())

	// Writer the body to file
	_, err = io.Copy(out, body)
	logErrorIfExists(err)

	fmt.Println("Perfect! Just saved in " + out.Name() + "!")
}

func mkDirIfNotExist(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		logErrorIfExists(err)
	}
}

func logErrorIfExists(err error) {
	if err != nil {
		log.Println(err)
	}
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
