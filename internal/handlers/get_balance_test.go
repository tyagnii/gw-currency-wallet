package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tyagnii/gw-currency-wallet/gen/mock/mock_db"
	"github.com/tyagnii/gw-currency-wallet/internal/db/models"
	"github.com/tyagnii/gw-currency-wallet/internal/logger"
	"github.com/tyagnii/gw-currency-wallet/pkg/cache"
	mock_exchanger "github.com/tyagnii/gw-proto/gen/mock"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_GetBalance(t *testing.T) {
	controller := gomock.NewController(t)
	mockDB := mock_db.NewMockDBConnector(controller)
	mockClient := mock_exchanger.NewMockExchangeServiceClient(controller)
	sLogger, _ := logger.NewSugaredLogger()
	Cache, _ := cache.NewCache()

	rw := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rw)

	type fields struct {
		dbCall *gomock.Call
		resp   models.Wallet
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test",
			fields: fields{
				dbCall: mockDB.EXPECT().
					GetBalance(ctx, models.User{Username: "user"}).
					Return(models.Wallet{
						Message: "abc",
						Balance: models.Currency{
							USD: 1,
							RUB: 2,
							EUR: 3,
						},
					},
						nil).Times(1),
				resp: models.Wallet{
					Message: "abc",
					Balance: models.Currency{
						USD: 1,
						RUB: 2,
						EUR: 3,
					},
				},
			},
			args: args{
				c: ctx,
			},
		},
		// TODO: Add test cases.
	}

	h := &Handler{
		dbconn:  mockDB,
		eClient: mockClient,
		sLogger: sLogger,
		Cache:   Cache,
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.args.c.AddParam("username", "user")
			// return nil error
			h.GetBalance(tt.args.c)

			// Marshall expected response
			jsonArgs, _ := json.Marshal(tt.fields.resp)
			assert.Equal(t, rw.Body.String(), string(jsonArgs))
			assert.Equal(t, rw.Code, http.StatusOK)
		})
	}
}
