package folders

import (
	"fmt"
)

/**
Function: retrieves a paginated list of folders associated with a specific OrgID.
Fetches all folders matching the OrgID and returns a subset of those folders based on
the pagination parameters provided in req

Explaination:
1. First calls FetchAllFoldersByOrgID to retrieve all folders associated with the given OrgID
2. Error Handles PageToken
3. Paginate by calculating the start and end indexes, then slice the list of folders to get current page
4. Set NextToken, either to next index or nil
5. Package pageFolders and nextToken into FetchFolderResponse and return

(tests for GetAllFoldersPagination are written in folders_test.go)
*/
func GetAllFoldersPagination(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	//use the same FetchAllFoldersByOrgID
	r, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err
	}

	//error if PageToken is negative
	if req.PageToken < 0 {
		return nil, fmt.Errorf("invalid PageToken: must be non negative")
	}

	//error if PageToken is beyond the available data
	if req.PageToken > len(r) {
		return nil, fmt.Errorf("invalid PageToken: exceeds available data")
	}

	start := req.PageToken
	end := start + req.PageSize
	if end > len(r) {
		end = len(r)
	}

	pageFolders := r[start:end] //current page of folders

	//find next token (end of curr / nil)
	var nextToken *int
	if end < len(r) {
		nextTokenValue := end
		nextToken = &nextTokenValue
	} else {
		nextToken = nil
	}

	ffr := &FetchFolderResponse{
		Folders:   pageFolders,
		NextToken: nextToken,
	}

	return ffr, nil
}
