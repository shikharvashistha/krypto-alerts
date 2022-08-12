package handlers

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shikharvashistha/krypto-alerts/pkg/store"
	"github.com/shikharvashistha/krypto-alerts/pkg/utils"
)

func Test_appSvc_DeleteAlert(t *testing.T) {
	type fields struct {
		Store  store.Store
		logger *utils.Logger
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &appSvc{
				Store:  tt.fields.Store,
				Logger: tt.fields.logger,
			}
			svc.DeleteAlert(tt.args.c)
		})
	}
}

func Test_appSvc_CreateAlert(t *testing.T) {
	type fields struct {
		Store  store.Store
		logger *utils.Logger
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &appSvc{
				Store:  tt.fields.Store,
				Logger: tt.fields.logger,
			}
			svc.CreateAlert(tt.args.c)
		})
	}
}

func Test_appSvc_GetAlerts(t *testing.T) {
	type fields struct {
		Store  store.Store
		logger *utils.Logger
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &appSvc{
				Store:  tt.fields.Store,
				Logger: tt.fields.logger,
			}
			svc.GetAlerts(tt.args.c)
		})
	}
}

func Test_appSvc_GetToken(t *testing.T) {
	type fields struct {
		Store  store.Store
		logger *utils.Logger
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &appSvc{
				Store:  tt.fields.Store,
				Logger: tt.fields.logger,
			}
			svc.GetToken(tt.args.c)
		})
	}
}
