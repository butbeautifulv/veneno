package intelligence

import "github.com/butbeautifulv/veneno/pkg/decision"

// Re-exports for engage wiring; decision logic lives in pkg/decision.
type (
	DecisionEngine  = decision.DecisionEngine
	TargetProfile   = decision.TargetProfile
	OptimizeContext = decision.OptimizeContext
)

var (
	DefaultDecisionEngine = decision.DefaultDecisionEngine
	AttackPatterns        = decision.AttackPatterns
	SelectPatternKey      = decision.SelectPatternKey
	BuildTargetProfile    = decision.BuildTargetProfile
)
