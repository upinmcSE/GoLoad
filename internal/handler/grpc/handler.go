package grpc

import (
	"context"

	"github.com/upinmcSE/GoLoad/internal/generated/grpc/go_load"
	"github.com/upinmcSE/GoLoad/internal/logic"
)

type Handler struct {
	go_load.UnimplementedGoLoadServiceServer
	accountLogic logic.Account
}

func NewHandler(
	accountLogic logic.Account,
) go_load.GoLoadServiceServer {
	return &Handler{
		accountLogic: accountLogic,
	}
}

// CreateAccount implements go_load.GoLoadServiceServer.
func (a Handler) CreateAccount(ctx context.Context, request *go_load.CreateAccountRequest) (*go_load.CreateAccountResponse, error) {
	output, err := a.accountLogic.CreateAccount(context.Background(), logic.CreateAccountParams{
		AccountName: request.GetAccountName(),
		Password:    request.GetPassword(),
	})

	if err != nil {
		return nil, err
	}

	return &go_load.CreateAccountResponse{
		AccountId: output.ID,
	}, nil
}

// CreateDownloadTask implements go_load.GoLoadServiceServer.
func (a *Handler) CreateDownloadTask(ctx context.Context, request *go_load.CreateDownloadTaskRequest) (*go_load.CreateDownloadTaskResponse, error) {
	panic("unimplemented")
}

// CreateSession implements go_load.GoLoadServiceServer.
func (a *Handler) CreateSession(context.Context, *go_load.CreateSessionRequest) (*go_load.CreateSessionResponse, error) {
	panic("unimplemented")
}

// DeleteDownloadTask implements go_load.GoLoadServiceServer.
func (a *Handler) DeleteDownloadTask(context.Context, *go_load.DeleteDownloadTaskRequest) (*go_load.DeleteDownloadTaskResponse, error) {
	panic("unimplemented")
}

// GetDownloadTaskFile implements go_load.GoLoadServiceServer.
func (a *Handler) GetDownloadTaskFile(*go_load.GetDownloadTaskFileRequest, go_load.GoLoadService_GetDownloadTaskFileServer) error {
	panic("unimplemented")
}

// GetDownloadTaskList implements go_load.GoLoadServiceServer.
func (a *Handler) GetDownloadTaskList(context.Context, *go_load.GetDownloadTaskListRequest) (*go_load.GetDownloadTaskListResponse, error) {
	panic("unimplemented")
}

// UpdateDownloadTask implements go_load.GoLoadServiceServer.
func (a *Handler) UpdateDownloadTask(context.Context, *go_load.UpdateDownloadTaskRequest) (*go_load.UpdateDownloadTaskResponse, error) {
	panic("unimplemented")
}

// mustEmbedUnimplementedGoLoadServiceServer implements go_load.GoLoadServiceServer.
func (a *Handler) mustEmbedUnimplementedGoLoadServiceServer() {
	panic("unimplemented")
}

//protoc -I=. --go_out=internal/generated --go-grpc_out=internal/generated --grpc-gateway_out=internal/generated --grpc-gateway_opt=generate_unbound_methods=true --openapiv2_out=. --openapiv2_opt=generate_unbound_methods=true --validate_out=lang=go:internal/generated api/go_load.proto
