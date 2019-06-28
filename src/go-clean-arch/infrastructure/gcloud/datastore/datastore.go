package datastore

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/pkg/errors"
)

type EnvVarEntity struct {
	AD_CHANNEL_API_BASE_URL  string
	APPLICATION_URL          string
	DAILY_REPORT_FROM        string
	DAILY_REPORT_TO          string
	DEVELOPMENT_BUILD        string
	FB_X_API_KEY             string
	INSTANCE_CONNECTION_NAME string
	SENDGRID_API_KEY         string
	SENDGRID_CC              string
	SENDGRID_FROM            string
	SENDGRID_TO              string
	SQL_DATABASE             string
	SQL_PASSWORD             string
	SQL_USER                 string
	STORAGE_BUCKET_NAME      string
}

type IDataStoreService interface {
	GetEnvVars() (*EnvVarEntity, error)
}

type DataStoreService struct {
	options *DataStoreServiceOptions
	client  *datastore.Client
	ctx     context.Context
}

type DataStoreServiceOptions struct {
	projectID string
}

func NewDataStoreService(options *DataStoreServiceOptions) (IDataStoreService, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, options.projectID)
	if err != nil {
		return nil, err
	}
	return &DataStoreService{options, client, ctx}, nil
}

func (svc *DataStoreService) GetEnvVars() (*EnvVarEntity, error) {
	var envVarEntities []EnvVarEntity
	kind := "envVars"
	q := datastore.NewQuery(kind)
	if _, err := svc.client.GetAll(svc.ctx, q, &envVarEntities); err != nil {
		return nil, errors.Wrap(err, "svc.client.GetAll(svc.ctx, q, &entities)")
	}

	return &envVarEntities[0], nil
}
