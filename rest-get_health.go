package sdk

/*
   Copyright 2016 Alexander I.Grafov <grafov@gmail.com>
   Copyright 2016-2019 The Grafana SDK authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

	   http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

   ॐ तारे तुत्तारे तुरे स्व
*/

import (
	"context"
	"encoding/json"
)

// HealthResponse represents the health of grafana server
type HealthResponse struct {
	Commit   string `json:"commit"`
	Database string `json:"database"`
	Version  string `json:"version"`
}

// GetHealth retrieves the health of the grafana server
// Reflects GET BaseURL API call.
func (r *Client) GetHealth(ctx context.Context) (HealthResponse, error) {
	var (
		raw []byte
		err error
	)
	if raw, _, err = r.get(ctx, "/api/health"); err != nil {
		return HealthResponse{}, err
	}

	health := HealthResponse{}
	if err := json.Unmarshal(raw, &health); err != nil {
		return HealthResponse{}, err
	}
	return health, nil
}
