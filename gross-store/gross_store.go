package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{"quarter_of_a_dozen": 3, "half_of_a_dozen": 6, "dozen": 12,
		"small_gross": 120, "gross": 144, "great_gross": 1728}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitsValue, okUnit := units[unit]
	itemsValue, okItem := bill[item]
	if !okUnit {
		return false
	} else if !okItem {
		bill[item] = unitsValue
		return true
	} else {
		bill[item] = unitsValue + itemsValue
		return true
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unitsValue, okUnit := units[unit]
	itemsValue, okItem := bill[item]
	if !okUnit || !okItem {
		return false
	} else if itemsValue < unitsValue {
		return false
	} else if itemsValue == unitsValue {
		delete(bill, item)
		return true
	} else {
		bill[item] = itemsValue - unitsValue
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	v, ok := bill[item]
	return v, ok
}
