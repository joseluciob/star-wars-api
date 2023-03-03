package persistence

import (
	"database/sql"
	"regexp"
	"star-wars-api/domain/entity"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type RepositoryPlanetSuite struct {
	suite.Suite
	conn   *sql.DB
	DB     *gorm.DB
	mock   sqlmock.Sqlmock
	repo   *PlanetRepo
	planet *entity.Planet
}

func (rs *RepositoryPlanetSuite) SetupSuite() {
	var (
		err error
	)
	rs.conn, rs.mock, err = sqlmock.New()
	assert.NoError(rs.T(), err)
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 rs.conn,
		PreferSimpleProtocol: true,
	})
	rs.DB, err = gorm.Open(dialector, &gorm.Config{})
	assert.NoError(rs.T(), err)
	rs.repo = NewPlanetRepository(rs.DB)
	assert.IsType(rs.T(), &PlanetRepo{}, rs.repo)
	rs.planet = &entity.Planet{
		ID:      1,
		Name:    "Iridonia",
		Climate: "unknown",
		Terrain: "rocky canyons, acid pools",
	}
}

func (rs *RepositoryPlanetSuite) AfterTest(_, _ string) {
	assert.NoError(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *RepositoryPlanetSuite) TestInsert() {
	rs.mock.ExpectBegin()
	rs.mock.ExpectExec(
		regexp.QuoteMeta(`INSERT INTO "planets" ("id","name","climate","terrain") VALUES ($1,$2,$3)`)).
		WithArgs(
			rs.planet.ID,
			rs.planet.Name,
			rs.planet.Climate,
			rs.planet.Terrain).
		WillReturnResult(sqlmock.NewResult(1, 1))
	rs.mock.ExpectCommit()
	p, err := rs.repo.Save(rs.planet)
	assert.NoError(rs.T(), err)
	assert.Equal(rs.T(), rs.planet, p)
}

func (rs *RepositoryPlanetSuite) TestUpdate() {
	rs.planet.Name = "Glee Anselm"
	rs.mock.ExpectBegin()
	rs.mock.ExpectExec(
		regexp.QuoteMeta(`UPDATE "planets" SET "name"=$1 WHERE "id" = $2`)).
		WithArgs(
			rs.planet.Name,
			rs.planet.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	rs.mock.ExpectCommit()
	p, err := rs.repo.Update(rs.planet)
	assert.NoError(rs.T(), err)
	assert.Equal(rs.T(), rs.planet, p)
}

func (rs *RepositoryPlanetSuite) TestDelete() {
	deleteAt := time.Now()
	rs.mock.ExpectBegin()
	rs.mock.ExpectExec(
		regexp.QuoteMeta(`UPDATE "planets" SET "deleted_at"=$1 WHERE id = $2 AND "planets"."deleted_at" IS NULL`)).
		WithArgs(deleteAt, rs.planet.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	rs.mock.ExpectCommit()
	err := rs.repo.Delete(rs.planet.ID)
	assert.NoError(rs.T(), err)
}

func (rs *RepositoryPlanetSuite) TestGetAll() {
	rows := sqlmock.NewRows([]string{"id", "name", "climate", "terrain"}).
		AddRow(
			rs.planet.ID,
			rs.planet.Name,
			rs.planet.Climate,
			rs.planet.Terrain).
		AddRow(
			5,
			"Sullust",
			"superheated",
			"mountains, volcanoes, rocky deserts")
	rs.mock.ExpectQuery(
		regexp.QuoteMeta(`SELECT * FROM "planets"`)).
		WithArgs().
		WillReturnRows(rows)
	planets, err := rs.repo.GetAll()
	assert.NoError(rs.T(), err)
	assert.Contains(rs.T(), planets, *rs.planet)
}

func TestMain(m *testing.M) {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2023, 03, 03, 20, 34, 58, 651387237, time.UTC)
	})
}

func TestSuite(t *testing.T) {

	suite.Run(t, new(RepositoryPlanetSuite))
}
