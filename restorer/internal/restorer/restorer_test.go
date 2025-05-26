package restorer

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	backuper "github.com/oiler-backup/mysql-adapter/backuper/backuper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	ctx        = context.Background()
	dbUser     = "testuser"
	dbPass     = "testpassword"
	dbName     = "testdb"
	backupName = "backup.dump"
)

func setupMySQLContainer() (*tc.Container, error) {
	req := tc.ContainerRequest{
		Image:           "mysql:8.0",
		ExposedPorts:    []string{"3306/tcp"},
		AlwaysPullImage: false,
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "rootpassword",
			"MYSQL_USER":          dbUser,
			"MYSQL_PASSWORD":      dbPass,
			"MYSQL_DATABASE":      dbName,
		},
		WaitingFor: wait.ForListeningPort("3306/tcp"),
	}

	mysqlC, err := tc.GenericContainer(ctx, tc.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	return &mysqlC, err
}

func Test_Redtore_UploadValidDump(t *testing.T) {
	mysqlC, err := setupMySQLContainer()
	require.NoError(t, err)
	defer func() {
		err := (*mysqlC).Terminate(ctx)
		if err != nil {
			panic(err)
		}
	}()

	dbhost, _ := (*mysqlC).ContainerIP(ctx)
	dbPort, _ := (*mysqlC).MappedPort(ctx, "3306")
	tempDir := t.TempDir()
	backupFile := filepath.Join(tempDir, backupName)

	b := backuper.NewBackuper(
		dbhost,
		dbPort.Port(),
		dbUser,
		dbPass,
		dbName,
		backupFile,
	)

	err = b.Backup(ctx, false)

	r := NewRestorer(
		dbhost,
		dbPort.Port(),
		dbUser,
		dbPass,
		dbName,
		backupFile,
	)

	err = r.Restore(ctx)
	require.NoError(t, err)

	fileInfo, err := os.Stat(backupFile)
	require.NoError(t, err)
	assert.Greater(t, fileInfo.Size(), int64(0))
}

func Test_Redtore_InvalidDump(t *testing.T) {
	mysqlC, err := setupMySQLContainer()
	require.NoError(t, err)
	defer func() {
		err := (*mysqlC).Terminate(ctx)
		if err != nil {
			panic(err)
		}
	}()

	dbhost, _ := (*mysqlC).ContainerIP(ctx)
	tempDir := t.TempDir()
	backupFile := filepath.Join(tempDir, backupName)

	r := NewRestorer(
		dbhost,
		"3306",
		dbUser,
		dbPass,
		dbName,
		backupFile,
	)

	err = r.Restore(ctx)
	require.ErrorContains(t, err, "failed executing mysql restore:")
}

func Test_Redtore_InvalidDBHost(t *testing.T) {
	dbhost := "wrong"
	dbPort := "3306"
	r := NewRestorer(
		dbhost,
		dbPort,
		dbUser,
		dbPass,
		dbName,
		backupName,
	)

	err := r.Restore(ctx)
	require.ErrorContains(t, err, "failed to connect to database:")
}
