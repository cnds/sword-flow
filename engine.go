package flow

import "github.com/jinzhu/gorm"

// TypePlace type of Place
type TypePlace uint

// TriggerType type of Trigger
type TriggerType string

// ArcType type of Arc
type ArcType string

// ArcDirection in or out
type ArcDirection uint

// CaseStatus status of CaseStatus
type CaseStatus string

//TokenStatus status of token
type TokenStatus string

const (
	// START start place
	START TypePlace = 0
	// END end place
	END TypePlace = 9
	// USER manually by a user of trigger
	USER TriggerType = "USER"
	// AUTO automatically by the system
	AUTO TriggerType = "AUTO"
	// MESSAGE by an external event
	MESSAGE TriggerType = "MESSAGE"
	// TIME after a time limit has expired
	TIME TriggerType = "TIME"
	// SEQ ordinary sequential flow, not a join or a split.
	SEQ ArcType = "SEQ"
	// EORS Explicit OR split.
	EORS ArcType = "EORS"
	// IORS Implicit OR split.
	IORS ArcType = "IORS"
	// ORJ explicit and implicit)
	ORJ ArcType = "ORJ"
	// ANDS and split
	ANDS ArcType = "ANDS"
	// ANDJ and join
	ANDJ ArcType = "ANDJ"
	// IN arc in direction
	IN ArcDirection = 0
	// OUT arc out direction
	OUT ArcDirection = 1
	// OPEN open status
	OPEN CaseStatus = "OPEN"
	// CLOSE close status
	CLOSE CaseStatus = "CLOSE"
	// SUSPEND suspended status
	SUSPEND CaseStatus = "SUSPEND"
	// CANCEL cancel status
	CANCEL CaseStatus = "CANCEL"
	// FREE free status of toke
	FREE TokenStatus = "FREE"
	//LOCK locked status of token
	LOCK TokenStatus = "LOCK"
	// CONSUME consumed status of token
	CONSUME TokenStatus = "CONSUME"
	// TCANCEL cancel status of token
	TCANCEL TokenStatus = "CANCEL"
)

// Option define option for flow engine
type Option struct {
	//TODO
}

// SwordEngine the engine for workflow
type SwordEngine struct {
	db *gorm.DB
	tx *gorm.DB
}
