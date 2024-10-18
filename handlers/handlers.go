package handlers

import (
	"github.com/jackc/pgx/v5"
	"github.com/mattg1243/willow-server/db"
)

type Handler struct {
	conn      *pgx.Conn
	queries   *db.Queries
	validator *Validator
}

func New(conn *pgx.Conn) *Handler {
	v := NewValidator()
	return &Handler{conn, db.New(conn), v}
}

// Updates running event balances update
// Going to try to avoid using this for now...
// func (h *Handler) SyncEventBalances(ctx context.Context, clientId uuid.UUID) error {
// 	events, err := h.queries.GetEvents(ctx, clientId)
// 	if err != nil {
// 		fmt.Printf("%q", err)
// 		return err
// 	}

// 	sort.Slice(events, func(i, j int) bool {
// 		return events[i].Date.Time.Before(events[j].Date.Time)
// 	})

// 	clientBalance := 0
// 	for i := 0; i < len(events); i++ {
// 		clientBalance += int(events[i].Amount)
// 		events[i].Newbalance = int32(clientBalance)
// 		h.queries.UpdateEvent(ctx, db.UpdateEventParams{ ID: events[i].ID, Newbalance: int32(clientBalance)})
// 	}

// 	err = h.queries.UpdateClient(ctx, db.UpdateClientParams{
// 		Balance: int32(clientBalance),
// 	})

// 	if err != nil {
// 		fmt.Printf("%q", err)
// 		return err
// 	}
// }