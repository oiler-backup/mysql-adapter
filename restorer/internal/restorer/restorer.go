// Package restorer contains entities to restore backup of MySQL Database.
package restorer

import (
	"context"
	"database/sql"
	"fmt"
	"os/exec"

	_ "github.com/go-sql-driver/mysql"
)

type Restorer struct {
	dbHost string
	dbPort string
	dbUser string
	dbPass string
	dbName string

	backupPath string
}

// NewRestorer is a constructor for Restorer.
// Accepts parameters to connect to database and backupPath where backup will be stored locally.
func NewRestorer(dbHost, dbPort, dbUser, dbPassword, dbName, backupPath string) Restorer {
	return Restorer{
		dbHost:     dbHost,
		dbPort:     dbPort,
		dbUser:     dbUser,
		dbPass:     dbPassword,
		dbName:     dbName,
		backupPath: backupPath,
	}
}

// Restore restores backup from local file.
// It uses mysql command with appropriate flags.
func (r Restorer) Restore(ctx context.Context) error {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", r.dbUser, r.dbPass, r.dbHost, r.dbPort, r.dbName)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return fmt.Errorf("failed to open driver for database: %v", err)
	}
	defer db.Close()

	err = db.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	connStr = fmt.Sprintf("mysql -h %s -P %s -u %s -p%s %s < %s",
		r.dbHost,
		r.dbPort,
		r.dbUser,
		r.dbPass,
		r.dbName,
		r.backupPath,
	)
	cmd := exec.Command("bash", "-c", connStr)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed executing mysql restore: %+v\n.Output:%s", err, string(output))
	}
	return nil
}
