package internal

import (
	"context"
	"github.com/jmichiels/cloud-functions-tests/internal/domain"
	"google.golang.org/api/iterator"
)

type clientFirestoreDocument struct {
	Id        string `firestore:"id"`
	FirstName string `firestore:"firstName"`
	LastName  string `firestore:"lastName"`
	Age       uint   `firestore:"age"`
}

const clientFirestoreCollectionPath = "clients"

// The methods below make firestoreDatabase implement clientRepository.

func (db *firestoreDatabase) storeClient(ctx context.Context, client *domain.Client) error {
	data := clientFirestoreDocument{
		Id:        client.Id.String(),
		FirstName: client.FirstName,
		LastName:  client.LastName,
		Age:       client.Age,
	}
	colRef := db.clt.Collection(clientFirestoreCollectionPath)
	docRef := colRef.Doc(data.Id)

	// Thanks to the fakeFirestoreTransaction, we can implement our repositories methods as if we were always in a
	// transaction without having to make duplicates (i.e. one for the transaction case, one for the normal case).

	return db.tx.Set(ctx, docRef, &data)
}

func (db *firestoreDatabase) getAllClients(ctx context.Context) ([]*domain.Client, error) {
	colRef := db.clt.Collection(clientFirestoreCollectionPath)
	iter := db.tx.Documents(ctx, colRef)
	var clients []*domain.Client
	for {
		snapshot, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var data clientFirestoreDocument
		if err := snapshot.DataTo(&data); err != nil {
			return nil, err
		}
		parsedId, err := domain.ParseUniqueIdFromString(data.Id)
		if err != nil {
			return nil, err
		}
		clients = append(clients, &domain.Client{
			Id:        parsedId,
			FirstName: data.FirstName,
			LastName:  data.LastName,
			Age:       data.Age,
		})
	}
	return clients, nil
}

func (db *firestoreDatabase) getClientById(ctx context.Context, id domain.UniqueId) (*domain.Client, error) {
	colRef := db.clt.Collection(clientFirestoreCollectionPath)
	docRef := colRef.Doc(id.String())
	snapshot, err := db.tx.Get(ctx, docRef)
	if err != nil {
		return nil, err
	}
	var data clientFirestoreDocument
	if err := snapshot.DataTo(&data); err != nil {
		return nil, err
	}
	parsedId, err := domain.ParseUniqueIdFromString(data.Id)
	if err != nil {
		return nil, err
	}
	return &domain.Client{
		Id:        parsedId,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Age:       data.Age,
	}, nil
}
