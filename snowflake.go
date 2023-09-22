package snowflake

import (
	"encoding/json"
	"os"
	"strconv"
	"time"
)

type SnowflakeStructure struct {
	Timestamp         uint64
	InternalWorkerId  uint8
	InternalProcessId uint8
	Increment         uint16
}

type SnowflakeBuilder struct {
	epoch uint64
}

func NewBuilder(epoch uint64) *SnowflakeBuilder {
	return &SnowflakeBuilder{
		epoch: epoch,
	}
}

func (sb SnowflakeBuilder) Epoch() uint64 {
	return sb.epoch
}

func (sb *SnowflakeBuilder) From(value uint64) *Snowflake {
	return &Snowflake{
		value:   value,
		builder: sb,
	}
}

func (sb *SnowflakeBuilder) Make(options *SnowflakeStructure) Snowflake {
	res := Snowflake{
		value:   0,
		builder: sb,
	}
	if options != nil {
		res.SetTimestamp(options.Timestamp)
		res.SetInternalWorkerId(options.InternalWorkerId)
		res.SetInternalProcessId(options.InternalProcessId)
		res.SetIncrement(options.Increment)
	}
	return res
}

func (sb *SnowflakeBuilder) DefaultGenerator(internalWorkerId uint8) *DefaultSnowflakeGenerator {
	return &DefaultSnowflakeGenerator{
		InternalWorkerId: internalWorkerId,
		Builder:          sb,
		Increment:        0,
	}
}

type Snowflake struct {
	value   uint64
	builder *SnowflakeBuilder
}

type SnowflakeGenerator[T any] interface {
	Make(data T) Snowflake
}

type DefaultSnowflakeGenerator struct {
	InternalWorkerId uint8
	Builder          *SnowflakeBuilder
	Increment        uint16
}

func (dsg *DefaultSnowflakeGenerator) Make(data uintptr) Snowflake {
	res := dsg.Builder.Make(&SnowflakeStructure{
		Timestamp:         uint64(time.Now().UTC().UnixMilli()),
		InternalWorkerId:  dsg.InternalWorkerId,
		InternalProcessId: uint8(os.Getpid()),
		Increment:         dsg.Increment,
	})
	dsg.Increment++
	return res
}

func (s Snowflake) Builder() *SnowflakeBuilder {
	return s.builder
}

func (s Snowflake) Value() uint64 {
	return s.value
}

func (s *Snowflake) SetBuilder(sb *SnowflakeBuilder) *Snowflake {
	s.builder = sb
	return s
}

func (s *Snowflake) SetValue(value uint64) *Snowflake {
	s.value = value
	return s
}

func (s Snowflake) Timestamp() uint64 {
	return (s.value >> 22) + s.builder.epoch
}

func (s Snowflake) InternalWorkerId() uint8 {
	return uint8((s.value & 0x3E0000) >> 17)
}

func (s Snowflake) InternalProcessId() uint8 {
	return uint8((s.value & 0x1F000) >> 12)
}

func (s Snowflake) Increment() uint16 {
	return uint16(s.value & 0xFFF)
}

func (s Snowflake) Structure() SnowflakeStructure {
	return SnowflakeStructure{
		Timestamp:         s.Timestamp(),
		InternalWorkerId:  s.InternalWorkerId(),
		InternalProcessId: s.InternalProcessId(),
		Increment:         s.Increment(),
	}
}

func (s *Snowflake) SetTimestamp(timestamp uint64) *Snowflake {
	s.value = (s.value & uint64(0x3FFFFF)) | ((timestamp - s.builder.epoch) << 22)
	return s
}

func (s *Snowflake) SetInternalWorkerId(internalWorkerId uint8) *Snowflake {
	s.value = (s.value & uint64(0xFFFFFFFFFFC1FFFF)) | (uint64(internalWorkerId&0x1F) << 17)
	return s
}

func (s *Snowflake) SetInternalProcessId(internalProcessId uint8) *Snowflake {
	s.value = (s.value & uint64(0xFFFFFFFFFFFE0FFF)) | (uint64(internalProcessId&0x1F) << 12)
	return s
}

func (s *Snowflake) SetIncrement(increment uint16) *Snowflake {
	s.value = (s.value & uint64(0xFFFFFFFFFFFFF000)) | uint64(increment&0xFFF)
	return s
}

func NewSnowflake(value uint64, builder *SnowflakeBuilder) Snowflake {
	return Snowflake{
		value:   value,
		builder: builder,
	}
}

func (s Snowflake) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatUint(s.value, 10))
}

func (s *Snowflake) UnmarshalJSON(data []byte) error {
	var vs string
	if err := json.Unmarshal(data, &vs); err != nil {
		return err
	}
	value, err := strconv.ParseUint(vs, 10, 64)
	if err != nil {
		return err
	}
	s.value = value
	s.builder = nil
	return nil
}

const (
	DiscordEpoch      uint64 = 1420070400000
)

var (
	DiscordBuilder    *SnowflakeBuilder = NewBuilder(DiscordEpoch)
	DiscordGenerator  *DefaultSnowflakeGenerator

	initalized        bool = false
)

func Initialize() {
	if initalized {
		panic("Initialized")
	}
	DiscordGenerator = DiscordBuilder.DefaultGenerator(0)
	initalized = true
}
