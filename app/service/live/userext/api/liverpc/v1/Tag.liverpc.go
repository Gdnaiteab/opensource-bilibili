// Code generated by protoc-gen-liverpc v0.1, DO NOT EDIT.
// source: v1/Tag.proto

package v1

import context "context"

import proto "github.com/golang/protobuf/proto"
import "go-common/library/net/rpc/liverpc"

var _ proto.Message // generate to suppress unused imports

// =============
// Tag Interface
// =============

type TagRPCClient interface {
	// * 查询预约状态
	//
	Test(ctx context.Context, req *TagTestReq, opts ...liverpc.CallOption) (resp *TagTestResp, err error)
}

// ===================
// Tag Live Rpc Client
// ===================

type tagRPCClient struct {
	client *liverpc.Client
}

// NewTagRPCClient creates a client that implements the TagRPCClient interface.
func NewTagRPCClient(client *liverpc.Client) TagRPCClient {
	return &tagRPCClient{
		client: client,
	}
}

func (c *tagRPCClient) Test(ctx context.Context, in *TagTestReq, opts ...liverpc.CallOption) (*TagTestResp, error) {
	out := new(TagTestResp)
	err := doRPCRequest(ctx, c.client, 1, "Tag.test", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}
