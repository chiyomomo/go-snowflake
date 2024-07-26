# Snowflake ID Generator

This package provides a simple and efficient implementation of the Snowflake ID generation algorithm in Go.

## Installation

To use this package, Go must be installed on the machine. Install the package using `go get`:

```sh
go get github.com/chiyomomo/snowflake
```

## Usage

#### Importing the Package

```go
import "github.com/chiyomomo/snowflake"
```

#### Generating a Snowflake ID

```go
id := snowflake.Generate()
fmt.Printf("Generated Snowflake ID: %d\n", id)
```

#### Setting Worker and Process IDs

Configure the default worker and process IDs before generating Snowflake IDs:

```go
snowflake.SetDefaultWorkerID(1)
snowflake.SetDefaultProcessID(1)
id := snowflake.Generate()
fmt.Printf("Generated Snowflake ID with custom worker and process IDs: %d\n", id)
```

#### Validating a Snowflake ID

Check if a given value is a valid Snowflake ID:

```go
isValid := snowflake.IsValidSnowflake(id)
fmt.Printf("Is valid Snowflake ID: %t\n", isValid)
```

#### Extracting Timestamp from a Snowflake ID

Extract the timestamp from a given Snowflake ID:

```go
timestamp, err := snowflake.GetTimestampFromSnowflake(id)
if err != nil {
    fmt.Printf("Error extracting timestamp: %s\n", err)
} else {
    fmt.Printf("Timestamp from Snowflake ID: %d\n", timestamp)
}
```
