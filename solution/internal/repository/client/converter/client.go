package converter

import (
	"server/internal/model"
	repo "server/internal/repository/client/model"
)

func ToClientFromRepo(client *repo.Client) *model.Client {
	return &model.Client{
		ID:       client.ClientID,
		Login:    client.Login,
		Age:      client.Age,
		Location: client.Location,
		Gender:   client.Gender,
	}
}

func ToClientsFromRepo(clients []*repo.Client) []*model.Client {
	models := make([]*model.Client, len(clients))
	for i, client := range clients {
		models[i] = ToClientFromRepo(client)
	}
	return models
}

func ToRepoFromClient(client *model.Client) *repo.Client {
	return &repo.Client{
		ClientID: client.ID,
		Login:    client.Login,
		Age:      client.Age,
		Location: client.Location,
		Gender:   client.Gender,
	}
}

func ToReposFromClients(clients []*model.Client) []*repo.Client {
	repos := make([]*repo.Client, len(clients))
	for i, client := range clients {
		repos[i] = ToRepoFromClient(client)
	}
	return repos
}
