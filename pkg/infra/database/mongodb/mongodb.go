package mongodb

import (
	context "context"
	atomic "sync/atomic"
	time "time"

	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"

	sys "go-rest-api-template/pkg/infra/system"
)

var (
	db      *mongo.Client
	healthy int32
)

// Connect - Connect at Mongodb
func Connect() {
	db = createClient()
	go monitor()
}

// IsAlive return Mongodb session status
func IsAlive() bool {
	return atomic.LoadInt32(&healthy) == 1
}

// RetrieveDatabase Return a Mongodb session
func RetrieveDatabase() *mongo.Database {
	return db.Database("gorest_adm")
}

// Disconnect - Disconnect at Mongodb
func Disconnect() {
	if db != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		db.Disconnect(ctx)
		sys.Info("[Mongodb session closed]")
	}
}

func createClient() *mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI(sys.Properties.Mongodb))

	if err != nil {
		sys.Error("[Could not create Mongodb client] err:%+v", err)
		return nil
	}

	err = client.Connect(context.TODO())
	if err != nil {
		sys.Error("[Could not connect at Mongodb] err:%+v", err)
		return nil
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		setStatusDown()
		sys.Warn("[Could create a Mongodb session] err:%+v", err)
	} else {
		setStatusUp()
		dataBases, _ := client.ListDatabases(context.TODO(), nil, nil)
		sys.Info("[Mongodb session created with databases: %+v]", dataBases)
	}

	return client
}

func monitor() {
	for {

		time.Sleep(30 * time.Second)

		if db == nil || db.Ping(context.TODO(), nil) != nil {
			setStatusDown()
			sys.Warn("[Mongodb session is not active, trying to reconnect]")
		} else {
			setStatusUp()
			sys.Info("[Mongodb session it's alive]")
		}

	}
}

func setStatusUp() {
	atomic.StoreInt32(&healthy, 1)
}

func setStatusDown() {
	atomic.StoreInt32(&healthy, 0)
}
