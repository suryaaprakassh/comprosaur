package context

import (
	"log/slog"
	"path/filepath"
)

type Context struct {
	path   string
	status string

	logger *slog.Logger
}

func (c *Context) AppendPath(path string) {
	c.path = filepath.Join(c.path, path)
}

func (c *Context) UpdatePath(path string) {
	c.path = path
}

func (c *Context) GetPath() string {
	return c.path
}

func (c *Context) UpdateStatus(status string) {
	c.status = status
}

func (c *Context) GetStatus() string {
	return c.status
}

func (c *Context) ResetStatus() {
	c.status = ""
}

func (c *Context) Logger() *slog.Logger {
	return c.logger
}

func New(path string, logger *slog.Logger) *Context {
	return &Context{
		path:   path,
		logger: logger,
	}
}

type CTX interface {
	//update the current path of the application
	UpdatePath(string)
	//update the current path of the application
	AppendPath(string)
	//get the current path of the application
	GetPath() string
	//update the application status
	UpdateStatus(string)
	//get the application status for rendering
	GetStatus() string
	//resets the application status
	ResetStatus()

	//gets default logger of the application
	Logger() *slog.Logger
}
