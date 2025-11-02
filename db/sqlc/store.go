package sqlc

type Store interface {
	Querier
}

/*
// единая точка для всех операций с БД.
type SQLStore struct {
	*Queries
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *SQLStore {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}*/

/*
// обёртка для выполнения операций в транзакциях.
func (s *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	// 1. НАЧАТЬ транзакцию
	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	// 2. Создать Queries для этой транзакции
	q := New(tx)
	// 3. ВЫПОЛНИТЬ твою функцию
	err = fn(q)
	if err != nil {
		// 4. ОТКАТИТЬ при ошибке
		tx.Rollback(ctx)
		return err
	}
	// 5. ПОДТВЕРДИТЬ если всё ок
	return tx.Commit(ctx)
}
*/
/*
func (store *Store) CreateOrderTx(ctx context.Context, userID uuid.UUID, item []CartItem) error {
	return store.execTx(ctx, func(q *Queries) error {
		order, err := q.CreateOrderTx(ctx, userID)
	})
}*/
