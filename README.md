# Go-Veeam-Auth

This is a lightweight Go package makes it easier to start working with Veeam APIs with Go.

The package implements all Veeam's products API OAuth based authorization and includes some convenience functions that make is easier to make requests.

There are no plans to make this a full SDK, there are just too many structs and calls across all the products for that to be practical.

NOTE: This is an open source project with no association with Veeam and is supplied under the MIT license.

### Add to your project

    go get github.com/shapedthought/go-veeam-auth

### Example of use:

    import "github.com/shapedthought/go-veeam-auth"

    func main() {

        // Get the profile required
        profile := gva.GetProfile("vbr")

        // Set the IP and username
        address := "192.168.0.106"
        userName := "administrator@yourdomain.co.uk"

        // Skip TLS verification, this is required if you are using self-signed certificates
        tr := &http.Transport{
            TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        }

        // Create the HTTP client
        client := &http.Client{Transport: tr}

        // call the ApiLogin method to get the token
        token := gva.ApiLogin(client, profile, address, userName)

        // use the BuildRequestUrl to construct the URL string for the request you want to make
        endPoint := "jobs"
        cs := gva.BuildRequestUrl(address, endPoint, profile)

        // Create your request
        r, err := http.NewRequest("GET", cs, nil)
        if err != nil {
            log.Fatal(err)
        }

        // Use the AddHeaders method to have gva add the headers for you
        gva.AddHeaders(r, profile, token)

        res, err := client.Do(r)
        if err != nil {
            log.Fatal(err)
        }

        defer res.Body.Close()

        body, _ := io.ReadAll(res.Body)

        fmt.Println(res)
        fmt.Println(string(body))
    }

### Profiles

First you need to get the profile you want to use, the options are:

- vbaws - Veeam Backup for AWS
- vbr - Veeam Backup and Replication
- vbaz - Veeam Backup for Azure
- vbgcp - Veeam Backup for GCP
- vone - Veeam ONE
- vb365 - Veeam Backup for M365

This will return all the elements required for the login for that API.

Next you will need to create a HTTP client using the usual method.

Finally need to use the ApiLogin method passing in:

- A pointer to a HTTP Client
- The Profile struct
- The API address as a string
- The Username

You'll noticed that there's no password specified here, well it's never a great idea to hardcode passwords so you will need to set it in the VEEAM_API_PASSWORD environmental variable:

Example:

    $env:VEEAM_API_PASSWORD = "yourPassword"

### Build your request URL

You can use the BuildRequestUrl method to construct the URL that you want to call.

    endPoint := "jobs"
    cs := gva.BuildRequestUrl(a, endPoint, profile)

The BuildRequestUrl takes:

- The API address (no need for http:// at the start or the port after)
- The endpoint
- The profile struct

The most important part of this is the endPoint variable, you only need the last part of the endpoint you wish to call.

For example to get VBR Jobs you would normally call:

    /api/v1/jobs

In this case you only need to supply the last element

    jobs

### Adding the Headers to your request

You will need to set up a request to send the calls you want to make, but to make life a bit easier you can add the required headers by calling the AddHeaders method.

    // create your request
    r, err := http.NewRequest("GET", cs, nil)
    if err != nil {
        log.Fatal(err)
    }

    // use the AddHeaders method to add the headers for you
    gva.AddHeaders(r, profile, token)

The function takes the Request pointer, the Profile struct and the Token struct.

It will automatically set up everything including the Bearer token.

## Making changes

We will aim to keep this package up-to-date with the latest Veeam releases, but if we missing something or you need to make a change, this is how you do it.

### API Versions

The current implementation uses the following API versions:

| Product | Version | x-api-version |
| ------- | ------- | ------------- |
| VBAWS   | v1      | 1.2-rev0      |
| VBR     | v1      | 1.0-rev2      |
| VBAZ    | v3      |               |
| VBGCP   | v1      | 1.0-rev0      |
| VONE    | v2      | 1.0-rev2      |
| VB365   | v6      |               |

If you need to update the x-api-version due to a minor update it it is recommended to do it manually.

    profile := gva.GetProfile("vbr")

    profile.Headers.XAPIVersion = "1.0-rev3"

If it is a major release then use the UpdateVersion method

    profile := gva.GetProfile("vbr")

    profile.UpdateVersion("v2", "2.0-rev1")

The method takes the version number as the first argument, and the x-api-version as the second.

### Port

The ports used are:

| Product | Port  |
| ------- | ----- |
| VBAWS   | 11005 |
| VBR     | 9419  |
| VBAZ    |       |
| VBGCP   | 13140 |
| VONE    | 1239  |
| VB365   | 4443  |

If you need to update the port run the UpdatePort method on the Profile struct

    profile := gva.GetProfile("vbr")

    profile.updatePort("1234")

## Issues and Contribution

If you have any issues with this package please raise an issue.

If you would like to contribute this project please feel free to send a pull request.
