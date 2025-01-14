## Dependency Risk Analysis Tool - DRAT

DRAT aims to provide risk indicator for libraries used by the developer within an organisation.

Below is what the result looks like after running against a python project. Note that these are just potential risks associate with these dependencies being identified and not actually vulnerabilities.

![drat-cli gui image](https://github.com/GovAuCSU/DRAT/raw/master/images/drat_cli_gui.png)



## Known problems:

 - Duplicate results
 - No caching for previous results 
 - Difficult to map relationships between result entries
 - Rate limiting force the use of personal token


## DRAT-cli
The app can be found under cmd/drat-cli. 

To run the app, you will need a github auth token to overcome the rate limit for github api. This token can be generated [here](https://github.com/settings/tokens). Unless you need private repository access for your project, a token without any permission should work fine.

To check out how it works, you can try running it with the following commands

To get a report for a single repository:
```bash
cd cmd/drat-cli
go run main.go -d 2 -v -r "https://github.com/AusDTO/dto-digitalmarketplace-supplier-frontend"
```

To get a report for a list of repository:
```bash
cd cmd/drat-cli
go run main.go -d 5 -v -f examplelist.lst -o output.json
```

Below is an example of the app's output. You can then open `cmd/drat-cli/static/index.html` in a browser and drag&drop `output.json` into the site to visualise the output and navigate around.

```json
[
    {
        "Ownername": "santrancisco",
        "Name": "cque",
        "URL": "https://api.github.com/repos/santrancisco/cque",
        "Dependencies": null,
        "RiskNotes": [
            "[INFO] Repository is not managed under an organisation",
            "[MEDIUM] Size of collaborator for the repository is 1 which is less than 3",
            "[MEDIUM] Repository has only been created for less than 30 days",
            "[GOOD] Repository has a wiki"
        ]
    },
    {
        "Ownername": "google",
        "Name": "go-github",
        "URL": "https://api.github.com/repos/google/go-github",
        "Dependencies": null,
        "RiskNotes": [
            "[MEDIUM] Repository has only been created for less than 30 days",
            "[GOOD] Repository has been forked 100 times",
            "[GOOD] Repository has been stared 3762 times",
            "[GOOD] Repository is being watched by 3762 people"
        ]
    },
    {
        "Ownername": "GovAuCSU",
        "Name": "certwatch",
        "URL": "https://api.github.com/repos/GovAuCSU/certwatch",
        "Dependencies": [
            "github.com/beorn7/perks",
            "github.com/bgentry/que-go",
            "github.com/cloudfoundry-community/go-cfenv",
            "github.com/gogo/protobuf",
            "github.com/golang/protobuf",
            "github.com/google/certificate-transparency-go",
            "github.com/jackc/pgx",
            "github.com/matttproud/golang_protobuf_extensions",
            "github.com/mitchellh/mapstructure",
            "github.com/pkg/errors",
            "github.com/prometheus/client_golang",
            "github.com/prometheus/client_model",
            "github.com/prometheus/common",
            "github.com/prometheus/procfs",
            "golang.org/x/crypto",
            "golang.org/x/net"
        ],
        "RiskNotes": [
            "[MEDIUM] Size of collaborator for the repository is 1 which is less than 3",
            "[MEDIUM] Repository has only been created for less than 30 days",
            "[GOOD] Repository has a wiki"
        ]
    },
    {
        "Ownername": "beorn7",
        "Name": "perks",
        "URL": "https://api.github.com/repos/beorn7/perks",
        "Dependencies": null,
        "RiskNotes": [
            "[INFO] Repository is not managed under an organisation",
            "[MEDIUM] Repository has only been created for less than 30 days",
            "[GOOD] Repository has a wiki"
        ]
    }
]
```

This app currently support 3 types of crawler plugins for github projects: NPM, go-dep, pip  and ruby-gem


## DRAT Webapp
Development for this app is currently on hold in favour of the drat-cli tool.

This tool's intial code and design is inspired by the certificate transparency work done by Adam Eijdenberg from Cloud.gov.au team. You can find the repository for that [here](https://github.com/govau/certwatch/tree/master/jobs)

#### Running locally

```bash
docker run -p 5432:5432 --name dratpg -e POSTGRES_USER=dratpg -e POSTGRES_PASSWORD=dratpg -d postgres
export VCAP_SERVICES='{"postgres": [{"credentials": {"username": "certpg", "host": "localhost", "password": "certpg", "name": "certpg", "port": 5434}, "tags": ["postgres"]}]}'
go run *.go
```

#### Running in docker with latest code

```bash
GOOS=linux GOARCH=amd64 go build -o bin/drat cmd/drat/main.go
docker-compose up
```

To checkout database:

```bash
psql "dbname=dratpg host=localhost user=dratpg password=dratpg port=5432"
```
