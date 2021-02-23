/*
 *  Brown University, CS138, Spring 2020
 *
 *  Purpose: Defines the RoutingTable type and provides methods for interacting
 *  with it.
 */

package pkg

import (
	"sync"
)

// RoutingTable has a number of levels equal to the number of digits in an ID
// (default 40). Each level has a number of slots equal to the digit base
// (default 16). A node that exists on level n thereby shares a prefix of length
// n with the local node. Access to the routing table protected by a mutex.
type RoutingTable struct {
	local RemoteNode                 // The local tapestry node
	Rows  [DIGITS][BASE][]RemoteNode // The rows of the routing table
	mutex sync.Mutex                 // To manage concurrent access to the routing table (could also have a per-level mutex)
}

// NewRoutingTable creates and returns a new routing table, placing the local node at the
// appropriate slot in each level of the table.
func NewRoutingTable(me RemoteNode) *RoutingTable {
	t := new(RoutingTable)
	t.local = me

	// Create the node lists with capacity of SLOTSIZE
	for i := 0; i < DIGITS; i++ {
		for j := 0; j < BASE; j++ {
			t.Rows[i][j] = make([]RemoteNode, 0, SLOTSIZE)
		}
	}

	// Make sure each row has at least our node in it
	for i := 0; i < DIGITS; i++ {
		slot := t.Rows[i][t.local.ID[i]]
		t.Rows[i][t.local.ID[i]] = append(slot, t.local)
	}

	return t
}

// Add adds the given node to the routing table.
//
// Returns true if the node did not previously exist in the table and was subsequently added.
// Returns the previous node in the table, if one was overwritten.
func (t *RoutingTable) Add(node RemoteNode) (added bool, previous *RemoteNode) {
	// TODO: students should implement this
	return
}

// Remove removes the specified node from the routing table, if it exists.
// Returns true if the node was in the table and was successfully removed.
func (t *RoutingTable) Remove(node RemoteNode) (wasRemoved bool) {

	t.mutex.Lock()
	defer t.mutex.Unlock()

	// TODO: students should implement this
	return
}

// GetLevel get all nodes on the specified level of the routing table, EXCLUDING the local node.
func (t *RoutingTable) GetLevel(level int) (nodes []RemoteNode) {

	t.mutex.Lock()
	defer t.mutex.Unlock()

	// TODO: students should implement this

	return
}

// FindNextHop searches the table for the closest next-hop node for the provided ID starting at the given level.
func (t *RoutingTable) FindNextHop(id ID, level int32) RemoteNode {

	t.mutex.Lock()
	defer t.mutex.Unlock()

	// TODO: students should implement this
}
