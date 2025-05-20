package model

type BaseCluster struct {
	Domain		string
	Name           string
	Platform       []string
	SshKey         string
	Node		*Node
	Network		*Network
	Image		*Image
	Storage        string
	Iam            string
	Cursor         int
	Selected       map[int]struct{}

}

type AdvancedSettings struct {
	HyperThreading bool
}

type Image struct {
	Image      string
	PullSecret string
}

type Network struct {
	Cidr string
	NetworkType string
	ApiVIP string
	IngressVIP string
	Region         string
}

type Node struct {
	MasterName string
	MasterCount uint
	MasterSize string
	WorkerName string
	WorkerCount uint
	WorkerSize string
}

