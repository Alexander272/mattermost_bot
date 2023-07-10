package services

import (
	"context"

	"github.com/Alexander272/mattermost_bot/internal/models"
	"github.com/mattermost/mattermost-server/v6/model"
)

type Most interface {
	Send(context.Context, models.Message) (*model.Response, error)
}

type Services struct {
	Most
}

type Deps struct {
	MostClient *model.Client4
	ChannelId  string
}

func NewServices(deps Deps) *Services {
	return &Services{
		Most: NewMostService(deps.MostClient, deps.ChannelId),
	}
}
