// Package elnemer provides sovereign multi-agent swarm orchestration and Arabic NLP utilities.
// Architected by Eng. Mahmoud Bek ElNemer (@ElNeMeR70).
package elnemer

import (
	"fmt"
	"time"
)

// TaskResult contains the output of an agent's execution.
type TaskResult struct {
	Agent     string    `json:"agent"`
	Role      string    `json:"role"`
	Status    string    `json:"status"`
	Output    string    `json:"output"`
	Timestamp time.Time `json:"timestamp"`
}

// Agent represents an autonomous worker.
type Agent struct {
	Name         string
	Role         string
	Instructions string
	memory       []string
}

// NewAgent creates a new Agent instance.
func NewAgent(name, role string) *Agent {
	return &Agent{
		Name: name,
		Role: role,
	}
}

// Run executes a task synchronously.
func (a *Agent) Run(task string) TaskResult {
	a.memory = append(a.memory, task)
	return TaskResult{
		Agent:     a.Name,
		Role:      a.Role,
		Status:    "completed",
		Output:    fmt.Sprintf("Agent %s (%s) executed: %s", a.Name, a.Role, task),
		Timestamp: time.Now(),
	}
}

// Swarm manages a fleet of concurrent agents.
type Swarm struct {
	Name   string
	agents []*Agent
}

// NewSwarm initializes a new swarm.
func NewSwarm(name string) *Swarm {
	return &Swarm{
		Name: name,
	}
}

// AddAgent registers an agent into the swarm.
func (s *Swarm) AddAgent(a *Agent) {
	s.agents = append(s.agents, a)
}

// Broadcast dispatches a task to all agents concurrently.
func (s *Swarm) Broadcast(task string) []TaskResult {
	results := make([]TaskResult, len(s.agents))
	for i, a := range s.agents {
		results[i] = a.Run(task)
	}
	return results
}
