package cosmosdb

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PercentileTargetListMetricsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *PercentileMetricListResult
}

type PercentileTargetListMetricsOperationOptions struct {
	Filter *string
}

func DefaultPercentileTargetListMetricsOperationOptions() PercentileTargetListMetricsOperationOptions {
	return PercentileTargetListMetricsOperationOptions{}
}

func (o PercentileTargetListMetricsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o PercentileTargetListMetricsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o PercentileTargetListMetricsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

// PercentileTargetListMetrics ...
func (c CosmosDBClient) PercentileTargetListMetrics(ctx context.Context, id TargetRegionId, options PercentileTargetListMetricsOperationOptions) (result PercentileTargetListMetricsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/percentile/metrics", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var model PercentileMetricListResult
	result.Model = &model

	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
