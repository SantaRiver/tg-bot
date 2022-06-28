package repository

import (
	"context"
	"homework-2/internal/app"
	pb "homework-2/pkg/api"
	"log"
)

func (r *repository) GetShare(ctx context.Context, shareReq *pb.Share) (shareRes *pb.Share, err error) {
	const query1 = `
		select id from shares where "figi" = $1
	`
	rows, err := r.pool.Query(ctx, query1, shareReq.Figi)
	if err != nil {
		err = app.NotFoundErr
		return
	}
	defer rows.Close()

	if !rows.Next() {
		err = app.NotFoundErr
		return
	}
	result, err := rows.Values()
	if err != nil {
		return
	}
	log.Println("GetShare:", result)
	return
}
