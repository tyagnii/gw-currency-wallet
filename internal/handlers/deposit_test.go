package handlers

import (
	"bytes"
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

func TestHandler_Deposit(t *testing.T) {
	controller := gomock.NewController(t)
	mockDB := mock_db.NewMockDBConnector(controller)
	mockClient := mock_exchanger.NewMockExchangeServiceClient(controller)
	sLogger, _ := logger.NewSugaredLogger()
	Cache, _ := cache.NewCache()

	rw := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rw)

	type args struct {
		hreq   models.DepositReq
		dbreq  models.Wallet
		reqUrl string
		c      *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",

			args: args{
				c: ctx,
				hreq: models.DepositReq{
					Currency: "USD",
					Amount:   1,
				},
				dbreq: models.Wallet{
					Balance: models.Currency{
						USD: 1,
						RUB: 0,
						EUR: 0,
					},
				},
				reqUrl: "/api/v1/wallet/deposit",
			},
		},
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

			b, _ := json.Marshal(tt.args.hreq)

			req, err := http.NewRequest(http.MethodPost, tt.args.reqUrl, bytes.NewBuffer(b))
			if err != nil {
				t.Fatal(err)
			}
			tt.args.c.Request = req

			mockDB.EXPECT().
				Deposit(tt.args.c, tt.args.dbreq).Return(nil).Times(1)

			h.Deposit(tt.args.c)
			assert.Equal(t, http.StatusOK, rw.Code)
		})
	}
}
