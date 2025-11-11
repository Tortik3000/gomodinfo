package moduleinfo

// UseCase wires ports to perform the use case
type UseCase struct {
	repo    repoContentProvider
	parser  goModParser
	checker versionChecker
}

func NewUseCase(
	repo repoContentProvider,
	parser goModParser,
	checker versionChecker,
) *UseCase {
	return &UseCase{
		repo:    repo,
		parser:  parser,
		checker: checker,
	}
}
