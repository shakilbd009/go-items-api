package elasticsearch

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/olivere/elastic"
	"github.com/shakilbd009/go-utils-lib/logger"
)

const (
	esHostIP = "es_hostIP"
)

var (
	esHost                   = os.Getenv(esHostIP)
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string, string) (*elastic.GetResult, error)
	Search(string, elastic.Query) (*elastic.SearchResult, error)
	Delete(index, docType, id string) (string, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	log := logger.GetLogger()
	client, err := elastic.NewClient(
		elastic.SetURL(esHost),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	Client.setClient(client)
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(index, docType string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().
		Index(index).
		Type(docType).
		BodyJson(doc).
		Do(ctx)
	if err != nil {
		logger.Error(
			fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Get(index, docType, id string) (*elastic.GetResult, error) {

	ctx := context.Background()
	result, err := c.client.Get().Index(index).Type(docType).Id(id).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to get index document by ID %s", id), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Search(index string, query elastic.Query) (*elastic.SearchResult, error) {

	ctx := context.Background()
	result, err := c.client.Search(index).Query(query).RestTotalHitsAsInt(true).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to search documents with index %s", index), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Delete(index, docType, id string) (string, error) {
	ctx := context.Background()
	result, err := c.client.Delete().Index(index).Type(docType).Id(id).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to delete document by ID %s", id), err)
		return "", err
	}
	return result.Result, nil
}
