package bitcask

import "errors"

var (
	ErrKeyIsEmpty             = errors.New("the key is empty")
	ErrIndexUpdatedFailed     = errors.New("failed to update index")
	ErrKeyNotFound            = errors.New("key not found in database")
	ErrDataFileNotFound       = errors.New("data file is not found")
	ErrDataDirectoryCorrupted = errors.New("the database directory maybe corrupted")
	ErrExceedMaxBatchNum      = errors.New("exceed the max batch num")
	ErrMergeIsProgress        = errors.New("merge is in progress, try again later")
	ErrDatabaseIsUsing        = errors.New("the database is used by another process")
)
