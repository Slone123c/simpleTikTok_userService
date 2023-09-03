// Code generated by Kitex v0.6.2. DO NOT EDIT.
package relationservice

import (
	server "github.com/cloudwego/kitex/server"
	RelationService "github.com/cold-runner/simpleTikTok/kitex_gen/RelationService"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler RelationService.RelationService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
