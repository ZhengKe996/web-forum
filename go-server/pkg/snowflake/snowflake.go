package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"time"
)

var node *snowflake.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2023-01-01", startTime)
	if err != nil {
		return err
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	return err
}

func GenID() string {
	return node.Generate().String()
}
