package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/v2/pkg/session"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStorage_GetStorageGroup(t *testing.T) {
	var result StorageGroup

	respData, err := loadTestData("TestStorage_GetStorageGroup.resp.json")
	if err != nil {
		t.Fatal(err)
	}

	if err := json.NewDecoder(bytes.NewBuffer(respData)).Decode(&result); err != nil {
		t.Fatal(err)
	}

	tests := map[string]struct {
		id               int
		responseStatus   int
		responseBody     string
		expectedPath     string
		expectedResponse *StorageGroup
		withError        error
	}{
		"200 OK": {
			id:               1,
			responseStatus:   http.StatusOK,
			responseBody:     string(respData),
			expectedPath:     "/storage/v1/storage-groups/1",
			expectedResponse: &result,
		},
		"400 Validation failure": {
			id:             -9,
			responseStatus: http.StatusBadRequest,
			responseBody: `{
				"type": "validation-error",
				"title": "Validation failure",
				"instance": "6d00fc96-5431-4efa-86eb-afbb6cbdb5bc",
				"status": 400,
				"detail": "Validation failed. Please review the errors.",
				"errors": [
					{
						"type": "error-types/invalid-value",
						"title": "Invalid value",
						"detail": "Unable to find the given storage group.",
						"field": "storageGroupId"
					}
				]
			}`,
			expectedPath: "/storage/v1/storage-groups/-9",
			withError: &Error{
				Type:       "validation-error",
				Title:      "Validation failure",
				Detail:     "Validation failed. Please review the errors.",
				StatusCode: http.StatusBadRequest,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			mockServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, test.expectedPath, r.URL.String())
				assert.Equal(t, http.MethodGet, r.Method)
				w.WriteHeader(test.responseStatus)
				_, err := w.Write([]byte(test.responseBody))
				assert.NoError(t, err)
			}))
			client := mockAPIClient(t, mockServer)
			ctx := session.ContextWithOptions(context.Background())
			result, err := client.GetStorageGroup(ctx, test.id)
			if test.withError != nil {
				assert.True(t, errors.Is(err, test.withError), "want: %s; got: %s", test.withError, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, test.expectedResponse, result)
		})
	}
}
