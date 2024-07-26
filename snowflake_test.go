package snowflake

import (
	"testing"
	"time"
)

// Test the generation of a Snowflake ID
func TestGenerate(t *testing.T) {
	id := Generate()
	if id == 0 {
		t.Errorf("Expected a non-zero snowflake ID")
	}
}

// Test the validation of a Snowflake ID
func TestIsValidSnowflake(t *testing.T) {
	id := Generate()

	if !IsValidSnowflake(id) {
		t.Errorf("Expected ID to be a valid snowflake")
	}

	if IsValidSnowflake("invalid_id") {
		t.Errorf("Expected string 'invalid_id' to be an invalid snowflake")
	}
}

// Test the extraction of timestamp from a Snowflake ID
func TestGetTimestampFromSnowflake(t *testing.T) {
	id := Generate()

	timestamp, err := GetTimestampFromSnowflake(id)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	now := uint64(time.Now().UnixNano() / 1e6)
	if timestamp > now {
		t.Errorf("Expected timestamp to be less than or equal to now")
	}

	if timestamp < _epoch {
		t.Errorf("Expected timestamp to be greater than or equal to the epoch")
	}
}
