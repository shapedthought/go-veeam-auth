package gva

func GetProfile(profile string) Profile {

	vb365 := Profile{
		Name: "vb365",
		Headers: Headers{
			Accept:      "application/json",
			ContentType: "application/x-www-form-urlencoded",
			XAPIVersion: "",
		},
		URL:        ":4443/v6/Token",
		Port:       "4443",
		APIVersion: "v6",
	}

	aws := Profile{
		Name: "aws",
		Headers: Headers{
			Accept:      "application/json",
			ContentType: "application/x-www-form-urlencoded",
			XAPIVersion: "1.2-rev0",
		},
		URL:        ":11005/api/v1/token",
		Port:       "11005",
		APIVersion: "v1",
	}

	vbr := Profile{
		Name: "vbr",
		Headers: Headers{
			Accept:      "application/json",
			ContentType: "application/x-www-form-urlencoded",
			XAPIVersion: "1.0-rev2",
		},
		URL:        ":9419/api/oauth2/token",
		Port:       "9419",
		APIVersion: "v1",
	}

	azure := Profile{
		Name: "azure",
		Headers: Headers{
			Accept:      "application/json",
			ContentType: "application/x-www-form-urlencoded",
			XAPIVersion: "",
		},
		URL:        "/api/oauth2/token",
		Port:       "",
		APIVersion: "v3",
	}

	gcp := Profile{
		Name: "gcp",
		Headers: Headers{
			Accept:      "application/json",
			ContentType: "application/x-www-form-urlencoded",
			XAPIVersion: "1.0-rev0",
		},
		URL:        ":13140/api/v1/token",
		Port:       "13140",
		APIVersion: "v1",
	}

	vone := Profile{
		Name: "vone",
		Headers: Headers{
			Accept:      "application/json",
			ContentType: "application/x-www-form-urlencoded",
			XAPIVersion: "1.0-rev2",
		},
		URL:        ":1239/api/token",
		Port:       ":1239",
		APIVersion: "v2",
	}

	var selectedProfile Profile

	switch profile {
	case "vbaws":
		selectedProfile = aws
	case "vbr":
		selectedProfile = vbr
	case "vbaz":
		selectedProfile = azure
	case "vbgcp":
		selectedProfile = gcp
	case "vone":
		selectedProfile = vone
	case "vb365":
		selectedProfile = vb365
	}

	return selectedProfile
}