package model

type Commit2Out struct {
	Phase1Out string   `json:"phase1_output"`
	ProveID   [32]byte `json:"prover_id"`
	SectorID  uint64   `json:"sector_id"`
	Miner     string   `json:"miner"`
}
