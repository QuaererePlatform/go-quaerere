package arangodb

import (
	"context"
	"log"

	"github.com/arangodb/go-driver"

	"github.com/QuaererePlatform/go-quaerere/internal/data_types/web_sites"
)

const WEB_SITE_COLLECTION = "WebSites"

func (s ArangoDBStorage) CreateWebSite(wp *web_sites.WebSite) (*driver.DocumentMeta, error) {
	ctx := context.Background()
	coll, _ := s.getCollection(ctx, WEB_SITE_COLLECTION)
	meta, err := coll.CreateDocument(ctx, wp)
	if err != nil {
		return nil, err
	}
	return &meta, nil
}

func (s ArangoDBStorage) ReadWebSite(key string) (*web_sites.WebSite, error) {
	ctx := context.Background()
	coll, _ := s.getCollection(ctx, WEB_SITE_COLLECTION)
	wp := new(web_sites.WebSite)
	meta, err := coll.ReadDocument(ctx, key, wp)
	log.Printf("Meta: %+v", meta)
	if err != nil {
		return nil, err
	}
	return wp, nil
}

func (s ArangoDBStorage) UpdateWebSite(key string, data map[string]interface{}) (*driver.DocumentMeta, error) {
	ctx := context.Background()
	coll, _ := s.getCollection(ctx, WEB_SITE_COLLECTION)
	meta, err := coll.UpdateDocument(ctx, key, data)
	if err != nil {
		return nil, err
	}
	return &meta, nil
}

func (s ArangoDBStorage) DeleteWebSite(key string) (*driver.DocumentMeta, error) {
	ctx := context.Background()
	coll, _ := s.getCollection(ctx, WEB_SITE_COLLECTION)
	meta, err := coll.RemoveDocument(ctx, key)
	if err != nil {
		return nil, err
	}
	return &meta, nil
}
