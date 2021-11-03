package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var file1 string
var file2 string

var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge two result files to one (json)",
	Run:   merge,
}

func init() {
	RootCmd.AddCommand(mergeCmd)
	f := mergeCmd.Flags()
	f.StringVar(&file1, "file1", ".", "Example: file1.json")
	f.StringVar(&file2, "file2", ".", "Example: file2.json")
	mergeCmd.MarkFlagRequired("file1")
	mergeCmd.MarkFlagRequired("file2")
}

func merge(cmd *cobra.Command, args []string) {
	data1, err := readJsonData(file1)
	if err != nil {
		panic(err)
	}
	data2, err := readJsonData(file2)
	if err != nil {
		panic(err)
	}
	var fsMerge []FolderStats
	for _, d1 := range data1 {
		for _, d2 := range data2 {
			if d1.Folder == d2.Folder {
				fsMerge = append(fsMerge,
					FolderStats{
						Folder:         d1.Folder,
						FileCount:      d1.FileCount + d2.FileCount,
						FileTypes:      mergeJSONMaps(d1.FileTypes, d2.FileTypes),
						AdditionalInfo: strings.Join([]string{d1.AdditionalInfo, d2.AdditionalInfo}, ","),
					})
			}
		}

	}
	//write results to file
	file, _ := json.MarshalIndent(fsMerge, "", " ")
	f1 := strings.Split(file1, ".")
	f2 := strings.Split(file2, ".")
	var filename = strings.Join([]string{f1[0], f2[0], "merge.json"}, "_")
	_ = ioutil.WriteFile(filename, file, 0644)
}

func readJsonData(file string) ([]FolderStats, error) {
	fs := []FolderStats{}

	content, err := os.ReadFile(file)
	if err != nil {
		return fs, err
	}
	if err := json.Unmarshal(content, &fs); err != nil {
		return fs, err
	}

	return fs, nil
}

func mergeJSONMaps(maps ...map[string]int) map[string]int {
	result := make(map[string]int)
	for _, m := range maps {
		for k, v := range m {
			result[k] = result[k] + v
		}
	}
	return result
}
