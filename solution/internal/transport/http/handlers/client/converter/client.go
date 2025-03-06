package converter

import (
	"server/internal/transport/http/converter"
	req "server/internal/transport/http/handlers/client/model"

	"server/internal/model"
)

func FromReqToClient(client *req.Client) *model.Client {
	return &model.Client{
		ID:       converter.FromStringPtr(client.ClientID),
		Login:    converter.FromStringPtr(client.Login),
		Age:      converter.FromIntPtr(client.Age),
		Location: converter.FromStringPtr(client.Location),
		Gender:   converter.FromStringPtr(client.Gender),
	}
}

func FromReqToClients(clients []*req.Client) []*model.Client {
	result := make([]*model.Client, len(clients))
	for i, r := range clients {
		result[i] = FromReqToClient(r)
	}
	return result
}

func FromClientToResp(client *model.Client) *req.Client {
	return &req.Client{
		ClientID: &client.ID,
		Login:    &client.Login,
		Age:      &client.Age,
		Location: &client.Location,
		Gender:   &client.Gender,
	}
}

func FromClientsToResp(clients []*model.Client) []*req.Client {
	result := make([]*req.Client, len(clients))
	for i, r := range clients {
		result[i] = FromClientToResp(r)
	}
	return result
}
