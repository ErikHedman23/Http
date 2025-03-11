package main

func getMailtoLinkForEmail(email string) string {
	if len(email) == 0 {
		return "mailto:"
	}
	return "mailto:" + email
}
