package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/gocql/gocql"
	"github.com/miekg/dns"
)

type Resolver struct {
	Entries map[string]string
	c       *gocql.ClusterConfig
	s       *gocql.Session
}

func NewResolver(entries map[string]string, c *gocql.ClusterConfig, s *gocql.Session) *Resolver {
	return &Resolver{Entries: entries, c: c, s: s}
}

func (res *Resolver) lookupAnswer(name, qtype string) ([]*dns.RR, error) {
	var answers []*dns.RR
	var qanswer string
	iter := res.s.Query("SELECT answer FROM qanswer WHERE question = ?", fmt.Sprintf("%s-%s", qtype, name)).Iter()
	for iter.Scan(&qanswer) {
		rra, err := dns.NewRR(qanswer)
		if err != nil {
			log.Println("Malformed answer:", err)
		} else {
			answers = append(answers, &rra)
		}
	}
	if err := iter.Close(); err != nil {
		log.Println(err)
	}
	if len(answers) != 0 {
		return answers, nil
	} else {
		return answers, fmt.Errorf("no entries")
	}
}

func (res *Resolver) handleRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	ar, err := res.lookupAnswer(r.Question[0].Name, dns.TypeToString[r.Question[0].Qtype])
	if err != nil {
		log.Println(err)
		dns.HandleFailed(w, r)
		return
	}
	for k, _ := range ar {
		m.Answer = append(m.Answer, *ar[k])
	}
	w.WriteMsg(m)
}
