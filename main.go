package main

import (
    "fmt"
    "os"
    "github.com/cristim/ec2-instances-info"
)

func main() {
	data, err := ec2instancesinfo.Data()

    if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
    
    // Print all available instance types
    for _, i := range *data {
        fmt.Println("Instance type", i.InstanceType)
    }
}
