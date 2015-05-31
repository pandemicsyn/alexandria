package main

import (
	"os"
	"runtime"

	log "github.com/Sirupsen/logrus"
	"github.com/gocql/gocql"
	"github.com/miekg/dns"
	"github.com/spf13/viper"
)

func configureLogging(v *viper.Viper) {
	level, err := log.ParseLevel(v.GetString("log_level"))
	if err != nil {
		log.Fatalln(err)
	}
	log.SetLevel(level)

	if v.GetString("log_format") == "text" {
		log.SetFormatter(&log.TextFormatter{DisableColors: true, FullTimestamp: true})
	} else if v.GetString("log_format") == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.Errorln("Error: log_type invalid, defaulting to text")
		log.SetFormatter(&log.TextFormatter{})
	}
	switch v.GetString("log_target") {
	case "stdout":
		log.SetOutput(os.Stdout)
	case "stderr":
		log.SetOutput(os.Stderr)
	default:
		log.Errorln("Error: log_target invalid, defaulting to Stdout")
		log.SetOutput(os.Stdout)
	}
}

func main() {
	//var err error
	v := viper.New()
	v.SetDefault("log_level", "info")
	v.SetDefault("log_format", "text")
	v.SetDefault("log_target", "stdout")
	v.SetDefault("max_procs", 1)
	v.SetEnvPrefix("alexandria")
	v.SetConfigName("alexandria")
	v.AddConfigPath("/etc/alexandria/")
	v.ReadInConfig()

	configureLogging(v)
	runtime.GOMAXPROCS(v.GetInt("max_procs"))

	log.Warningln("alexandria starting up")
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "alexandria"
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	/*
		cluster := gocql.NewCluster("127.0.0.1")
		cluster.Keyspace = "alexandria"

		session, err := cluster.CreateSession()
		log.Println("Session error ", err)

		defer session.Close()

		var src, target, rtype string
		var rttl int
		if err := session.Query("SELECT * FROM domains WHERE src = 'ip.velocillama.com'").Scan(&src, &target, &rttl, &rtype); err != nil {
			log.Fatal(err)
		}
		log.Println(src, target, rtype, rttl)
	*/
	resolver := NewResolver(map[string]string{}, cluster, session)
	resolver.Entries["MX-ronin.io."] = "ronin.io. 3600 IN MX 10 mx.ronin.io."
	udpServer := &dns.Server{Addr: ":53", Net: "udp"}
	dns.HandleFunc(".", resolver.handleRequest)
	log.Fatalln(udpServer.ListenAndServe())
}
