package gomodparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Tortik3000/gomodinfo/internal/model"
	modelErr "github.com/Tortik3000/gomodinfo/internal/model/errors/parser"
)

func TestParser_Parse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		modBytes []byte
		wantInfo *model.ModuleInfo
		wantErr  error
	}{
		{
			name:     "empty go.mod",
			modBytes: []byte{},
			wantInfo: nil,
			wantErr:  modelErr.ErrEmptyGoMod,
		},
		{
			name:     "invalid syntax",
			modBytes: []byte("module github.com/example/test\nrequire ("),
			wantInfo: nil,
			wantErr:  modelErr.ErrInvalidGoModSyntax,
		},
		{
			name:     "missing module directive",
			modBytes: []byte("go 1.20\nrequire github.com/stretchr/testify v1.8.4"),
			wantInfo: nil,
			wantErr:  modelErr.ErrMissingModuleDirective,
		},
		{
			name:     "missing go version",
			modBytes: []byte("module github.com/example/test\nrequire github.com/stretchr/testify v1.8.4"),
			wantInfo: nil,
			wantErr:  modelErr.ErrMissingGoVersion,
		},
		{
			name: "valid go.mod",
			modBytes: []byte(`module github.com/example/test
			go 1.21
			require (
				github.com/stretchr/testify v1.8.4
				golang.org/x/mod v0.12.0
			)`),
			wantInfo: &model.ModuleInfo{
				Name:    "github.com/example/test",
				Version: "1.21",
				Deps: []*model.Dependency{
					{Name: "github.com/stretchr/testify", CurrentVersion: "v1.8.4"},
					{Name: "golang.org/x/mod", CurrentVersion: "v0.12.0"},
				},
			},
			wantErr: nil,
		},
		{
			name: "valid go.mod without deps",
			modBytes: []byte(`module github.com/example/naked
			go 1.22
			`),
			wantInfo: &model.ModuleInfo{
				Name:    "github.com/example/naked",
				Version: "1.22",
				Deps:    []*model.Dependency{},
			},
			wantErr: nil,
		},
	}

	parser := New()

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			info, err := parser.Parse(tt.modBytes)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, info)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, info)
			assert.Equal(t, tt.wantInfo.Name, info.Name)
			assert.Equal(t, tt.wantInfo.Version, info.Version)
			assert.Len(t, info.Deps, len(tt.wantInfo.Deps))

			for i, dep := range info.Deps {
				assert.Equal(t, tt.wantInfo.Deps[i].Name, dep.Name)
				assert.Equal(t, tt.wantInfo.Deps[i].CurrentVersion, dep.CurrentVersion)
			}
		})
	}
}
