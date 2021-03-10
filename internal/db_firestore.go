package internal

import (
	"cloud.google.com/go/firestore"
	"context"
)

// Implementation of repositories based on firestore.
type firestoreDatabase struct {
	clt *firestore.Client

	// Will be an actualFirestoreTransaction if inside a transaction, or a fakeFirestoreTransaction otherwise.
	tx firestoreTransaction
}

func newFirestoreDatabase(clt *firestore.Client) *firestoreDatabase {
	return &firestoreDatabase{
		clt: clt,
		tx: &fakeFirestoreTransaction{
			clt: clt,
		},
	}
}

func (db *firestoreDatabase) runTransaction(ctx context.Context, f func(ctx context.Context, tx repositories) error) error {
	return db.clt.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		return f(ctx, &firestoreDatabase{
			clt: db.clt,
			tx: &actualFirestoreTransaction{
				tx: tx,
			},
		})
	})
}

// Implemented by actualFirestoreTransaction and fakeFirestoreTransaction.
type firestoreTransaction interface {
	//todo Create(ctx context.Context,dr *firestore.DocumentRef, data interface{}) error
	//todo Delete(ctx context.Context,dr *firestore.DocumentRef, opts ...firestore.Precondition) error
	//todo DocumentRefs(ctx context.Context,cr *firestore.CollectionRef) *firestore.DocumentRefIterator
	Documents(ctx context.Context, q firestore.Queryer) *firestore.DocumentIterator
	Get(ctx context.Context, dr *firestore.DocumentRef) (*firestore.DocumentSnapshot, error)
	//todo GetAll(ctx context.Context,drs []*firestore.DocumentRef) ([]*firestore.DocumentSnapshot, error)
	Set(ctx context.Context, dr *firestore.DocumentRef, data interface{}, opts ...firestore.SetOption) error
	//todo Update(ctx context.Context,dr *firestore.DocumentRef, data []firestore.Update, opts ...firestore.Precondition) error
}

// Implements firestoreTransaction.
type fakeFirestoreTransaction struct {
	clt *firestore.Client
}

func (fakeTx *fakeFirestoreTransaction) Documents(ctx context.Context, q firestore.Queryer) *firestore.DocumentIterator {
	if colRef, ok := q.(firestore.CollectionRef); ok {
		return colRef.Documents(ctx)
	}
	return q.(firestore.Query).Documents(ctx)
}

func (fakeTx *fakeFirestoreTransaction) Get(ctx context.Context, dr *firestore.DocumentRef) (*firestore.DocumentSnapshot, error) {
	return dr.Get(ctx)
}

func (fakeTx *fakeFirestoreTransaction) Set(ctx context.Context, dr *firestore.DocumentRef, data interface{}, opts ...firestore.SetOption) error {
	_, err := dr.Set(ctx, data, opts...)
	return err
}

// Implements firestoreTransaction.
type actualFirestoreTransaction struct {
	tx *firestore.Transaction
}

func (actualTx *actualFirestoreTransaction) Documents(ctx context.Context, q firestore.Queryer) *firestore.DocumentIterator {
	return actualTx.tx.Documents(q)
}

func (actualTx *actualFirestoreTransaction) Get(ctx context.Context, dr *firestore.DocumentRef) (*firestore.DocumentSnapshot, error) {
	return actualTx.tx.Get(dr)
}

func (actualTx *actualFirestoreTransaction) Set(ctx context.Context, dr *firestore.DocumentRef, data interface{}, opts ...firestore.SetOption) error {
	return actualTx.tx.Set(dr, data, opts...)
}
