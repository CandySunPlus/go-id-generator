package idgen

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
)

// IDGenerator struct
type IDGenerator struct {
	client *redis.Client
	hash   string
}

// NewIDGenerator construct
func NewIDGenerator(host string, port int, hash string) *IDGenerator {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: "",
		DB:       0,
	})

	return &IDGenerator{client, hash}
}

// Next func
func (idGen *IDGenerator) Next(tableName string, shardID int) (int64, error) {
	result, err := idGen.client.EvalSha(idGen.hash, []string{tableName, strconv.Itoa(shardID)}).Result()
	var paramSlice []int64

	if params, ok := result.([]interface{}); ok {
		for _, param := range params {
			paramSlice = append(paramSlice, param.(int64))
		}
		return buildID(paramSlice...), nil
	}

	if err != nil {
		return 0, err
	}

	return 0, errors.New("id generator result parse error")
}

func buildID(args ...int64) int64 {
	var id int64
	for index, arg := range args {
		switch index {
		case 0:
			id = arg * 1000
			break
		case 1:
			id += arg / 1000
			break
		case 2:
			id = id<<(12+10) + (arg << 10)
			break
		case 3:
			id += arg
			break
		}
	}
	return id
}

// ParseID method
func ParseID(id int64) []int64 {
	miliSecond := id >> 22
	// 保留中间 12bit
	sharedID := (id & (0xFFF << 10)) >> 10
	// 最后 10bit
	seq := id & 0x3FF

	return []int64{miliSecond, sharedID, seq}
}
