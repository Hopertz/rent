// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: houses.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createHouse = `-- name: CreateHouse :one
INSERT INTO house (location, block, partition, occupied) VALUES ($1,$2,$3,$4) RETURNING id
`

type CreateHouseParams struct {
	Location  string `json:"location"`
	Block     string `json:"block"`
	Partition int16  `json:"partition"`
	Occupied  bool   `json:"occupied"`
}

func (q *Queries) CreateHouse(ctx context.Context, arg CreateHouseParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createHouse,
		arg.Location,
		arg.Block,
		arg.Partition,
		arg.Occupied,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getHouseById = `-- name: GetHouseById :one
SELECT id,location, block, partition , Occupied, occupiedBy, version FROM house WHERE id = $1
`

func (q *Queries) GetHouseById(ctx context.Context, id uuid.UUID) (House, error) {
	row := q.db.QueryRowContext(ctx, getHouseById, id)
	var i House
	err := row.Scan(
		&i.ID,
		&i.Location,
		&i.Block,
		&i.Partition,
		&i.Occupied,
		&i.Occupiedby,
		&i.Version,
	)
	return i, err
}

const getHouseByIdWithTenant = `-- name: GetHouseByIdWithTenant :one
SELECT h.id,h.location, h.block, h.partition , h.Occupied, 
CONCAT(t.first_name || ' ' || t.last_name) AS tenant_name, t.id AS tenant_id, h.version 
FROM house h
Join tenant t ON h.occupiedBy = t.id
WHERE h.id = $1
`

type GetHouseByIdWithTenantRow struct {
	ID         uuid.UUID   `json:"id"`
	Location   string      `json:"location"`
	Block      string      `json:"block"`
	Partition  int16       `json:"partition"`
	Occupied   bool        `json:"occupied"`
	TenantName interface{} `json:"tenant_name"`
	TenantID   uuid.UUID   `json:"tenant_id"`
	Version    uuid.UUID   `json:"version"`
}

func (q *Queries) GetHouseByIdWithTenant(ctx context.Context, id uuid.UUID) (GetHouseByIdWithTenantRow, error) {
	row := q.db.QueryRowContext(ctx, getHouseByIdWithTenant, id)
	var i GetHouseByIdWithTenantRow
	err := row.Scan(
		&i.ID,
		&i.Location,
		&i.Block,
		&i.Partition,
		&i.Occupied,
		&i.TenantName,
		&i.TenantID,
		&i.Version,
	)
	return i, err
}

const getHouses = `-- name: GetHouses :many
SELECT id,location, block, partition , occupied FROM house
`

type GetHousesRow struct {
	ID        uuid.UUID `json:"id"`
	Location  string    `json:"location"`
	Block     string    `json:"block"`
	Partition int16     `json:"partition"`
	Occupied  bool      `json:"occupied"`
}

func (q *Queries) GetHouses(ctx context.Context) ([]GetHousesRow, error) {
	rows, err := q.db.QueryContext(ctx, getHouses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetHousesRow{}
	for rows.Next() {
		var i GetHousesRow
		if err := rows.Scan(
			&i.ID,
			&i.Location,
			&i.Block,
			&i.Partition,
			&i.Occupied,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateHouseById = `-- name: UpdateHouseById :exec
UPDATE house
SET location = $1, block = $2, partition = $3, occupied = $4, 
version = uuid_generate_v4(), occupiedBy = $5
WHERE id = $6 AND version = $7
`

type UpdateHouseByIdParams struct {
	Location   string        `json:"location"`
	Block      string        `json:"block"`
	Partition  int16         `json:"partition"`
	Occupied   bool          `json:"occupied"`
	Occupiedby uuid.NullUUID `json:"occupiedby"`
	ID         uuid.UUID     `json:"id"`
	Version    uuid.UUID     `json:"version"`
}

func (q *Queries) UpdateHouseById(ctx context.Context, arg UpdateHouseByIdParams) error {
	_, err := q.db.ExecContext(ctx, updateHouseById,
		arg.Location,
		arg.Block,
		arg.Partition,
		arg.Occupied,
		arg.Occupiedby,
		arg.ID,
		arg.Version,
	)
	return err
}
