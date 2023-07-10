package services

import (
	"context"
	"fmt"

	"github.com/Alexander272/mattermost_bot/internal/models"
	"github.com/mattermost/mattermost-server/v6/model"
)

type MostService struct {
	MostClient *model.Client4
	ChannelId  string
}

func NewMostService(MostClient *model.Client4, ChannelId string) *MostService {
	return &MostService{
		MostClient: MostClient,
		ChannelId:  ChannelId,
	}
}

func (s *MostService) Send(ctx context.Context, message models.Message) (res *model.Response, err error) {
	attachment := &model.SlackAttachment{
		Fallback: message.Service.Id,
		Color:    "#f12e2e",
		// Title:    "Error SealurPro",
		Fields: []*model.SlackAttachmentField{
			{
				Title: "Ошибка",
				Value: fmt.Sprintf("#### %s", message.Data.Error),
			},
			// {
			// 	Title: "Дата",
			// 	Value: "07.07.2030 15:54:45",
			// },
			{
				Title: "IP",
				Short: true,
				Value: message.Data.IP,
			},
			{
				Title: "URL",
				Short: true,
				Value: message.Data.URL,
			},
			{
				Title: "Пользователь",
				Short: true,
				Value: message.Data.User,
			},
			{
				Title: "Компания",
				Short: true,
				Value: message.Data.Company,
			},
		},
	}

	post := &model.Post{
		ChannelId: s.ChannelId,
		// Message:   "### ```SealurPro```\n #pro_error\n ### Дата: 07.07.2030 15:54:45",
		Message: fmt.Sprintf("### ```%s```\n #%s_error\n #### Дата: %s", message.Service.Name, message.Service.Id, message.Data.Date),
	}

	post.AddProp("attachments", []*model.SlackAttachment{attachment})

	_, res, err = s.MostClient.CreatePost(post)
	if err != nil {
		return nil, fmt.Errorf("failed to send message. error: %w", err)
	}

	return res, nil
}
