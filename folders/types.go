package folders

import "github.com/gofrs/uuid"

type FetchFolderRequest struct {
	OrgID uuid.UUID
	PageSize  int
	PageToken int
}

type FetchFolderResponse struct {
	Folders []*Folder
	NextToken *int
}