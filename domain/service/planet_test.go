package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"star-wars-api/configs"
	"star-wars-api/infrastructure/persistence"
	"star-wars-api/infrastructure/provider"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewRepositoriesMock(config *configs.Configs) *persistence.Repositories {

	conn, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 conn,
		PreferSimpleProtocol: true,
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	return &persistence.Repositories{
		Planet: persistence.NewPlanetRepository(db),
		Film:   persistence.NewFilmRepository(db),
	}
}

func Test_Import(t *testing.T) {

	type args struct {
		ctx context.Context
	}

	type fields struct {
		repos    *persistence.Repositories
		provider func(ctrl *gomock.Controller) *provider.MockProvider
	}

	cfg := &configs.Configs{}
	repos := NewRepositoriesMock(cfg)

	response := provider.GetPlanetsResponse{}
	if err := json.Unmarshal(getFile("test_endpoint_planets"), &response); err != nil {
		t.Errorf("failed to read testdata file: %v", err)
		return
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		assert  bool
		wantErr bool
	}{
		{
			name: "should call the import planets method and not return errors",
			fields: fields{
				repos: repos,
				provider: func(ctrl *gomock.Controller) *provider.MockProvider {
					p := provider.NewMockProvider(ctrl)
					p.EXPECT().
						GetPlanets(gomock.All(), gomock.All()).
						Times(1).
						Return(&provider.GetPlanetsResponse{}, nil)

					return p
				},
			},
			args: args{
				ctx: context.Background(),
			},
			assert:  false,
			wantErr: false,
		},
		{
			name: "should call the import planets method and simulate an API iteration",
			fields: fields{
				repos: repos,
				provider: func(ctrl *gomock.Controller) *provider.MockProvider {
					p := provider.NewMockProvider(ctrl)
					p.EXPECT().
						GetPlanets(gomock.All(), gomock.All()).
						Times(2).
						Return(&response, nil).
						DoAndReturn(func(ctx context.Context, page int) (*provider.GetPlanetsResponse, error) {
							if page == 1 {
								return &response, nil
							} else {
								return &provider.GetPlanetsResponse{}, nil
							}
						}).AnyTimes()

					return p
				},
			},
			args: args{
				ctx: context.Background(),
			},
			assert:  false,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			planet := NewPlanetService(tt.fields.repos, tt.fields.provider(ctrl))
			planet.Import(tt.args.ctx)
		})
	}
}

func getFile(name string) []byte {
	file := fmt.Sprintf("./testdata/%s.json", name)
	p, _ := ioutil.ReadFile(file)

	return p
}
