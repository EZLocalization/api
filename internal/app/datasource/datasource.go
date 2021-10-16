// package datasource

// import (
// 	"context"
// 	"time"
// )

// type Configuration struct {
// 	PostgresConfig *postgres.Configuration
// }

// type DataSource struct {
// 	postgres *postgres.Postgres
// }

// // New datasource
// func New(ctx context.Context, cfg *Configuration) (*DataSource, error) {

// 	// create postgres instance
// 	postgres, err := postgres.New(ctx, cfg.PostgresConfig)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Database Schema Setup
// 	postgres.AutoMigrate(
// 	// &models.User{},
// 	)

// 	return &DataSource{
// 		postgres: postgres,
// 	}, nil
// }

// // Terminate datasource dataset terminate
// func (d *DataSource) Terminate(ctx context.Context) {

// 	// logger.Debug("terminateing datasource...")

// 	// terminate postgres
// 	if d.postgres == nil {
// 		d.postgres.Close(ctx)
// 	}

// 	// logger.Debug("terminated datasource")
// }

// // GetPostgres get postgres instance
// func (d *DataSource) GetPostgres() *postgres.Postgres {
// 	timeoutContext, _ := context.WithTimeout(context.Background(), 5*time.Second)
// 	d.postgres.DB = d.postgres.WithContext(timeoutContext)
// 	return d.postgres
// }
