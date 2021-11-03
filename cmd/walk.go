package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var walkCMD = &cobra.Command{
	Use:   "walk",
	Short: "Walk through folders and list all filetypes",
	Run:   walk,
}

type FolderStats struct {
	Folder         string         `json:"folder"`
	FileCount      int            `json:"file_count"`
	FileTypes      map[string]int `json:"file_types"`
	AdditionalInfo string         `json:"additional_info"`
}

//FileTypeCounts map[string]int `json:"file_type_counts"`

type FileInfo struct {
	Extension string `json:"extension"`
	Size      string `json:"size"`
	//TODO: more info ...
}

var path string

func init() {
	RootCmd.AddCommand(walkCMD)
	f := walkCMD.Flags()
	f.StringVar(&path, "path", ".", "Default.: C:\"")
}

func walk(cmd *cobra.Command, args []string) {
	fss, err := walkWolder(path)
	if err != nil {
		fmt.Println("error: ", err)
	}

	file, _ := json.MarshalIndent(fss, "", " ")
	var hostname = "no_hostname"
	hostname, _ = os.Hostname()
	_ = ioutil.WriteFile(hostname+".json", file, 0644)
}

func walkWolder(walkPath string) ([]FolderStats, error) {
	var fss []FolderStats

	err := filepath.Walk(walkPath,
		func(path string, info os.FileInfo, err error) error {
			fs := FolderStats{}
			if err != nil {
				fs.Folder = path //error open folder
				fs.AdditionalInfo = err.Error()
			}
			if info.IsDir() {
				if path != walkPath {
					fmt.Println("Walk folder: ", path)
					files, err := getExt(path)
					if err != nil {
						return err
					}
					fs.Folder = path
					fs.FileCount = len(files)
					fs.FileTypes = countUniques(files)
					fss = append(fss, fs)
				}
			}
			return nil
		})
	if err != nil {
		return nil, err
	}

	return fss, nil
}

func getExt(path string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, filepath.Ext(info.Name()))
		}
		return nil
	})
	return files, err
}

func getFiles(path string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, info.Name())
		}
		return nil
	})
	return files, err
}

func countUniques(arr []string) map[string]int {
	m := make(map[string]int)
	for _, i := range arr {
		m[i] = m[i] + 1
	}

	return m
}
