# Weather-Monster


Dependencies

Golang 1.16 or higher

Docker for Mac/Windows

Steps to run the project 

    1. Create a base folder and clone the projects github.com/shyam81992/Site-Information.
    2. Run go mod tidy to install the dependencies.
    
    3. To run the testcases run the command . ../config/dev.sh && go test ./... -v

Note: If you are using windows please install git bash so that you can run the shell script  ". config/dev.sh"


Post /static/siteinfo -b {url : "string"} 
status 200
response {
    "Page Title":"pagetitle",
    "Headings":{"h1": 10},
    "Internal Links":10",
    "External Links":     10,
	"Inaccessible Links": 10,
	"loginform": true
}

Post /dynamic/siteinfo -b {url : "string", email : "string"}  status 200 
response {"message" : "ok"}

/dynamic/siteinfo 
Follows async model. Since to scrap the dynamic page we need to use chromedp or go-selenium and both the packages will be using chrome in background (If we implement this logic in the Site-Information service, we might face performace issues when the number of requests increases). A seperate microservice is created for this (Site-Information-Job).

Site-Information-Job (Work in progess)
Used to scrap the dynamic page using chromedp or go-selenium and notify the website information to the user through email.

Sample selenium code under automation/selenium.go_

go-selenium supports language translation where chromedp doesn't. 

Area of Improvement:
Need to verify if go-selenium language translation is stable. If not then we need to use chromdp + Internationalization or chromdp + third party api which tranlates (ex google translate or deepL).


