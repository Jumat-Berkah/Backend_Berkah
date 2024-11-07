package helper

import (
	"Backend_berkah/model"
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(mconn model.DBInfo) (*mongo.Database, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mconn.DBString))
	if err != nil {
		mconn.DBString = SRVLookup(mconn.DBString)
		client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mconn.DBString))
		if err != nil {
			return nil, err
		}
	}
	db := client.Database(mconn.DBName)
	return db, nil
}

func SRVLookup(srvuri string) (mongouri string) {
	// Check if the URI is a valid SRV URI format
	atsplits := strings.Split(srvuri, "@")
	if len(atsplits) < 2 {
		fmt.Println("Invalid SRV URI format: missing '@' separator")
		return ""
	}

	userpass := strings.Split(atsplits[0], "//")
	if len(userpass) < 2 {
		fmt.Println("Invalid userpass format: missing '//' separator")
		return ""
	}

	// Construct the MongoDB URI with userpass information
	mongouri = "mongodb://" + userpass[1] + "@"
	slashsplits := strings.Split(atsplits[1], "/")
	if len(slashsplits) < 2 {
		fmt.Println("Invalid domain or database name format: missing '/' separator")
		return ""
	}

	// Parse domain and database name
	domain := slashsplits[0]
	dbname := slashsplits[1]

	fmt.Println("Parsed domain:", domain)
	fmt.Println("Parsed database name:", dbname)

	// Resolver setup
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, "8.8.8.8:53")
		},
	}

	// Perform SRV lookup
	_, srvs, err := r.LookupSRV(context.Background(), "mongodb", "tcp", domain)
	if err != nil {
		fmt.Println("Error in SRV Lookup:", err)
		return ""
	}

	// Build the SRV list
	var srvlist string
	for _, srv := range srvs {
		srvlist += strings.TrimSuffix(srv.Target, ".") + ":" + strconv.FormatUint(uint64(srv.Port), 10) + ","
	}

	// Fetch TXT records, if any
	txtrecords, txtErr := r.LookupTXT(context.Background(), domain)
	if txtErr != nil {
		fmt.Println("Error fetching TXT records:", txtErr)
	}

	var txtlist string
	for _, txt := range txtrecords {
		txtlist += txt
	}

	// Construct the final MongoDB URI
	mongouri = mongouri + strings.TrimSuffix(srvlist, ",") + "/" + dbname + "?ssl=true&" + txtlist
	fmt.Println("Constructed MongoDB URI:", mongouri)

	return mongouri
}



func GetRandomDoc[T any](db *mongo.Database, collection string, size uint) (result []T, err error) {
	filter := mongo.Pipeline{
		{{Key: "$sample", Value: bson.D{{Key: "size", Value: size}}}},
	}
	ctx := context.Background()
	cursor, err := db.Collection(collection).Aggregate(ctx, filter)
	if err != nil {
		return
	}

	err = cursor.All(ctx, &result)

	return
}

func GetAllDoc[T any](db *mongo.Database, collection string) (doc T, err error) {
	ctx := context.Background()
	cur, err := db.Collection(collection).Find(ctx, bson.M{})
	if err != nil {
		return
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, &doc)
	if err != nil {
		return
	}
	return
}

func GetOneDoc[T any](db *mongo.Database, collection string, filter bson.M) (doc T, err error) {
	err = db.Collection(collection).FindOne(context.Background(), filter).Decode(&doc)
	if err != nil {
		return
	}
	return
}

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}, err error) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		return
	}
	return insertResult.InsertedID, nil
}

// With replaceOne() you can only replace the entire document,
// while updateOne() allows for updating fields. Since replaceOne() replaces the entire document - fields in the old document not contained in the new will be lost.
// With updateOne() new fields can be added without losing the fields in the old document.
func UpdateDoc(db *mongo.Database, collection string, filter bson.M, updatefield bson.M) (updateresult *mongo.UpdateResult, err error) {
	updateresult, err = db.Collection(collection).UpdateOne(context.TODO(), filter, updatefield)
	if err != nil {
		return
	}
	return
}

func ReplaceOneDoc(db *mongo.Database, collection string, filter bson.M, doc interface{}) (updatereseult *mongo.UpdateResult, err error) {
	updatereseult, err = db.Collection(collection).ReplaceOne(context.TODO(), filter, doc)
	if err != nil {
		return
	}
	return
}