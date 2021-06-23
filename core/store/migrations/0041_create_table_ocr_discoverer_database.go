package migrations

import (
	"gorm.io/gorm"
)

const up41 = `
CREATE TABLE offchainreporting_discoverer_database (
	local_peer_id text NOT NULL REFERENCES encrypted_p2p_keys (peer_id) DEFERRABLE INITIALLY IMMEDIATE,
	remote_peer_id text NOT NULL,
	ann bytea NOT NULL,
	created_at timestamptz not null,
	updated_at timestamptz not null,
	PRIMARY KEY(local_peer_id, remote_peer_id)
);
`
const down41 = `
DROP TABLE offchainreporting_discoverer_database;
`

func init() {
	Migrations = append(Migrations, &Migration{
		ID: "0041_create_table_offchainreporting_discoverer_database",
		Migrate: func(db *gorm.DB) error {
			return db.Exec(up41).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Exec(down41).Error
		},
	})
}
