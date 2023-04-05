package config

// DagOptions is the options for the build command
type DagOptions struct {
	// Set is the set flag
	Set []string
	// Values is the values flag
	Values []string
	// SkipDeps is the skip deps flag
	SkipDeps bool
	// DetailedExitcode is the detailed exit code
	DetailedExitcode bool
	// IncludeTests is the include tests flag
	IncludeTests bool
	// SkipNeeds is the include crds flag
	SkipNeeds bool
	// IncludeNeeds is the include needs flag
	IncludeNeeds bool
	// IncludeTransitiveNeeds is the include transitive needs flag
	IncludeTransitiveNeeds bool
	// SkipDagOnInstall is the skip diff on install flag
	SkipDagOnInstall bool
	// ShowSecrets is the show secrets flag
	ShowSecrets bool
	// NoHooks skips hooks during diff
	NoHooks bool
	// Suppress is the suppress flag
	Suppress []string
	// SuppressSecrets is the suppress secrets flag
	SuppressSecrets bool
	// Concurrency is the concurrency flag
	Concurrency int
	// Validate is the validate flag
	Validate bool
	// Context is the context flag
	Context int
	// Output is output flag
	Output string
	// ReuseValues is true if the helm command should reuse the values
	ReuseValues bool
	// ResetValues is true if helm command should reset values to charts' default
	ResetValues bool
	// Propagate '--post-renderer' to helmv3 template and helm install
	PostRenderer string
}

// NewDagOptions creates a new Apply
func NewDagOptions() *DagOptions {
	return &DagOptions{}
}

// DagImpl is impl for applyOptions
type DagImpl struct {
	*GlobalImpl
	*DagOptions
}

// NewDagImpl creates a new DagImpl
func NewDagImpl(g *GlobalImpl, t *DagOptions) *DagImpl {
	return &DagImpl{
		GlobalImpl: g,
		DagOptions: t,
	}
}

// Concurrency returns the concurrency
func (t *DagImpl) Concurrency() int {
	return t.DagOptions.Concurrency
}

// IncludeNeeds returns the include needs
func (t *DagImpl) IncludeNeeds() bool {
	return t.DagOptions.IncludeNeeds || t.IncludeTransitiveNeeds()
}

// IncludeTransitiveNeeds returns the include transitive needs
func (t *DagImpl) IncludeTransitiveNeeds() bool {
	return t.DagOptions.IncludeTransitiveNeeds
}

// Set returns the Set
func (t *DagImpl) Set() []string {
	return t.DagOptions.Set
}

// SkipDeps returns the skip deps
func (t *DagImpl) SkipDeps() bool {
	return t.DagOptions.SkipDeps
}

// SkipNeeds returns the skip needs
func (t *DagImpl) SkipNeeds() bool {
	if !t.IncludeNeeds() {
		return t.DagOptions.SkipNeeds
	}

	return false
}

// Validate returns the validate
func (t *DagImpl) Validate() bool {
	return t.DagOptions.Validate
}

// Values returns the values
func (t *DagImpl) Values() []string {
	return t.DagOptions.Values
}

// Context returns the context
func (t *DagImpl) Context() int {
	return t.DagOptions.Context
}

// DetailedExitCode returns the detailed exit code
func (t *DagImpl) DetailedExitcode() bool {
	return t.DagOptions.DetailedExitcode
}

// Output returns the output
func (t *DagImpl) DagOutput() string {
	return t.DagOptions.Output
}

// IncludeTests returns the include tests
func (t *DagImpl) IncludeTests() bool {
	return t.DagOptions.IncludeTests
}

// ShowSecrets returns the show secrets
func (t *DagImpl) ShowSecrets() bool {
	return t.DagOptions.ShowSecrets
}

// NoHooks skips hooks.
func (t *DagImpl) NoHooks() bool {
	return t.DagOptions.NoHooks
}

// ShowCRDs returns the show crds
func (t *DagImpl) SkipCRDs() bool {
	return false
}

// SkipDagOnInstall returns the skip diff on install
func (t *DagImpl) SkipDagOnInstall() bool {
	return t.DagOptions.SkipDagOnInstall
}

// Suppress returns the suppress
func (t *DagImpl) Suppress() []string {
	return t.DagOptions.Suppress
}

// SuppressDag returns the suppress diff
func (t *DagImpl) SuppressDag() bool {
	return false
}

// SuppressSecrets returns the suppress secrets
func (t *DagImpl) SuppressSecrets() bool {
	return t.DagOptions.SuppressSecrets
}

// ReuseValues returns the ReuseValues.
func (t *DagImpl) ReuseValues() bool {
	if !t.ResetValues() {
		return t.DagOptions.ReuseValues
	}
	return false
}

func (t *DagImpl) ResetValues() bool {
	return t.DagOptions.ResetValues
}

// PostRenderer returns the PostRenderer.
func (t *DagImpl) PostRenderer() string {
	return t.DagOptions.PostRenderer
}
