package rule_engine

import (
	"context"
	"github.com/pkg/errors"
)

const (
	RULE_STATUS_CREATED = "created"
	RULE_STATUS_STARTED = "started"
	RULE_STATUS_STOPPED = "stopped"
	RULE_STATUS_UNKNOWN = "unknown"
)

var (
	// ErrConflict indicates usage of the existing email during account
	// registration.
	ErrConflict = errors.New("email already taken")

	// ErrMalformedEntity indicates malformed entity specification
	// (e.g. invalid realmname or password).
	ErrMalformedEntity = errors.New("malformed entity specification")

	// ErrUnauthorizedAccess indicates missing or invalid credentials provided
	// when accessing a protected resource.
	ErrUnauthorizedAccess = errors.New("missing or invalid credentials provided")

	// ErrNotFound indicates a non-existent entity request.
	ErrNotFound = errors.New("non-existent entity")

	// ErrruleEngineNotFound indicates a non-existent realm request.
	ErrruleEngineNotFound = errors.New("non-existent ruleEngine")

	// ErrScanMetadata indicates problem with metadata in db.
	ErrScanMetadata = errors.New("Failed to scan metadata")

	// ErrMissingEmail indicates missing email for password reset request.
	ErrMissingEmail = errors.New("missing email for password reset")

	// ErrUnauthorizedPrincipal indicate the pricipal can not be recognized
	ErrUnauthorizedPrincipal = errors.New("unauthorized principal")
)

//Service service
type Service interface {
	AddNewruleEngine(context.Context, RuleEngine) error
	GetruleEngineInfo(context.Context, string) (RuleEngine, error)
	UpdateruleEngine(context.Context, RuleEngine) (RuleEngine, error)
	RevokeruleEngine(context.Context, string) error
	ListruleEngine(context.Context, uint64, uint64) (RuleEnginePage, error)
	UpdateruleEngineStatus(context.Context, string, string) error
}

/*
var _ Service = (*ruleEngineService)(nil)

type ruleEngineService struct {
	ruleEngines      RuleEngineRepository
	mutex            sync.RWMutex
	instanceManager  instanceManager
	ruleEnginesCache RuleEngineCache
}

//New new
func New(ruleEngines RuleEngineRepository, instancemanager instanceManager, ruleEngineCache RuleEngineCache) Service {
	return &ruleEngineService{
		ruleEngines:      ruleEngines,
		instanceManager:  instancemanager,
		ruleEnginesCache: ruleEngineCache,
	}
}

func (svc ruleEngineService) AddNewruleEngine(ctx context.Context, ruleEngine RuleEngine) error {
	return svc.ruleEngines.Save(ctx, ruleEngine)
}

func (svc ruleEngineService) GetruleEngineInfo(ctx context.Context, ruleEngineID string) (RuleEngine, error) {
	ruleEngine, err := svc.ruleEngines.Retrieve(ctx, ruleEngineID)
	if err != nil {
		return RuleEngine{}, errors.Wrap(ErrruleEngineNotFound, err.Error())
	}

	return ruleEngine, nil
}

func (svc ruleEngineService) UpdateruleEngine(ctx context.Context, ruleEngine RuleEngine) (RuleEngine, error) {

	old_ruleEngine, err := svc.ruleEngines.Retrieve(ctx, ruleEngine.ID)
	if err != nil {
		return RuleEngine{}, errors.Wrap(ErrruleEngineNotFound, err.Error())
	}
	if old_ruleEngine.Status == RULE_STATUS_STARTED {
		return RuleEngine{}, status.Error(codes.FailedPrecondition, "")
	}

	return svc.ruleEngines.Update(ctx, ruleEngine)
}

func (svc ruleEngineService) RevokeruleEngine(ctx context.Context, ruleEngineID string) error {

	ruleEngine, err := svc.ruleEngines.Retrieve(ctx, ruleEngineID)
	if err != nil {
		return errors.Wrap(ErrruleEngineNotFound, err.Error())
	}
	if ruleEngine.Status == RULE_STATUS_STARTED {
		return status.Error(codes.FailedPrecondition, "")
	}

	return svc.ruleEngines.Revoke(ctx, ruleEngineID)
}

func (svc ruleEngineService) ListruleEngine(ctx context.Context, offset uint64, limit uint64) (RuleEnginePage, error) {
	return svc.ruleEngines.List(ctx, offset, limit)
}

func (svc ruleEngineService) UpdateruleEngineStatus(ctx context.Context, ruleEngineID string, updatestatus string) error {

	ruleEngine, err := svc.ruleEngines.Retrieve(ctx, ruleEngineID)
	if err != nil {
		return errors.Wrap(ErrruleEngineNotFound, err.Error())
	}

	switch updatestatus {
	case UPDATE_RULE_STATUS_START:
		if ruleEngine.Status != RULE_STATUS_CREATED && ruleEngine.Status != RULE_STATUS_STOPPED {
			return status.Error(codes.FailedPrecondition, "")
		}

		return svc.instanceManager.startRuleEngine(&ruleEngine)
	case UPDATE_RULE_STATUS_STOP:
		if ruleEngine.Status != RULE_STATUS_STARTED {
			return status.Error(codes.FailedPrecondition, "")
		}

		return svc.instanceManager.stopRuleEngine(&ruleEngine)
	}
	return nil
}
*/
