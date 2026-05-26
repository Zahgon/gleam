package util

import (
	"google.golang.org/grpc"
)

func grpcDial(address string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	// opts = append(opts, grpc.WithBlock())
	// opts = append(opts, grpc.WithTimeout(time.Duration(5*time.Second)))
	return nil, nil
}

// grpc.WithInsecure(),

// client ping server if no activity for this long

func GleamGrpcDial(address string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseServerToGrpcAddress(server string) (serverGrpcAddress string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}
