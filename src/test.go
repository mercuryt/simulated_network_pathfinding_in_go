package main

type ServerTestData struct {
	name      string
	latitude  float64
	longitude float64
}

func main() {
	servers := []ServerTestData{
		ServerTestData{"New York", 40.7834300, -73.9662500},
		ServerTestData{"Paris", 48.8534100, 2.3488000},
		ServerTestData{"London", 51.5085300, -0.1257400},
		ServerTestData{"Frankfurt", 50.122008, 8.690800},
		ServerTestData{"Rome", 41.894143, 12.479683},
		ServerTestData{"Moscow", 55.716034, 37.640622},
		ServerTestData{"Hong Kong", 22.337174, 114.127436},
		ServerTestData{"Beijing", 39.894307, 116.336993},
		ServerTestData{"Tokyo", 35.699474, 139.722737},
		ServerTestData{"Sydney", -33.869639, 151.126925},
		ServerTestData{"Auckland", -36.883566, 174.748696},
		ServerTestData{"Bangkok", 13.697759, 100.494355},
		ServerTestData{"Los Angeles", 34.024732, -118.332127},
		ServerTestData{"Rio De Janeiro", -22.906953, -43.203326},
		ServerTestData{"Honalulu", 21.327975, -157.830672},
		ServerTestData{"Singapore", 1.344245, 103.803059},
	}

	serverNodeMap := newServerNodeMap()
	for index, std := range servers {
		serverNodeMap.addServerNode(index+1, std.name, std.latitude, std.longitude)
	}

	serverNodeMap.setLatency("Moscow", "Beijing", 1000)
	serverNodeMap.printNode("Moscow")
	serverNodeMap.printNode("Beijing")

	path := serverNodeMap.getPath("New York", "Singapore")
	path.print()
}
