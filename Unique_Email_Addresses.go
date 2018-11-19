// Runtime: 12 ms, faster than 97.30% of Go online submissions for Unique Email Addresses.
func numUniqueEmails(emails []string) int {
	m := make(map[string]struct{}, len(emails))
	for _, email := range emails {
		split := strings.Split(email, "@")
		name, domain := strings.Split(split[0], "+")[0], split[1]
		name = strings.Replace(name, ".", "", -1)
		new := name + "@" + domain
		if _, ok := m[new]; ok {
			continue
		}
		m[new] = struct{}{}
	}
	return len(m)
}