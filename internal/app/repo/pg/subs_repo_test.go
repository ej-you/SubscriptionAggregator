package pg

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"SubscriptionAggregator/internal/app/entity"
	"SubscriptionAggregator/internal/app/errors"
	"SubscriptionAggregator/internal/app/repo"
	"SubscriptionAggregator/internal/pkg/database"
)

const (
	_dbDSN = "user=aggregator password=p@SSw0rd host=127.0.0.1 port=5432 dbname=aggregator_db sslmode=disable connect_timeout=10"
)

var (
	_repo repo.SubsRepoDB

	_subsUUID = uuid.NewString()
	_userUUID = "44601fee-2bf1-4721-ae6f-7636e79a0cba"
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

	startDate := time.Now().UTC()
	newSubs := entity.Subscription{
		ID:          _subsUUID,
		ServiceName: "Yandex Plus",
		Price:       400,
		UserID:      _userUUID,
		StartDate:   &startDate,
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
	t.Log("Update subs")

	startDate := time.Now().UTC()
	updateValues := entity.Subscription{
		ID:          _subsUUID,
		ServiceName: "Kinopoisk",
		Price:       350,
		UserID:      _userUUID,
		StartDate:   &startDate,
	}

	updatedSubs, err := _repo.Update(&updateValues)
	require.NoError(t, err)

	t.Logf("Updated subs: %+v", updatedSubs)
}

func TestSubs_UpdateUnexisting(t *testing.T) {
	t.Log("Try to update unexisting subs")

	startDate := time.Now().UTC()
	updateValues := entity.Subscription{
		ID:          uuid.NewString(),
		ServiceName: "Kion",
		Price:       500,
		UserID:      _userUUID,
		StartDate:   &startDate,
	}

	_, err := _repo.Update(&updateValues)
	require.Error(t, err)
	require.ErrorIs(t, err, errors.ErrNotFound)

	t.Log("Unexisting subs")
}

func TestSubs_UpdateOneField(t *testing.T) {
	t.Log("Update just one subs field")

	updateValues := entity.Subscription{
		ID:          _subsUUID,
		ServiceName: "Ivi",
	}

	updatedSubs, err := _repo.Update(&updateValues)
	require.NoError(t, err)

	t.Logf("Updated subs: %+v", updatedSubs)
}

func TestSubs_GetSum(t *testing.T) {
	t.Log("Get sum of prices")

	subs := entity.SubscriptionSumFilter{
		UserID:      _userUUID,
		ServiceName: "Ivi",
	}

	total, err := _repo.GetSum(&subs)
	require.NoError(t, err)

	t.Logf("Total subs prices: %v", total)
	require.Equal(t, total, 350)
}

func TestSubs_Delete(t *testing.T) {
	t.Log("Remove subs by ID")

	err := _repo.Delete(_subsUUID)
	require.NoError(t, err)

	t.Logf("Subs with ID %s was deleted successfully", _subsUUID)
}
