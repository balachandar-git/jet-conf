package grpc

import (
	"time"

	"github.com/balachandar-git/jet-config/source"
	proto "github.com/balachandar-git/jet-config/source/grpc/proto"
)

func toChangeSet(c *proto.ChangeSet) *source.ChangeSet {
	return &source.ChangeSet{
		Data:      c.Data,
		Checksum:  c.Checksum,
		Format:    c.Format,
		Timestamp: time.Unix(c.Timestamp, 0),
		Source:    c.Source,
	}
}
