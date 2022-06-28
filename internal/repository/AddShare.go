package repository

import (
	"context"
	"fmt"
	"homework-2/internal/app"
	pb "homework-2/pkg/api"
)

func (r *repository) AddShare(ctx context.Context, shareReq *pb.Share) error {
	fmt.Println("Repository: AddShare")
	const query1 = `
		select id from shares where "figi" = $1
	`
	const query2 = `
		insert into shares
		("figi", "class_code", "currency", "name", "country_of_risk") 
		values ($1, $2, $3, $4, $5);
	`
	rows, err := r.pool.Query(ctx, query1, shareReq.Figi)
	if err != nil {
		return err
	}
	if rows.Next() {
		return app.AlreadyExistErr
	}
	defer rows.Close()
	_, err = r.pool.Exec(ctx, query2, shareReq.Figi, shareReq.ClassCode, shareReq.Currency, shareReq.Name, shareReq.CountryOfRisk)
	return nil

}
