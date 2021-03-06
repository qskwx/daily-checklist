package iface

import "daily-checklist/internal/session"

func constructMessage(ss session.Session) string {
	message := ""
	for _, category := range ss.Categories() {
		message += category.Name() + ":\n"
		for _, activity := range category.Activities() {
			message += "\t\t\t\t"
			var doneSymbol string
			if activity.Done() {
				doneSymbol = "+"
			} else {
				doneSymbol = "-"
			}
			message += doneSymbol + activity.Id() + doneSymbol + " " + activity.Activity() + "\n"
		}
	}
	return message
}
