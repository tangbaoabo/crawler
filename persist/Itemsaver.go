package persist

import (
	"crawler/engine"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v6"
	"log"
)

func ItemSaver(index string) (item chan engine.Item, err error) {
	out := make(chan engine.Item)
	client, e := elastic.NewClient(elastic.SetSniff(false))
	if e != nil {
		return nil, e

	}
	go func() {
		count := 0
		for {
			item := <-out
			log.Printf("Item Save got Item #%d %v", count, item)
			count++
			err := save(item, client, index)
			if err != nil {
				log.Printf("持久化到es出错啦，%v", err)
				continue
			}
		}
	}()
	return out, nil
}

func save(item engine.Item, client *elastic.Client, index string) (err error) {
	if item.Type == "" {
		return errors.New("must apply Type")
	}
	indexService := client.Index().
		Index(index).
		Type(item.Type).
		Id(item.Id).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err = indexService.Do(context.Background())
	if err != nil {
		panic(err)
		return err
	}
	return nil

}
