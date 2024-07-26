package snowflake

import (
	"fmt"
	"strconv"
	"time"
)

const (
	_epoch           = 1420070400000
	_timestampBits   = 42
	_workerBits      = 5
	_processBits     = 5
	_sequenceBits    = 12
	_maxWorkerID     = (1 << _workerBits) - 1
	_maxProcessID    = (1 << _processBits) - 1
	_maxSequence     = (1 << _sequenceBits) - 1
	_defaultSequence = 0
)

var (
	defaultWorkerID  uint64 = 1
	defaultProcessID uint64 = 1
	generator        *SnowflakeGenerator
)

type SnowflakeGenerator struct {
	WorkerID   uint64
	ProcessID  uint64
	Sequence   uint64
	lastMillis uint64
}

func (g *SnowflakeGenerator) GenerateDefaultSnowflake() uint64 {
	millis := uint64(time.Now().UnixNano() / 1e6)

	if millis == g.lastMillis {
		g.Sequence = (g.Sequence + 1) & _maxSequence
		if g.Sequence == 0 {
			// Sequence rollover, wait until next millisecond
			for millis <= g.lastMillis {
				millis = uint64(time.Now().UnixNano() / 1e6)
			}
		}
	} else {
		g.Sequence = _defaultSequence
	}

	g.lastMillis = millis

	return ((millis - _epoch) << (_workerBits + _processBits + _sequenceBits)) |
		((g.WorkerID & _maxWorkerID) << (_processBits + _sequenceBits)) |
		((g.ProcessID & _maxProcessID) << _sequenceBits) |
		(g.Sequence & _maxSequence)
}

func init() {
	generator = &SnowflakeGenerator{
		WorkerID:  defaultWorkerID,
		ProcessID: defaultProcessID,
		Sequence:  _defaultSequence,
	}
}

// SetDefaultWorkerID sets the default worker ID
func SetDefaultWorkerID(workerID uint64) {
	defaultWorkerID = workerID
}

// SetDefaultProcessID sets the default process ID
func SetDefaultProcessID(processID uint64) {
	defaultProcessID = processID
}

// Generate generates a snowflake using the default generator
func Generate() uint64 {
	return generator.GenerateDefaultSnowflake()
}

// IsValidSnowflake checks if the given value is a valid  snowflake
func IsValidSnowflake(value interface{}) bool {
	var id uint64

	switch v := value.(type) {
	case int:
		id = uint64(v)
	case int32:
		id = uint64(v)
	case uint64:
		id = v
	case string:
		// Try to parse the string as an integer
		intValue, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return false // Not a valid integer
		}
		id = intValue
	default:
		// If the type is not int, int32, uint64, or string, it's not a valid type for a snowflake
		return false
	}

	// Extract timestamp, worker ID, process ID, and sequence
	timestamp := (id >> (_workerBits + _processBits + _sequenceBits)) + _epoch
	workerID := (id >> (_processBits + _sequenceBits)) & _maxWorkerID
	processID := (id >> _sequenceBits) & _maxProcessID
	sequence := id & _maxSequence

	// Validate timestamp (within a reasonable range)
	now := uint64(time.Now().UnixNano() / 1e6)
	if timestamp < _epoch || timestamp > now {
		return false
	}

	// Validate worker ID, process ID, and sequence
	return workerID <= _maxWorkerID && processID <= _maxProcessID && sequence <= _maxSequence
}

func GetTimestampFromSnowflake(snowflake uint64) (uint64, error) {
	// Extract timestamp from the snowflake
	timestamp := (snowflake >> (_workerBits + _processBits + _sequenceBits)) + _epoch

	// Validate the timestamp
	now := uint64(time.Now().UnixNano() / 1e6)
	if timestamp < _epoch || timestamp > now {
		return 0, fmt.Errorf("invalid timestamp extracted from snowflake")
	}

	return timestamp, nil
}
