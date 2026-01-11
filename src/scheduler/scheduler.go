package scheduler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgtype"
)

type CommandRequest struct {
	Command     string
	ScheduledAt string
}

type Task struct {
	Id          string
	Command     string
	ScheduledAt pgtype.Timestamp
	Status      string
	CompletedAt pgtype.Timestamp
	StartedAt   pgtype.Timestamp
	LockedUntil pgtype.Timestamp
}

type SchedulerServer struct {
	serverPort         string
	dbConnectionString string
	ctx                context.Context
	cancel             context.CancelFunc
	httpServer         *http.Server
}

func NewServer(port string, dbConnectionString string) *SchedulerServer {
	fmt.Println("Creating new server")
	ctx, cancel := context.WithCancel(context.Background())
	return &SchedulerServer{
		serverPort:         port,
		dbConnectionString: dbConnectionString,
		ctx:                ctx,
		cancel:             cancel,
	}
}

func (server *SchedulerServer) StartServer() {
	fmt.Println("starting server")
}
