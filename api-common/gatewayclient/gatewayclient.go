package gatewayclient

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"

	"git.ana/xjtuana/api-gateway/dto"
	"github.com/google/uuid"
)

type Route struct {
	Url    string
	Method string
	Scope  []string
}

type GatewayClient struct {
	id         uuid.UUID
	serverAddr *net.UDPAddr
	name       string
	selfIp     string
	httpPort   uint16
	stop       bool
	ticker     *time.Ticker
	mutex      sync.Mutex // protects following
	routes     []Route
}

// gatewayServer must be like 127.0.0.1:8000
func New(gatewayServer string, name string, selfIp string, httpPort uint16) (*GatewayClient, error) {
	serverAddr, err := net.ResolveUDPAddr("udp", gatewayServer)
	if err != nil {
		fmt.Println("Error resolving server address:", err)
		return nil, err
	}
	randomID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return &GatewayClient{
		serverAddr: serverAddr,
		id:         randomID,
		name:       name,
		selfIp:     selfIp,
		httpPort:   httpPort,
		stop:       true,
		ticker:     time.NewTicker(time.Second * 5),
	}, nil
}

func (c *GatewayClient) heartbeatAndRegister() {
	c.mutex.Lock()
	// TODO: Send Heartbeat
	for _, route := range c.routes {
		c.sendRegister(route.Url, route.Method, route.Scope)
	}
	c.mutex.Unlock()
}

func (c *GatewayClient) Start() {
	if !c.stop {
		return
	}
	go func() {
		for range c.ticker.C {
			c.heartbeatAndRegister()
		}
	}()
}

func (c *GatewayClient) Register(url, method string, scope []string) {
	c.mutex.Lock()
	c.routes = append(c.routes, Route{Url: url, Method: method, Scope: scope})
	c.mutex.Unlock()
}

func (c *GatewayClient) sendRegister(url, method string, scope []string) error {
	pkt := dto.Packet{
		Id:     c.id,
		Type:   "register",
		Name:   c.name,
		Ip:     c.selfIp,
		Port:   c.httpPort,
		Scope:  scope,
		Url:    url,
		Method: method,
	}

	conn, err := net.DialUDP("udp", nil, c.serverAddr)
	if err != nil {
		fmt.Println("Error creating UDP connection:", err)
		return err
	}
	defer conn.Close()

	data, err := json.Marshal(pkt)
	if err != nil {
		fmt.Println("Error marshalling UDP packet:", err)
	}

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("Error sending data:", err)
		return err
	}

	return nil
}
