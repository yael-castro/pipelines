package logic

func New(repo Repository) Logic {
	return logic{
		Repository: repo,
	}
}

type logic struct {
	Repository
}
