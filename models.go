package flow

import (
	"time"

	"github.com/jinzhu/gorm"
)

// SfWorkFlows This table holds the definition for each workflow process, such as 'Order Fulfillment'.
type SfWorkFlows struct {
	gorm.Model
	Name        string          `gorm:"type:varchar(100)"`
	description string          `gorm:"type:text"`
	IsValid     bool            `gorm:"default:false"`
	IsDraft     bool            `gorm:"default:true"`
	CreatorID   uint            `gorm:"type:bigint"`
	ErrorMsg    string          `gorm:"type:text"`
	Places      []SfPlaces      `gorm:"foreignkey:WorkflowID"`
	Transitions []SfTransitions `gorm:"foreignkey:WorkflowID"`
	Arcs        []SfArcs        `gorm:"foreignkey:WorkflowID"`
	Cases       []SfCases       `gorm:"foreignkey:WorkflowID"`
	WorkItems   []SfWorkItems   `gorm:"foreignkey:WorkflowID"`
	Tokens      []SfTokens      `gorm:"foreignkey:WorkflowID"`
}

// SfPlaces This table holds the details for each place within a workflow process.
type SfPlaces struct {
	gorm.Model
	WorkflowID  uint        `gorm:"type:bigint"`
	Name        string      `gorm:"type:text"`
	Description string      `gorm:"type:text"`
	SortOrder   uint        `gorm:"default:0"`
	PlaceType   uint        `gorm:"default:0"`
	WorkFlow    SfWorkFlows `gorm:"foreignkey:WorkflowID"`
	Arcs        []SfArcs    `gorm:"foreignkey:PlaceID"`
	Tokens      []SfTokens  `gorm:"foreignkey:PlaceID"`
}

// SfTransitions his table holds the details for each transition within a workflow process,
// such as 'Charge Customer', 'Pack Order' and 'Ship Order'.
// Each record will point to an application task within the MENU database.
type SfTransitions struct {
	gorm.Model
	WorkflowID  uint                      `gorm:"type:bigint"`
	Name        string                    `gorm:"type:varchar(100)"`
	Description string                    `gorm:"type:text"`
	SortOrder   uint                      `gorm:"default:0"`
	WorkFlow    SfWorkFlows               `gorm:"foreignkey:WorkflowID"`
	Arcs        []SfArcs                  `gorm:"foreignkey:TransitionID"`
	Assignments []SfTransitionAssignments `gorm:"foreignkey:TransitionID"`
	TiggerLimit uint                      `gorm:"type:int"`
	TiggerType  uint                      `gorm:"default:0"`
}

// SfArcs This table holds the details for each arc within a workflow process.
// An arc links a place to a transition.
type SfArcs struct {
	gorm.Model
	WorkflowID         uint   `gorm:"type:bigint"`
	TransitionID       uint   `gorm:"type:bigint"`
	PlaceID            uint   `gorm:"type:bigint"`
	Direction          uint   `gorm:"default:0"`
	ArcType            uint   `gorm:"default:0"`
	ConditionField     string `gorm:"type:varchar(50)"`
	ConditionOp        string `gorm:"type:varchar(50)"`
	ConditionValue     string `gorm:"type:varchar(50)"`
	ConditionExp       string `gorm:"type:varchar(50)"`
	ConditionFieldType string `gorm:"type:varchar(50)"`
}

// SfTransitionAssignments This Tbale holds assign for RBAC
type SfTransitionAssignments struct {
	gorm.Model
	WorkflowID     uint          `gorm:"type:bigint"`
	TransitionID   uint          `gorm:"type:bigint"`
	AssignableType string        `gorm:"type:varchar(100)"`
	AssignableID   string        `gorm:"type:varchar(100)"`
	Workflow       SfWorkFlows   `gorm:"foreignkey:WorkflowID"`
	Transition     SfTransitions `gorm:"foreignkey:TransitionID"`
}

// SfCases This identifies when a particular instance of a workflow was started,
// its context and its current status.
type SfCases struct {
	gorm.Model
	WorkflowID     uint          `gorm:"type:bigint"`
	TargetableType uint          `gorm:"type:varchar(100)"`
	TargetableID   uint          `gorm:"type:varchar(100)"`
	State          uint          `gorm:"type:varchar(100)"`
	Workflow       SfWorkFlows   `gorm:"foreignkey:WorkflowID"`
	WorkItems      []SfWorkItems `gorm:"foreignkey:CaseID"`
	Tokens         []SfTokens    `gorm:"foreignkey:CaseID"`
}

// SfTokens This identifies when a token was inserted into a particular place.
type SfTokens struct {
	gorm.Model
	WorkflowID       uint        `gorm:"type:bigint"`
	CaseID           uint        `gorm:"type:bigint"`
	TargetableType   string      `gorm:"type:varchar(100)"`
	TargetableID     string      `gorm:"type:varchar(100)"`
	PlaceID          uint        `gorm:"type:bigint"`
	State            uint        `gorm:"default:0"`
	WorkItemID       uint        `gorm:"type:bigint"`
	LockedWorkItemID uint        `gorm:"type:bigint"`
	WorkFlow         SfWorkFlows `gorm:"foreignkey:WorkflowID"`
	Case             SfCases     `gorm:"foreignkey:CaseID"`
	Place            SfPlaces    `gorm:"foreignkey:PlaceID"`
	ProducedAt       *time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	LockedAt         *time.Time
	CanceledAt       *time.Time
	ConsumedAt       *time.Time
}

// SfWorkItems A record is created here when a transition is enabled or able to fire.
// Entries which are to be triggered by a human participant will appear on the Menu/Home Page of relevant users
// so that they can see what tasks are pending and select any for processing.
// Each of these entries will be a hyperlink which, when pressed,
// will cause the relevant application task to be activated with the correct context already loaded.
type SfWorkItems struct {
	gorm.Model
	WorkflowID     uint          `gorm:"type:bigint"`
	CaseID         uint          `gorm:"type:bigint"`
	TransitionID   uint          `gorm:"type:bigint"`
	TargetableType string        `gorm:"type:varchar(100)"`
	TargetableID   string        `gorm:"type:varchar(100)"`
	State          uint          `gorm:"default:0"`
	UserType       string        `gorm:"type:varchar(100)"`
	UserID         uint          `gorm:"type:bigint"`
	WorkFlow       SfWorkFlows   `gorm:"foreignkey:WorkflowID"`
	Transition     SfTransitions `gorm:"foreignkey:TransitionID"`
	Case           SfCases       `gorm:"foreignkey:CaseID"`
	EnabledAt      *time.Time    `gorm:"default:CURRENT_TIMESTAMP"`
	StartedAt      *time.Time
	CanceledAt     *time.Time
	FinishedAt     *time.Time
	OverriddenAt   *time.Time
	Deadline       *time.Time
}
