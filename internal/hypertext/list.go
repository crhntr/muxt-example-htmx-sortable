package hypertext

import "github.com/crhntr/muxt-example-htmx-sortable/internal/database"

type ListData struct {
	Err   error
	List  database.List
	Tasks []database.TasksByListIDRow
}

type UpdateListForm struct {
	TaskIDs []int32 `name:"task-id"`
}
