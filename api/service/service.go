package service

import (
	"github.com/the-web3/event-watcher/api/models"
	"github.com/the-web3/event-watcher/database/worker"
	"strconv"
)

type Service interface {
	GetDepositTokensList(*models.QueryDTParams) (*models.DepositTokensResponse, error)

	QueryDTListParams(page string, pageSize string, order string) (*models.QueryDTParams, error)
}

type HandlerSvc struct {
	v                 *Validator
	depositTokensView worker.DepositTokensView
}

func (h HandlerSvc) GetDepositTokensList(params *models.QueryDTParams) (*models.DepositTokensResponse, error) {
	panic("implement me")
}

func (h HandlerSvc) QueryDTListParams(page string, pageSize string, order string) (*models.QueryDTParams, error) {
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}
	pageVal := h.v.ValidatePage(pageInt)

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		return nil, err
	}
	pageSizeVal := h.v.ValidatePageSize(pageSizeInt)
	orderBy := h.v.ValidateOrder(order)

	return &models.QueryDTParams{
		Page:     pageVal,
		PageSize: pageSizeVal,
		Order:    orderBy,
	}, nil
}

func New(v *Validator, dtv worker.DepositTokensView) Service {
	return &HandlerSvc{
		v:                 v,
		depositTokensView: dtv,
	}
}

func (h HandlerSvc) GetDepositList(params *models.QueryDTParams) (*models.DepositTokensResponse, error) {
	depositList, total := h.depositTokensView.QueryDepositTokensList(params.Page, params.PageSize, params.Order)
	return &models.DepositTokensResponse{
		Current: params.Page,
		Size:    params.PageSize,
		Total:   int64(total),
		Result:  depositList,
	}, nil
}
