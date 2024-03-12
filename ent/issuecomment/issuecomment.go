// Code generated by ent, DO NOT EDIT.

package issuecomment

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the issuecomment type in the database.
	Label = "issue_comment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNodeID holds the string denoting the node_id field in the database.
	FieldNodeID = "node_id"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// FieldHTMLURL holds the string denoting the html_url field in the database.
	FieldHTMLURL = "html_url"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldIssueURL holds the string denoting the issue_url field in the database.
	FieldIssueURL = "issue_url"
	// FieldAuthorAssociation holds the string denoting the author_association field in the database.
	FieldAuthorAssociation = "author_association"
	// FieldReactions holds the string denoting the reactions field in the database.
	FieldReactions = "reactions"
	// EdgeIssue holds the string denoting the issue edge name in mutations.
	EdgeIssue = "issue"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the issuecomment in the database.
	Table = "issue_comments"
	// IssueTable is the table that holds the issue relation/edge.
	IssueTable = "issue_comments"
	// IssueInverseTable is the table name for the Issue entity.
	// It exists in this package in order to avoid circular dependency with the "issue" package.
	IssueInverseTable = "issues"
	// IssueColumn is the table column denoting the issue relation/edge.
	IssueColumn = "issue_comments"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "issue_comments"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_comments_created"
)

// Columns holds all SQL columns for issuecomment fields.
var Columns = []string{
	FieldID,
	FieldNodeID,
	FieldURL,
	FieldBody,
	FieldHTMLURL,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldIssueURL,
	FieldAuthorAssociation,
	FieldReactions,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "issue_comments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"issue_comments",
	"user_comments_created",
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

// AuthorAssociation defines the type for the "author_association" enum field.
type AuthorAssociation string

// AuthorAssociation values.
const (
	AuthorAssociationCOLLABORATOR           AuthorAssociation = "COLLABORATOR"
	AuthorAssociationCONTRIBUTOR            AuthorAssociation = "CONTRIBUTOR"
	AuthorAssociationFIRST_TIMER            AuthorAssociation = "FIRST_TIMER"
	AuthorAssociationFIRST_TIME_CONTRIBUTOR AuthorAssociation = "FIRST_TIME_CONTRIBUTOR"
	AuthorAssociationMANNEQUIN              AuthorAssociation = "MANNEQUIN"
	AuthorAssociationMEMBER                 AuthorAssociation = "MEMBER"
	AuthorAssociationNONE                   AuthorAssociation = "NONE"
	AuthorAssociationOWNER                  AuthorAssociation = "OWNER"
)

func (aa AuthorAssociation) String() string {
	return string(aa)
}

// AuthorAssociationValidator is a validator for the "author_association" field enum values. It is called by the builders before save.
func AuthorAssociationValidator(aa AuthorAssociation) error {
	switch aa {
	case AuthorAssociationCOLLABORATOR, AuthorAssociationCONTRIBUTOR, AuthorAssociationFIRST_TIMER, AuthorAssociationFIRST_TIME_CONTRIBUTOR, AuthorAssociationMANNEQUIN, AuthorAssociationMEMBER, AuthorAssociationNONE, AuthorAssociationOWNER:
		return nil
	default:
		return fmt.Errorf("issuecomment: invalid enum value for author_association field: %q", aa)
	}
}

// OrderOption defines the ordering options for the IssueComment queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByNodeID orders the results by the node_id field.
func ByNodeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNodeID, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByBody orders the results by the body field.
func ByBody(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBody, opts...).ToFunc()
}

// ByHTMLURL orders the results by the html_url field.
func ByHTMLURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHTMLURL, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByIssueURL orders the results by the issue_url field.
func ByIssueURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIssueURL, opts...).ToFunc()
}

// ByAuthorAssociation orders the results by the author_association field.
func ByAuthorAssociation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthorAssociation, opts...).ToFunc()
}

// ByIssueField orders the results by issue field.
func ByIssueField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIssueStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newIssueStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IssueInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, IssueTable, IssueColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}