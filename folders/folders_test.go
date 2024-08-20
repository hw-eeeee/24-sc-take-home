package folders_test

import (
	"testing"
	"math"
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

func Test_GetAllFolders_Pagination(t *testing.T) {
	orgID := uuid.FromStringOrNil(folders.DefaultOrgID)

	t.Run("First page of folders", func(t *testing.T) {
		req := &folders.FetchFolderRequest{
			OrgID:    orgID,
			PageSize: 3,
			PageToken: 0,
		}

		res, err := folders.GetAllFoldersPagination(req)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 3, len(res.Folders)) //3 folders in each page

		//there's a next token, should be set to the next index (3)
		if assert.NotNil(t, res.NextToken) {
			assert.Equal(t, 3, *res.NextToken)
		}
	})

	t.Run("Fetch second page folders", func(t *testing.T) {
		req := &folders.FetchFolderRequest{
			OrgID:    orgID,
			PageSize: 3,
			PageToken: 0,
		}

		res, err := folders.GetAllFoldersPagination(req)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 3, len(res.Folders)) //3 folders in each page

		//there's a next token, should be set to the next index (3)
		if assert.NotNil(t, res.NextToken) {
			assert.Equal(t, 3, *res.NextToken)
		}

		// Use the next token to get the next page of results
		secondReq := &folders.FetchFolderRequest{
			OrgID:    orgID,
			PageSize: 3,
			PageToken: *res.NextToken,
		}

		secondRes, err := folders.GetAllFoldersPagination(secondReq)

		assert.NoError(t, err)
		assert.NotNil(t, secondRes)
		assert.Equal(t, 3, len(secondRes.Folders))

		if assert.NotNil(t, secondRes.NextToken) {
			assert.Equal(t, 6, *secondRes.NextToken) // next token is at index 6
		}
	})

	t.Run("Negative PageToken returns error", func(t *testing.T) {
		req := &folders.FetchFolderRequest{
			OrgID:    orgID,
			PageSize: 3,
			PageToken: -1,
		}

		res, err := folders.GetAllFoldersPagination(req)

		//expect error, nil and error message
		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "invalid PageToken: must be non negative")
	})

	t.Run("PageToken bigger than avaiable data error", func(t *testing.T) {
		req := &folders.FetchFolderRequest{
			OrgID:    orgID,
			PageSize: 3,
			PageToken: math.MaxInt, // beyond the available data
		}

		res, err := folders.GetAllFoldersPagination(req)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "invalid PageToken: exceeds available data")
	})

	t.Run("Req page size larger than available data", func(t *testing.T) {
		req := &folders.FetchFolderRequest{
			OrgID:    orgID,
			PageSize: math.MaxInt,
			PageToken: 0,
		}

		res, err := folders.GetAllFoldersPagination(req)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.LessOrEqual(t, len(res.Folders), math.MaxInt)
		assert.Nil(t, res.NextToken)
	})

}
