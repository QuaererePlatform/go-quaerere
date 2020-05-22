package arangodb

import (
	"context"
	"log"

	"github.com/arangodb/go-driver"
	"github.com/jinzhu/copier"

	"github.com/QuaererePlatform/go-quaerere/internal/common/web_pages"
)

const WEB_PAGE_COLLECTION = "WebPages"

type (
	WebPageStore struct {
		driver.Collection
	}
)

func (s WebPageStore) Create(wp *web_pages.WebPage) (*DocumentMeta, error) {
	ctx := context.Background()
	meta, err := s.CreateDocument(ctx, wp)
	if err != nil {
		return nil, err
	}

	m := new(DocumentMeta)
	err = copier.Copy(&meta, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s WebPageStore) Read(key string) (*web_pages.WebPage, error) {
	ctx := context.Background()
	coll, err := s.getCollection(ctx, WEB_PAGE_COLLECTION)
	if err != nil {
		return nil, err
	}
	wp := new(web_pages.WebPage)
	meta, err := coll.ReadDocument(ctx, key, wp)
	log.Printf("Meta: %+v", meta)
	if err != nil {
		return nil, err
	}
	return wp, nil
}

func (s WebPageStore) Update(key string, data map[string]interface{}) (*DocumentMeta, error) {
	ctx := context.Background()
	coll, _ := s.getCollection(ctx, WEB_PAGE_COLLECTION)
	meta, err := coll.UpdateDocument(ctx, key, data)
	if err != nil {
		return nil, err
	}

	m := new(DocumentMeta)
	err = copier.Copy(&meta, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s WebPageStore) Delete(key string) (*DocumentMeta, error) {
	ctx := context.Background()
	coll, _ := s.getCollection(ctx, WEB_PAGE_COLLECTION)
	meta, err := coll.RemoveDocument(ctx, key)
	if err != nil {
		return nil, err
	}

	m := new(DocumentMeta)
	err = copier.Copy(&meta, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s ArangoDBStorage) CreateWebPage(wp *web_pages.WebPage) (*driver.DocumentMeta, error) {
	log.Printf("arangodb.CreateWebPage() before getCollection s: %+v", s)
	ctx := context.Background()
	coll, err := s.getCollection(ctx, WEB_PAGE_COLLECTION)
	log.Printf("arangodb.CreateWebPage() coll: %+v", coll)
	log.Printf("arangodb.CreateWebPage() after getCollections s: %+v", s)
	if err != nil {
		return nil, err
	}
	meta, err := coll.CreateDocument(ctx, wp)
	if err != nil {
		return nil, err
	}
	return &meta, nil
}

func (s ArangoDBStorage) ReadWebPage(key string) (*web_pages.WebPage, error) {
	ctx := context.Background()
	coll, err := s.getCollection(ctx, WEB_PAGE_COLLECTION)
	if err != nil {
		return nil, err
	}
	wp := new(web_pages.WebPage)
	meta, err := coll.ReadDocument(ctx, key, wp)
	log.Printf("Meta: %+v", meta)
	if err != nil {
		return nil, err
	}
	return wp, nil
}

func (s ArangoDBStorage) UpdateWebPage(key string, data map[string]interface{}) (*driver.DocumentMeta, error) {
	ctx := context.Background()
	coll, _ := s.getCollection(ctx, WEB_PAGE_COLLECTION)
	meta, err := coll.UpdateDocument(ctx, key, data)
	if err != nil {
		return nil, err
	}
	return &meta, nil
}

func (s ArangoDBStorage) DeleteWebPage(key string) (*driver.DocumentMeta, error) {
	ctx := context.Background()
	coll, _ := s.getCollection(ctx, WEB_PAGE_COLLECTION)
	meta, err := coll.RemoveDocument(ctx, key)
	if err != nil {
		return nil, err
	}
	return &meta, nil
}
