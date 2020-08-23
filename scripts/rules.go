package scripts

import "github.com/pranav93/RackspaceAssignment/models"

// CreateRulesDB creates rules in in-memory key-val store
func CreateRulesDB() {
	models.RulesDBMap = map[int]models.Rule{}
	models.ActionsDBMap = map[int]models.Action{}
	models.ResultsDBMap = map[int]models.Result{}

	// Rule 1
	action1 := models.Action{
		ID:          1,
		RuleID:      1,
		ProductCode: "CF1",
		Operator:    "ge",
		Qty:         1,
	}
	result1 := models.Result{
		ID:          1,
		RuleID:      1,
		ProductCode: "CF1",
		Qty:         -1, // unlimited bogo
	}
	rule1 := models.Rule{
		ID:     1,
		Rtype:  "getFree",
		Name:   "BOGO",
		Action: action1,
		Result: result1,
	}
	models.ActionsDBMap[1] = action1
	models.ResultsDBMap[1] = result1
	models.RulesDBMap[1] = rule1

	// Rule 2
	action2 := models.Action{
		ID:          2,
		RuleID:      2,
		ProductCode: "AP1",
		Operator:    "ge",
		Qty:         3,
	}
	result2 := models.Result{
		ID:               2,
		RuleID:           2,
		ProductCode:      "AP1",
		Qty:              -1,
		ResultType:       "applyPrice",
		AppliedPrice:     4.5,
		AppliedPriceType: "absolute",
	}
	rule2 := models.Rule{
		ID:     2,
		Rtype:  "applyPrice",
		Name:   "APPL",
		Action: action2,
		Result: result2,
	}
	models.ActionsDBMap[2] = action2
	models.ResultsDBMap[2] = result2
	models.RulesDBMap[2] = rule2

	// Rule 3
	action3 := models.Action{
		ID:          3,
		RuleID:      3,
		ProductCode: "CH1",
		Operator:    "ge",
		Qty:         1,
	}
	result3 := models.Result{
		ID:          3,
		RuleID:      3,
		ProductCode: "MK1",
		Qty:         1,
	}
	rule3 := models.Rule{
		ID:     2,
		Rtype:  "getFree",
		Name:   "CHMK",
		Action: action3,
		Result: result3,
	}
	models.ActionsDBMap[3] = action3
	models.ResultsDBMap[3] = result3
	models.RulesDBMap[3] = rule3

	// Rule 4
	action4 := models.Action{
		ID:          4,
		RuleID:      4,
		ProductCode: "OM1",
		Operator:    "ge",
		Qty:         1,
	}
	result4 := models.Result{
		ID:               4,
		RuleID:           4,
		ProductCode:      "AP1",
		Qty:              1,
		ResultType:       "applyPrice",
		AppliedPrice:     50,
		AppliedPriceType: "percent",
	}
	rule4 := models.Rule{
		ID:     4,
		Rtype:  "applyPrice",
		Name:   "APOM",
		Action: action4,
		Result: result4,
	}
	models.ActionsDBMap[4] = action4
	models.ResultsDBMap[4] = result4
	models.RulesDBMap[4] = rule4
}
