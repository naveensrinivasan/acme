package certs

import (
	"log"
	"time"

	"github.com/naveensrinivasan/acme/internal/cache"
	"github.com/pborman/uuid"
)

const renewDuration = 10 * time.Minute

// Cert manages certificates.
type Cert interface {
	Get(name string) string
}

// internal struct with the interface implementation.
type cert struct {
	cache cache.Cache
}

// New provides a Cert which manages certs.
func New(cache cache.Cache) Cert {
	return cert{cache: cache}
}

// certificate struct is cached for certificate retrieval and last sync time
// the last sync time is cached because if the acme server restarts we need to know
// when was the last time the certs were refreshed so that the code can retrieve the cert.
// For this implementation it is using a timer simulating refresh every 10 minutes.
// When the acme server restarts it will lookup the cache for the certs that were last synced
// and decide when is the next time the timer to sync.This is the thought process and
// still not implemented.
type certificate struct {
	cert         string    // cert
	lastSyncTime time.Time // last time the cert was generated.
}

// Get generates and cache's certificates.
func (c cert) Get(name string) string {
	// return the cert if exists in the cache
	val, exists := c.cache.Get(name)
	if exists {
		certs, _ := val.(certificate)
		return certs.cert
	}

	time.Sleep(10 * time.Second) // simulating lets encrypt call

	uid := uuid.New() // generating a random uid for the cert
	t := time.Now()
	c.cache.Set(name,
		certificate{
			cert:         uid,
			lastSyncTime: t}) // caching the cert

	log.Printf("generated cert for %s %s \n", name, t)

	// renew certs every 10 minutes for each domain
	ticker := time.NewTicker(renewDuration)
	go func() {
		for t := range ticker.C {
			uid = uuid.New()
			c.cache.Set(name, certificate{cert: uid, lastSyncTime: t})
			log.Printf("cert updated for domain %s %s \n", name, t)
		}
	}()
	return uid
}
