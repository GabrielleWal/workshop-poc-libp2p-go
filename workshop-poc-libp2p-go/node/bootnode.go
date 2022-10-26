package node

import (
    "fmt"
    "github.com/libp2p/go-libp2p"
    "github.com/libp2p/go-libp2p-core/host"
)

type BootNode struct {
    Host host.Host
}

func (b *BootNode) Start() error {
    node, err := libp2p.New()
    if err != nil {
        return err
    }
    b.Host = node
    return nil
}

func (b *BootNode) Stop() error {
    err := b.Host.Close()
    if (err != nil) {
        return err
    }
    return nil
}

func (b *BootNode) GetMultiAddress() (string, error) {
    if len(b.Host.Addrs()) < 3 {
        return "", fmt.Errorf("no multiadress available")
    }
    return b.Host.Addrs()[02].String() + "/p2p/" + b.Host.ID().String(), nil
}

func NewBootNode() *BootNode {
    return &BootNode{}
}

func (b *BootNode) StartDiscovery() string {
}
