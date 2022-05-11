package grpc_client

import "AggreBot/internal/pkg/api"

func (c *Client) GetUserSources(userId int64) ([]*api.Source, error) {
	sourcesResponse, err := c.api.GetUserSources(
		c.ctx,
		&api.UserId{
			Id: userId,
		},
	)
	return sourcesResponse.Sources, err
}

func (c *Client) AddSource(userId int64, name, url string) error {
	_, err := c.api.AddSource(
		c.ctx,
		&api.AddSourceRequest{
			UserId: userId,
			Name:   name,
			Url:    url,
		},
	)
	return err
}

func (c *Client) UpdateSourceName(sourceId int64, name string) error {
	_, err := c.api.UpdateSourceName(
		c.ctx,
		&api.UpdateSourceNameRequest{
			Id:   sourceId,
			Name: name,
		},
	)
	return err
}

func (c *Client) UpdateSourceToggleActive(sourceId int64) (*api.SourceToggleActiveResponse, error) {
	toggleResponse, err := c.api.UpdateSourceToggleActive(
		c.ctx,
		&api.SourceId{
			Id: sourceId,
		},
	)
	return toggleResponse, err
}

func (c *Client) DeleteSource(sourceId int64) error {
	_, err := c.api.DeleteSource(
		c.ctx,
		&api.SourceId{
			Id: sourceId,
		},
	)
	return err
}