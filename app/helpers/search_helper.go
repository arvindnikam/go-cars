package helpers

import (
	"regexp"
	"strings"

	"github.com/spf13/cast"
)

var (
	SEARCH_OPTIONS = []string{"limit", "offset", "sort_column", "sort_order"}
)

func GetSearchOptions(request map[string]interface{}) map[string]interface{} {
	default_options := map[string]interface{}{
		"limit":       10,
		"offset":      nil,
		"sort_column": "updated_at",
		"sort_order":  "desc",
	}

	options := map[string]interface{}{}
	for _, key := range SEARCH_OPTIONS {
		if value, present := request[key]; present {
			options[key] = value
		} else {
			options[key] = default_options[key]
		}
	}

	if _, present := options["offset"]; !present {
		pag := options["page"]
		lim := options["limit"]
		// page := cast.ToInt(options["page"])
		// limit := cast.ToInt(options["limit"])

		var page int
		if pag == nil || pag == 0 {
			page = 1
		} else {
			page = cast.ToInt(pag)
		}

		var limit int
		if lim == nil || lim == 0 {
			limit = 10
		} else {
			limit = cast.ToInt(lim)
		}
		options["offset"] = (page - 1) * limit
	}

	return options
}

func ParseConditions(conditions map[string]map[string]interface{}) string {
	searchQueryArray := []string{}

	for uid, operator_value_hash := range conditions {
		uid = underscore(uid)
		operator := operator_value_hash["operator"]
		val := operator_value_hash["value"]
		value := cast.ToString(val)

		switch operator {
		case "eq":
			searchQueryArray = append(searchQueryArray, uid+" = '"+value+"'")
		case "not":
			searchQueryArray = append(searchQueryArray, uid+" != '"+value+"'")
		case "is_empty":
			searchQueryArray = append(searchQueryArray, uid+" IS NULL")
		case "not_empty":
			searchQueryArray = append(searchQueryArray, uid+" IS NOT NULL")
		case "lte":
			searchQueryArray = append(searchQueryArray, uid+" <= '"+value+"'")
		case "gte":
			searchQueryArray = append(searchQueryArray, uid+" >= '"+value+"'")
		case "lt":
			searchQueryArray = append(searchQueryArray, uid+" < '"+value+"'")
		case "gt":
			searchQueryArray = append(searchQueryArray, uid+" > '"+value+"'")
		default:
			searchQueryArray = append(searchQueryArray, uid+" = '"+value+"'")
		}
	}

	searchQuery := strings.Join(searchQueryArray, " AND ")
	return searchQuery
}

func underscore(word string) string {
	re := regexp.MustCompile("[A-Z]")
	word = re.ReplaceAllStringFunc(word, func(s string) string {
		return "_" + strings.ToLower(s)
	})

	word = strings.Trim(word, "_")
	return word
}
