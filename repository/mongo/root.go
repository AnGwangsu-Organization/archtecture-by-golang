package mongo

import (
	"context"
	"eCommerce/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"strings"
)

type Mongo struct {
	config *config.Config

	client *mongo.Client
	db     *mongo.Database

	// TODO : 사용할 컬렉션
	user    *mongo.Collection
	content *mongo.Collection
	history *mongo.Collection
}

func NewMongo(config *config.Config) (*Mongo, error) {
	m := &Mongo{
		config: config,
	}

	ctx := context.Background()
	var err error
	if m.client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.Mongo.Uri)); err != nil {
		panic(err)
	} else if err = m.client.Ping(ctx, nil); err != nil {
		panic(err)
	} else {
		// * DB 설정
		m.db = m.client.Database(config.Mongo.Db)

		m.user = m.db.Collection("user")
		m.content = m.db.Collection("content")
		m.history = m.db.Collection("history")

		if err = createIndex(m.user, []string{"user"}, []string{"user"}); err != nil {
			panic(err)
		} else if err = createIndex(m.content, []string{"name"}, []string{"name"}); err != nil {
			panic(err)
		} else if err = createIndex(m.history, []string{"user"}, []string{}); err != nil {
			panic(err)
		}
	}

	return m, nil
}

// * 유틸 함수
func createIndex(collection *mongo.Collection, indexes, uniques []string) error {
	// 없는 경우에만 생성을 진행하는 코드

	// 인덱스 구조체
	type indexOptions struct {
		key    string
		order  int64
		unique bool
	}

	var indexsOpt []indexOptions

	for _, field := range indexes {
		noU := false
		// unique한지 체크
		for _, unique := range uniques {
			if unique == field {
				indexsOpt = append(indexsOpt, indexOptions{key: field, order: -1, unique: true})
				noU = true
				break
			}
		}

		// 같은 인덱스가 없으면 false
		if noU {
			indexsOpt = append(indexsOpt, indexOptions{key: field, order: -1, unique: false})
		}
	}

	ctx := context.Background()

	needToCreate := make(map[string]indexOptions)

	if indexCursor, err := collection.Indexes().List(ctx); err != nil {
		panic(err)
	} else {
		defer indexCursor.Close(ctx)

		for indexCursor.Next(ctx) {
			if v, ok := indexCursor.Current.Lookup("name").StringValueOK(); !ok || v == "_id_" {
				continue
			} else {

				split := strings.Split(v, "_")

				if len(split) == 2 {
					if order, err := strconv.Atoi(split[1]); err == nil {
						if order == 1 || order == -1 {
							needToCreate[split[0]] = indexOptions{split[0], int64(order), false}
						}
					}
				}
			}

		}
	}

	for _, i := range indexsOpt {
		if value, ok := needToCreate[i.key]; ok {
			opt := options.Index()

			if value.unique {
				opt.SetUnique(value.unique)
			}

			m := mongo.IndexModel{
				Keys:    bson.D{{Key: value.key, Value: 1}},
				Options: opt,
			}

			if _, err := collection.Indexes().CreateOne(ctx, m); err != nil {
				return err
			}
		}
	}

	return nil
}
