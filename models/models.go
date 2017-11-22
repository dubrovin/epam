package models

import "time"

// Product -
type Product struct {
	ID   int64         `json:"id"`
	Hash string        `json:"hash"`
	TTL  time.Duration `json:"ttl"`
}

// GetID -
func (p *Product) GetID() int64 {
	return p.ID
}

// SetID -
func (p *Product) SetID(id int64) error {
	p.ID = id
	return nil
}

// GetHash -
func (p *Product) GetHash() string {
	return p.Hash
}

// SetHash -
func (p *Product) SetHash(hash string) error {
	p.Hash = hash
	return nil
}

func (p *Product) GetTTL() time.Duration {
	return p.TTL
}

func (p *Product) SetTTL(ttl time.Duration) error {
	p.TTL = ttl
	return nil
}
