/* other utility functions */

package utils

import (
	"os"
	"reflect"
	"time"
)


func FileExists(filename string) bool {

	// checks if a file exists and is not a directory

	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}


func TsToString(ts time.Time) string {

	// convert a given time stamp `ts` to string of format `2006-01-02T15:04:05Z`

	return ts.Format("2006-01-02T15:04:05Z")
	
}


func PartitionList(list []string, chunkSize int) [][]string {

	// powered by ChatGPT

	var partitions [][]string

	for i := 0; i < len(list); i += chunkSize {
		end := i + chunkSize
		if end > len(list) {
			end = len(list)
		}
		partitions = append(partitions, list[i:end])
	}

	return partitions
}


func PartitionGenericList(list interface{}, chunkSize int) [][]interface{} {

	// works on any list type

	listValue := reflect.ValueOf(list)
	if listValue.Kind() != reflect.Slice {
		panic("Input is not a slice")
	}

	var partitions [][]interface{}

	for i := 0; i < listValue.Len(); i += chunkSize {
		end := i + chunkSize
		if end > listValue.Len() {
			end = listValue.Len()
		}
		partitions = append(partitions, sliceToInterfaceSlice(listValue.Slice(i, end)))
	}

	return partitions
}

func sliceToInterfaceSlice(slice reflect.Value) []interface{} {
	interfaceSlice := make([]interface{}, slice.Len())
	for i := 0; i < slice.Len(); i++ {
		interfaceSlice[i] = slice.Index(i).Interface()
	}
	return interfaceSlice
}