package infrastructure

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// Use a service account
func FirestoreInit(ctx context.Context) (*firestore.Client, error) {

	sa := option.WithCredentialsFile("secrets/ServiceAccount.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Println("APP")
		log.Fatalln(err)
	}
	//log.Println(sa)

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Println("Firestore init error")
		log.Fatalln(err)
		return nil, err
	}
	//   defer client.Close()
	return client, nil
}
