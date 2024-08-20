package folders_test

import (
	"testing"
	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllFolders(t *testing.T) {
	t.Run("Default and valid OrgID returns folders", func(t *testing.T) {
		req := &folders.FetchFolderRequest{
			OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
		}

		res, err := folders.GetAllFolders(req)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Greater(t, len(res.Folders), 0) // at least 1 folder

		//check has valid id, name, OrgId
		for _, folder := range res.Folders {
			assert.NotEqual(t, uuid.Nil, folder.Id)
			assert.NotEmpty(t, folder.Name)
			assert.Equal(t, req.OrgID, folder.OrgId)
		}
	})

	t.Run("Invalid OrgID returns NilOrgID + no folders", func(t *testing.T) {
		req := &folders.FetchFolderRequest{
			OrgID: uuid.FromStringOrNil("hello"),
		}

		res, err := folders.GetAllFolders(req)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 0, len(res.Folders)) //no folders
		assert.Equal(t, uuid.Nil, req.OrgID) //OrgID should be nil bc invalid UUID input
	})

	t.Run("Non-existent OrgID returns no folders", func(t *testing.T) {
		req := &folders.FetchFolderRequest{
			OrgID: uuid.Must(uuid.NewV4()), //generate a new UUID
		}

		res, err := folders.GetAllFolders(req)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 0, len(res.Folders)) //no folder
	})

	t.Run("Nil OrgID returns no folders", func(t *testing.T) {
		req := &folders.FetchFolderRequest{
			OrgID: uuid.Nil, // use nil UUID
		}

		res, err := folders.GetAllFolders(req)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 0, len(res.Folders)) //no folder
	})

	t.Run("No OrgID provided returns no folders", func(t *testing.T) {
        req := &folders.FetchFolderRequest{} //create req with no OrgID

        res, err := folders.GetAllFolders(req)

        assert.NoError(t, err)
        assert.NotNil(t, res)
        assert.Equal(t, 0, len(res.Folders))
    })

}
