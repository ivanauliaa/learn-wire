package user

import (
	"database/sql"
	"sync"

	"github.com/google/wire"
	"github.com/ivanauliaa/learn-wire/domain"
)

var (
	hdl     *handler
	hdlOnce sync.Once

	svc     *service
	svcOnce sync.Once

	repo     *repository
	repoOnce sync.Once

	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideHandler,
		ProvideService,
		ProvideRepository,

		wire.Bind(new(domain.UserHandlerInterface), new(*handler)),
		wire.Bind(new(domain.UserServiceInterface), new(*service)),
		wire.Bind(new(domain.UserRepositoryInterface), new(*repository)),
	)
)

func ProvideRepository(db *sql.DB) *repository {
	repoOnce.Do(func() {
		repo = &repository{
			db: db,
		}
	})

	return repo
}

func ProvideService(repo domain.UserRepositoryInterface) *service {
	svcOnce.Do(func() {
		svc = &service{
			repo: repo,
		}
	})

	return svc
}

func ProvideHandler(svc domain.UserServiceInterface) *handler {
	hdlOnce.Do(func() {
		hdl = &handler{
			svc: svc,
		}
	})

	return hdl
}
