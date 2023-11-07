package bigquery

import (
	"context"
	"errors"

	bigquery "cloud.google.com/go/bigquery"
	option "google.golang.org/api/option"

	"github.com/gasprawira/tokopedia-library/helper"
)

func CreateNewClient(projectID string, credentialJSON []byte) (client *bigquery.Client, err error) {

	ctx := context.Background()

	if helper.GetDataCenterName() == helper.DC_GCP {
		client, err = bigquery.NewClient(ctx, projectID)
	} else {
		client, err = bigquery.NewClient(ctx, projectID, option.WithCredentialsJSON(credentialJSON))
	}

	if err != nil {
		return nil, err
	}
	if client == nil {
		err = errors.New("BigQuery instance is nil")
	}
	return
}
