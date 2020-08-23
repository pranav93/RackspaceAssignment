package models

// RulesDBMap key-val storage for rules
var RulesDBMap map[int]Rule

// ActionsDBMap key-val storage for actions
var ActionsDBMap map[int]Action

// ResultsDBMap key-val storage for results
var ResultsDBMap map[int]Result

// GetAllRulesDB gets all rules from key-val store
func GetAllRulesDB() map[int]Rule {
	return RulesDBMap
}
