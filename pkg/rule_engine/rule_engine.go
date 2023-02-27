package rule_engine

import (
	"context"
	"time"
)

const (
	UPDATE_RULE_STATUS_START = "start"
	UPDATE_RULE_STATUS_STOP  = "stop"
)

//RuleEngine RuleEngine
type RuleEngine struct {
	Name         string
	ID           string
	Description  string
	DebugMode    bool
	Status       string
	Payload      []byte
	Root         bool
	Channel      string
	SubTopic     string
	CreateAt     time.Time
	LastUpdateAt time.Time
}

type PageMetadata struct {
	Total  uint64
	Offset uint64
	Limit  uint64
	Name   string
}

type RuleEnginePage struct {
	PageMetadata
	RuleChains []RuleEngine
}

// Validate returns an error if representtation is invalid
func (r RuleEngine) Validate() error {
	if r.ID == "" {
		return ErrMalformedEntity
	}
	return nil
}

//RuleEngineRepository specifies realm persistence API
type RuleEngineRepository interface {
	//Save save the RuleEngine
	Save(context.Context, RuleEngine) error

	//Update the RuleEngine
	Update(context.Context, RuleEngine) (RuleEngine, error)

	//Retrieve return RuleEngine by RuleEngine id
	Retrieve(context.Context, string) (RuleEngine, error)

	//Revoke remove RuleEngine by RuleEngine id
	Revoke(context.Context, string) error

	//List return all RuleEngines
	List(context.Context, uint64, uint64) (RuleEnginePage, error)
}

// RuleEngineCache contains thing caching interface.
type RuleEngineCache interface {
	// Save stores pair thing key, thing id.
	Save(context.Context, string, string) error

	// ID returns thing ID for given key.
	ID(context.Context, string) (string, error)

	// Remove thing from cache.
	Remove(context.Context, string) error
}
