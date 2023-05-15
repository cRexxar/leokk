package snowflake

import (
	"math/rand"
	"time"

	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

func init() {
	sf = createIdGenerator()
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func createIdGenerator() *sonyflake.Sonyflake {
	config := sonyflake.Settings{
		MachineID: getMachineId,
	}
	return sonyflake.NewSonyflake(config)
}

func getMachineId() (uint16, error) {
	return uint16(1), nil
}

func Next() (id int64) {
	var i uint64
	if sf != nil {
		i, _ = sf.NextID()
		id = int64(i)
	} else {
		id = rand.Int63()
	}
	return
}

func GetOne() int64 {
	return Next()
}
