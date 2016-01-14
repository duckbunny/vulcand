// Copyright © 2015 Jason Smith <jasonrichardsmith@gmail.com>.
//
// Use of this source code is governed by the LGPL-3
// license that can be found in the LICENSE file.

package vulcand

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/coreos/etcd/Godeps/_workspace/src/golang.org/x/net/context"
	"github.com/coreos/etcd/client"
	"github.com/duckbunny/etcd"
	"github.com/duckbunny/herald"
	"github.com/duckbunny/service"
)

var (
	KVpath string = "/vulcand/backends"
	TTL    int    = 15
	// Title for specifying herald in flags
	Title string = "vulcand"
)

type Backend struct {
	Type `json:"Type"`
}

type Server struct {
	URL `json:"URL"`
}

func init() {
	ttl := os.Getenv("VULCAND_TTL")
	if ttl != "" {
		newttl, err := strconv.Atoi(ttl)
		if err != nil {
			log.Fatal(err)
		}
		TTL = newttl
	}
	flag.IntVar(&TTL, "vulcand-ttl", TTL, "TTL for vulcand microservice heartbeats.")
}

type Vulcand struct {
	*etcd.Etcd
}

func New() *Etcd {
	return new(Vulcand{etcd.New()})
}

func (v *Vulcand) Start(s *service.Service) error {
	if err != nil {
		return err
	}
	key := BackendPath(s)
	b := Backend{s.Protocol}
	js, err := json.Marshal(b)
	if err != nil {
		return err
	}
	_, err = v.Kapi.Set(context.Background(), key, string(js), nil)
	if err != nil {
		return ProcessEtcdErrors(err)
	}
	err = setServer(s)
	if err != nil {
		return err
	}
	go heartBeat(s)
}

func heartBeat(s service.Service) {
	for _ = range time.Tick(time.Duration(TTL-1) * time.Second) {
		setServer(s)
	}
}

func setServer(s service.Service) error {
	key := ServerPath(s)
	s := Server{Url(s)}
	so := client.SetOptions{TTL: TTL}
	_, err = v.Kapi.Set(context.Background(), key, string(js), so)
	if err != nil {
		return ProcessEtcdErrors(err)
	}

}

func (v *Vulcand) Stop(s *service.Service) error {
	key := ServerPath(s)
	p, err = v.Kapi.Delete(context.Background(), key, nil)
	if err != nil {
		return ProcessEtcdErrors(err)
	}
}

func BackendPath(s *service.Service) string {
	return fmt.Sprintf("%v/backend", basePath(s))
}

func ServerPath(s *service.Service) string {
	return fmt.Sprintf("%v/servers/%v.%v", basePath(s), environment.Host(), s.Port)
}

func BasePath(s *service.Service) string {
	return fmt.Sprintf("%v/%v.%v.%v", KVPath, s.Domain, s.Title, s.Version)
}
func Url(s *service.Service) string {
	return fmt.Sprintf("%v:%v", environment.Host(), s.Port)
}

// Register this herald with consul
func Register() {
	c := New()
	herald.AddPool(Title, c)
	herald.AddDeclaration(Title, c)
}
