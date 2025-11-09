package resources

import (
	"gold-gym-be/internal/service/goldgym" // sesuaikan path

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"

	"go.uber.org/zap"
)

type BootResources struct {
	DBLocal      *sqlx.DB
	DBProd       *sqlx.DB
	Redis        *redis.Client
	GoldSvcLocal goldgym.Service
	GoldSvcProd  goldgym.Service
	Tracer       opentracing.Tracer
	Logger       *zap.Logger
}
