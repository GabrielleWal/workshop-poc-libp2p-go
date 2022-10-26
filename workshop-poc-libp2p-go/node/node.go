package node

import (
)

type Node interface {
    Start() error
    Stop() error
    GetMultiAddress() string
	StartDiscovery() string
}