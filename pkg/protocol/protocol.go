package protocol

import "github.com/BladekTech/blaze/internal/blaze/util"

type Command struct {
	Name string
	Args []string
}

type Status int
type Data struct {
	inner string
}

const DEFAULT_PORT int16 = 7854

const (
	STATUS_OK Status = iota
	STATUS_MALFORMED_REQUEST
	STATUS_INVALID_KEY
	STATUS_NO_SUCH_COMMAND
	STATUS_NO_SUCH_KEY
	STATUS_KEY_ALREADY_EXISTS
)

const (
	CMD_PING string = "ping"
	CMD_GET  string = "get"
	CMD_SET  string = "set"
)

type Result struct {
	Status Status
	Error  *string
	Data   Data
}

func (result Result) IsError() bool {
	return result.Status != STATUS_OK
}

func (data Data) ToBytes() []byte {
	return util.StrToByteSlice(data.inner)
}

func NewData(inner string) Data {
	return Data{
		inner,
	}
}
