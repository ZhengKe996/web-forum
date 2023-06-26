package snowflake

import (
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func Init(startTime string, machineID int64) (err error) {
	node, err = snowflake.NewNode(machineID)
	if err != nil {
		return err
	}
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}
