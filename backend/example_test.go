package boschtail

import (
	"boschtail/ent"
	"boschtail/ent/order"
	"context"
	"log"
	"testing"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
)

func TestOrder(t *testing.T) {
	// Create an ent.Client with in-memory SQLite database.
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	order1, err := client.Order.Create().Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a todo: %v", err)
	}

	want := order.Status("ORDERED")
	if order1.Status != want || err != nil {
		t.Fatalf(`Want: %v --> return value: %v`, want, order1.Status)
	}
}
