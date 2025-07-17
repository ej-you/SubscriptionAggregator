package pg

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"SubscriptionAggregator/internal/app/entity"
	"SubscriptionAggregator/internal/app/repo"
	"SubscriptionAggregator/internal/pkg/database"
)

const (
	_dbDSN = "user=aggregator password=p@SSw0rd host=127.0.0.1 port=5432 dbname=aggregator_db sslmode=disable connect_timeout=10"
)

var (
	_repo repo.SubsRepoDB

	_subsUUID = uuid.New().String()
	_userUUID = "60601fee-2bf1-4721-ae6f-7636e79a0cba"
)

func TestMain(m *testing.M) {
	// open DB connection
	dbStorage, err := database.New(_dbDSN,
		database.WithTranslateError(),
		database.WithIgnoreNotFound(),
	)
	if err != nil {
		log.Fatalf("get db connection: %v", err)
	}
	_repo = NewSubsRepoDB(dbStorage)
	// run tests
	os.Exit(m.Run())
}

func TestSubs_Create(t *testing.T) {
	t.Log("Create new subs")

	newSubs := entity.Subscription{
		ID:          _subsUUID,
		ServiceName: "Yandex Plus",
		Price:       400,
		UserID:      "60601fee-2bf1-4721-ae6f-7636e79a0cba",
		StartDate:   time.Now().UTC(),
	}

	err := _repo.Create(&newSubs)
	require.NoError(t, err)

	_subsUUID = newSubs.ID
	t.Logf("New subs: %+v", newSubs)
}

func TestSubs_GetByID(t *testing.T) {
	t.Log("Get subs by ID")

	subs, err := _repo.GetByID(_subsUUID)
	require.NoError(t, err)

	t.Logf("Subscription: %+v", subs)
}

func TestSubs_GetList(t *testing.T) {
	t.Log("Get all subs")

	subsList, err := _repo.GetList()
	require.NoError(t, err)

	t.Logf("All subs: %v", subsList)
}

func TestSubs_Update(t *testing.T) {
	t.Log("Update book by ID")

	updatedSubs := entity.Subscription{
		ID:          _subsUUID,
		ServiceName: "Kion",
		Price:       500,
		UserID:      "60601fee-2bf1-4721-ae6f-7636e79a0cba",
		StartDate:   time.Now().UTC(),
	}

	err := _repo.Update(&updatedSubs)
	require.NoError(t, err)

	t.Logf("Updated subs: %+v", updatedSubs)
}

func TestSubs_Remove(t *testing.T) {
	t.Log("Remove subs by ID")

	err := _repo.Delete(_subsUUID)
	require.NoError(t, err)

	t.Logf("Subs with ID %s was deleted successfully", _subsUUID)
}
