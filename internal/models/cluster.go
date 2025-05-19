package model

type BaseCluster struct {
	domain		string
	name           string
	platform       []string
	sshKey         string
	node		*Node
	network		*Network
	image		*Image
	storage        string
	iam            string
	cursor         int
	selected       map[int]struct{}
	keyPrompt string
}






type AdvancedSettings struct {
	hyperThreading bool
}
