package snowflake

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpack(t *testing.T) {
	assert.Equal(
		t,
		SnowflakeStructure{
			Timestamp:         1568274192129,
			InternalWorkerId:  1,
			InternalProcessId: 0,
			Increment:         10,
		},
		DiscordBuilder.From(621611758141964298).Structure(),
	)
	assert.Equal(
		t,
		SnowflakeStructure{
			Timestamp:         1658487387287,
			InternalWorkerId:  2,
			InternalProcessId: 1,
			Increment:         71,
		},
		DiscordBuilder.From(999993323446079559).Structure(),
	)
	assert.Equal(
		t,
		SnowflakeStructure{
			Timestamp:         1675971236426,
			InternalWorkerId:  2,
			InternalProcessId: 2,
			Increment:         1,
		},
		DiscordBuilder.From(1073325901825187841).Structure(),
	)
	assert.Equal(
		t,
		SnowflakeStructure{
			Timestamp:         1636450447419,
			InternalWorkerId:  1,
			InternalProcessId: 1,
			Increment:         40,
		},
		DiscordBuilder.From(907563698409836584).Structure(),
	)
	assert.Equal(
		t,
		SnowflakeStructure{
			Timestamp:         1631877663788,
			InternalWorkerId:  1,
			InternalProcessId: 0,
			Increment:         20,
		},
		DiscordBuilder.From(888384053735194644).Structure(),
	)
	assert.Equal(
		t,
		SnowflakeStructure{
			Timestamp:         1636888207754,
			InternalWorkerId:  1,
			InternalProcessId: 2,
			Increment:         31,
		},
		DiscordBuilder.From(909399798333972511).Structure(),
	)
	assert.Equal(
		t,
		SnowflakeStructure{
			Timestamp:         1623344062350,
			InternalWorkerId:  1,
			InternalProcessId: 0,
			Increment:         80,
		},
		DiscordBuilder.From(852591535089385552).Structure(),
	)
}

func TestPack(t *testing.T) {
	assert.Equal(
		t,
		uint64(621611758141964298),
		DiscordBuilder.Make(&SnowflakeStructure{
			Timestamp:         1568274192129,
			InternalWorkerId:  1,
			InternalProcessId: 0,
			Increment:         10,
		}).Value(),
	)
	assert.Equal(
		t,
		uint64(999993323446079559),
		DiscordBuilder.Make(&SnowflakeStructure{
			Timestamp:         1658487387287,
			InternalWorkerId:  2,
			InternalProcessId: 1,
			Increment:         71,
		}).Value(),
	)
	assert.Equal(
		t,
		uint64(1073325901825187841),
		DiscordBuilder.Make(&SnowflakeStructure{
			Timestamp:         1675971236426,
			InternalWorkerId:  2,
			InternalProcessId: 2,
			Increment:         1,
		}).Value(),
	)
	assert.Equal(
		t,
		uint64(907563698409836584),
		DiscordBuilder.Make(&SnowflakeStructure{
			Timestamp:         1636450447419,
			InternalWorkerId:  1,
			InternalProcessId: 1,
			Increment:         40,
		}).Value(),
	)
	assert.Equal(
		t,
		uint64(888384053735194644),
		DiscordBuilder.Make(&SnowflakeStructure{
			Timestamp:         1631877663788,
			InternalWorkerId:  1,
			InternalProcessId: 0,
			Increment:         20,
		}).Value(),
	)
	assert.Equal(
		t,
		uint64(909399798333972511),
		DiscordBuilder.Make(&SnowflakeStructure{
			Timestamp:         1636888207754,
			InternalWorkerId:  1,
			InternalProcessId: 2,
			Increment:         31,
		}).Value(),
	)
	assert.Equal(
		t,
		uint64(852591535089385552),
		DiscordBuilder.Make(&SnowflakeStructure{
			Timestamp:         1623344062350,
			InternalWorkerId:  1,
			InternalProcessId: 0,
			Increment:         80,
		}).Value(),
	)
}

func TestJson(t *testing.T) {
	d, err := json.Marshal([]*Snowflake{
		DiscordBuilder.From(621611758141964298),
		DiscordBuilder.From(999993323446079559),
		DiscordBuilder.From(1073325901825187841),
		DiscordBuilder.From(907563698409836584),
		DiscordBuilder.From(888384053735194644),
		DiscordBuilder.From(909399798333972511),
		DiscordBuilder.From(852591535089385552),
	})
	if err != nil {
		panic(err)
	}
	assert.Equal(
		t,
		`["621611758141964298","999993323446079559","1073325901825187841","907563698409836584","888384053735194644","909399798333972511","852591535089385552"]`,
		string(d),
	)
	var res []Snowflake
	err = json.Unmarshal(
		[]byte(`["621611758141964298","999993323446079559","1073325901825187841","907563698409836584","888384053735194644","909399798333972511","852591535089385552"]`),
		&res,
	)
	if err != nil {
		panic(err)
	}
	assert.Equal(
		t,
		[]Snowflake{
			NewSnowflake(621611758141964298, nil),
			NewSnowflake(999993323446079559, nil),
			NewSnowflake(1073325901825187841, nil),
			NewSnowflake(907563698409836584, nil),
			NewSnowflake(888384053735194644, nil),
			NewSnowflake(909399798333972511, nil),
			NewSnowflake(852591535089385552, nil),
		},
		res,
	)
}
