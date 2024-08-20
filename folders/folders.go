package folders

import (
	"github.com/gofrs/uuid"
)


/**
New Suggested Function
Retrieve all folders associated with a specific `OrgID` by:
	1. fetching a list of filtered folder pointers from source using `FetchAllFoldersByOrgID` function
	2. copy the filtered folder pointers into a new slice `fp`
	3. creating and returning a `FetchFolderResponse` containing these folder pointers

Improvments:
	- Removed redundant loops that converts pointers->values->pointers, now function directly uses
	  `r` retrieved from `FetchAllFoldersByOrgID` for `fp`
	- Error handles potential error from `FetchAllFoldersByOrgID`, function will now immediately
	  return `nil` and the error
	- Unused variables from the previous `GetAllFolders` (`err`,`f1`,`fs`,`k`,`k1`) have been removed
*/
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	//fetch folders by OrgID
	r, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err
	}

	// copy into folders
	var fp []*Folder
	for _, v1 := range r {
		fp = append(fp, v1)
	}

	// create FetchFolderResponse with folder pointers
	ffr := &FetchFolderResponse{Folders: fp}
	return ffr, nil
}


/**
Original Function
Retrieve all folders associated with a specific `OrgID` by:
	1. fetching a list of filtered folder pointers from source using `FetchAllFoldersByOrgID` function
	2. convert these pointers to folder structs and store them in a temp variable `f`
	3. convert folder values in `f`` back into pointers, storing them in `fp`
	4. package these folder pointers `fp` into `FetchFolderResponse` struct and return this response

Deprecated: due to redudant loops from pointers -> folder structs -> values -> pointers
see improvement comments below + suggested improved function above

*NOTE: To compile, some lines were modified (these are marked with CHANGED)*
*/
func DeprecatedGetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	// err,f1,fs declared but not not used, can be removed (compiler provides error messages)
	// CHANGED: comment lines 57-61 out to compile
	// var (
	// 	err error
	// 	f1  Folder
	// 	fs  []*Folder
	// )

	// f unnecessary, used temporarily to hold values that are later converted back to pointers
	f := []Folder{}
	r, _ := FetchAllFoldersByOrgID(req.OrgID) // filter folders with `orgID` but ignores error

	// loop not necessary- no need to convert pointers to values
	// CHANGED: replace k with _ on line 69 to compile (k never used)
	for _, v := range r {
		f = append(f, *v)
	}

	// should use r to directly append to fp
	var fp []*Folder

	//loop redundant, converts values in `f` back into pointers (original `r` contains all pointers)
	// CHANGED: replace k1 with _ on line 78 to compile (k1 never used)
	for _, v1 := range f {
		fp = append(fp, &v1)
	}
	var ffr *FetchFolderResponse
	ffr = &FetchFolderResponse{Folders: fp}
	return ffr, nil
}

/**
Retrieves all folders associated with a given `OrgID`. Filters through a sample Dataset
and returns pointers to the folders that match `OrgID`
*/
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
