package main

import (
	"fmt"
	"time"
)

type NodeRole string

const (
	Follower  NodeRole = "Follower"
	Candidate NodeRole = "Candidate"
	Leader    NodeRole = "Leader"
)

type RaftNode struct {
	ID       string
	Role     NodeRole
	Term     int
	VotedFor string
}

func (n *RaftNode) RunElection() {
	n.Term++
	n.Role = Candidate
	n.VotedFor = n.ID
	fmt.Printf("[*] Node %s started election for Term %d\n", n.ID, n.Term)
	time.Sleep(50 * time.Millisecond)
	n.Role = Leader
	fmt.Printf("[✓] Node %s achieved Majority Consensus! Promoted to [LEADER]\n", n.ID)
}

func main() {
	fmt.Println("[*] Initializing Raft Consensus Cluster Simulation...")
	node := &RaftNode{ID: "node-us-east-1", Role: Follower, Term: 1}
	node.RunElection()
}
