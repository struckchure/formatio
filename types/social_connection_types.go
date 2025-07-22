package types

import "pkg.formatio/lib"

type ListConnectionsArgs struct {
	lib.BaseListFilterArgs

	UserId         string `json:"userId"`
	ConnectionType string `json:"connectionType"`
}

type CreateConnectionArgs struct {
	UserId         string `json:"userId"`
	ConnectionId   string `json:"connectionId"`
	ConnectionType string `json:"connectionType"`
}

type GetConnectionArgs struct {
	ID string `json:"id"`
}

type UpdateConnectionArgs struct {
	UserId         string `json:"userId"`
	ConnectionId   string `json:"connectionId"`
	ConnectionType string `json:"connectionType"`
}

type DeleteConnectionArgs struct {
	ID string `json:"id"`
}
