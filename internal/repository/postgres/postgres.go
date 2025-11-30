package postgres

import (
	"blog/config"
	"blog/internal/entity"
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Predicate string

var (
	EqualPredicate              = Predicate("=")
	NotEqualPredicate           = Predicate("<>")
	GreaterThanPredicate        = Predicate(">")
	GreaterThanOrEqualPredicate = Predicate(">=")
	SmallerThanPredicate        = Predicate("<")
	SmallerThanOrEqualPredicate = Predicate("<=")
	LikePredicate               = Predicate("LIKE")
	ILikePredicate              = Predicate("ILIKE") // PostgreSQL-specific case-insensitive matching
)

var _ Repo = (*dbRepo)(nil)

type Repo interface {
	i()
	GetDbR() *gorm.DB
	GetDbW() *gorm.DB
	DbRClose() error
	DbWClose() error
}

type dbRepo struct {
	DbR *gorm.DB
	DbW *gorm.DB
}

func New() (Repo, error) {
	cfg := config.Conf.Postgres
	dbr, err := dbConnect(cfg.Username, cfg.Password, fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), cfg.Dbname, cfg.SSLMode)
	if err != nil {
		return nil, err
	}

	dbw, err := dbConnect(cfg.Username, cfg.Password, fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), cfg.Dbname, cfg.SSLMode)
	if err != nil {
		return nil, err
	}

	return &dbRepo{
		DbR: dbr,
		DbW: dbw,
	}, nil
}

func (d *dbRepo) i() {}

func (d *dbRepo) GetDbR() *gorm.DB {
	return d.DbR
}

func (d *dbRepo) GetDbW() *gorm.DB {
	return d.DbW
}

func (d *dbRepo) DbRClose() error {
	sqlDB, err := d.DbR.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (d *dbRepo) DbWClose() error {
	sqlDB, err := d.DbW.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func dbConnect(user, pass, addr, dbName, sslMode string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
		addr, // host:port format
		user,
		pass,
		dbName,
		"",
		sslMode)

	// Reformat DSN, correctly separate host and port
	host, port := parseAddress(addr)
	dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Shanghai",
		host,
		port,
		user,
		pass,
		dbName,
		sslMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("[db connection failed] Database name: %s", dbName))
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	cfg := config.Conf.Postgres
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Second * cfg.ConnMaxLifeTime)

	if cfg.Migrate {
		err = db.AutoMigrate(
			&entity.User{},
			&entity.Post{},
			&entity.PostTag{},
			&entity.Category{},
			&entity.Tag{},
			&entity.Media{},
			&entity.Analytics{},
		)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

// parseAddress parses address string, separates host and port
func parseAddress(addr string) (host, port string) {
	if len(addr) == 0 {
		return "localhost", "5432"
	}

	for i := len(addr) - 1; i >= 0; i-- {
		if addr[i] == ':' {
			return addr[:i], addr[i+1:]
		}
	}

	return addr, "5432"
}

type Transaction interface {
	ExecTx(ctx context.Context, fn func(ctx context.Context) error) error
}

type contextTxKey struct{}

func (d *dbRepo) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.DbW.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func (d *dbRepo) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.DbW
}
