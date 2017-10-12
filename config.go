package paxi

import (
	"encoding/json"
	"os"
	"paxi/glog"
	"strconv"
)

// default values
const (
	PORT             = 1735
	HTTP_PORT        = 8080
	CHAN_BUFFER_SIZE = 1024 * 10000
	BUFFER_SIZE      = 1024 * 10000
)

// Algorithm type
type Algorithm int

// add algorithm type here
const (
	WPaxos Algorithm = iota
	EPaxos
	KPaxos
	WanKeeper
	Cosmos
)

type Config struct {
	ID             ID            `json:"id"`
	Addrs          map[ID]string `json:"address"`      // address for node communication
	HTTPAddrs      map[ID]string `json:"http_address"` // address for client
	Algorithm      Algorithm     `json:"algorithm"`
	F              int           `json:"f"` // number of failure nodes
	Threshold      int           `json:"threshold"`
	BackOff        int           `json:"backoff"`
	Thrifty        bool          `json:"thrifty"`
	ChanBufferSize int           `json:"chan_buffer_size"`
	BufferSize     int           `json:"buffer_size"`
	ConfigFile     string        `json:"file"`
	Consistency    int           `json:"consistency"`
	// Transport      string        `json:"transport"`
	// RecvRoutines   int           `json:"recv_routines"`
	// Codec          string        `json:"codec"`
	// Batching       bool          `json:"batching"`
}

func MakeDefaultConfig() *Config {
	id := NewID(0, 0)
	config := new(Config)
	config.ID = NewID(1, 1)
	config.Addrs = map[ID]string{id: "chan://*:" + strconv.Itoa(PORT)}
	config.Algorithm = WPaxos
	config.ChanBufferSize = CHAN_BUFFER_SIZE
	config.BufferSize = BUFFER_SIZE
	config.ConfigFile = "config.json"
	return config
}

// String is implemented to print the config
func (c *Config) String() string {
	config, err := json.Marshal(c)
	if err != nil {
		glog.Errorln(err)
	}
	return string(config)
	//return fmt.Sprintf("Config[MyID:%s,Address:%s]", c.ID.String(), c.Addrs[c.ID])
}

func (c *Config) Load() error {
	file, err := os.Open(c.ConfigFile)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	return decoder.Decode(c)
}

func (c *Config) Save() error {
	file, err := os.Create(c.ConfigFile)
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(file)
	return encoder.Encode(c)
}