// Code generated by entc, DO NOT EDIT.

package workflowinstance

const (
	// Label holds the string label denoting the workflowinstance type in the database.
	Label = "workflow_instance"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldInstanceID holds the string denoting the instanceid field in the database.
	FieldInstanceID = "instance_id"
	// FieldInvokedBy holds the string denoting the invokedby field in the database.
	FieldInvokedBy = "invoked_by"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldRevision holds the string denoting the revision field in the database.
	FieldRevision = "revision"
	// FieldBeginTime holds the string denoting the begintime field in the database.
	FieldBeginTime = "begin_time"
	// FieldEndTime holds the string denoting the endtime field in the database.
	FieldEndTime = "end_time"
	// FieldFlow holds the string denoting the flow field in the database.
	FieldFlow = "flow"
	// FieldInput holds the string denoting the input field in the database.
	FieldInput = "input"
	// FieldOutput holds the string denoting the output field in the database.
	FieldOutput = "output"
	// FieldStateData holds the string denoting the statedata field in the database.
	FieldStateData = "state_data"
	// FieldMemory holds the string denoting the memory field in the database.
	FieldMemory = "memory"
	// FieldDeadline holds the string denoting the deadline field in the database.
	FieldDeadline = "deadline"
	// FieldAttempts holds the string denoting the attempts field in the database.
	FieldAttempts = "attempts"
	// FieldErrorCode holds the string denoting the errorcode field in the database.
	FieldErrorCode = "error_code"
	// FieldErrorMessage holds the string denoting the errormessage field in the database.
	FieldErrorMessage = "error_message"
	// FieldStateBeginTime holds the string denoting the statebegintime field in the database.
	FieldStateBeginTime = "state_begin_time"
	// FieldController holds the string denoting the controller field in the database.
	FieldController = "controller"
	// EdgeWorkflow holds the string denoting the workflow edge name in mutations.
	EdgeWorkflow = "workflow"
	// EdgeInstance holds the string denoting the instance edge name in mutations.
	EdgeInstance = "instance"
	// Table holds the table name of the workflowinstance in the database.
	Table = "workflow_instances"
	// WorkflowTable is the table that holds the workflow relation/edge.
	WorkflowTable = "workflow_instances"
	// WorkflowInverseTable is the table name for the Workflow entity.
	// It exists in this package in order to avoid circular dependency with the "workflow" package.
	WorkflowInverseTable = "workflows"
	// WorkflowColumn is the table column denoting the workflow relation/edge.
	WorkflowColumn = "workflow_instances"
	// InstanceTable is the table that holds the instance relation/edge.
	InstanceTable = "workflow_events"
	// InstanceInverseTable is the table name for the WorkflowEvents entity.
	// It exists in this package in order to avoid circular dependency with the "workflowevents" package.
	InstanceInverseTable = "workflow_events"
	// InstanceColumn is the table column denoting the instance relation/edge.
	InstanceColumn = "workflow_instance_instance"
)

// Columns holds all SQL columns for workflowinstance fields.
var Columns = []string{
	FieldID,
	FieldInstanceID,
	FieldInvokedBy,
	FieldStatus,
	FieldRevision,
	FieldBeginTime,
	FieldEndTime,
	FieldFlow,
	FieldInput,
	FieldOutput,
	FieldStateData,
	FieldMemory,
	FieldDeadline,
	FieldAttempts,
	FieldErrorCode,
	FieldErrorMessage,
	FieldStateBeginTime,
	FieldController,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "workflow_instances"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"workflow_instances",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
