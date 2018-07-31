package score

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/GovAuCSU/DRAT/crawl"
)

// func setupTest() (r *github.Repository, u *github.User, c *github.Client, err error) {
// 	err = nil
// 	ctx := context.Background()
// 	c = github.NewClient(&http.Client{})
// 	if os.Getenv("GITHUB_AUTH_TOKEN") != "" {
// 		ts := oauth2.StaticTokenSource(
// 			&oauth2.Token{AccessToken: os.Getenv("GITHUB_AUTH_TOKEN")},
// 		)

// 		tc := oauth2.NewClient(ctx, ts)
// 		c = github.NewClient(tc)
// 	}
// 	r, _, err = c.Repositories.Get(ctx, "google", "go-github")
// 	if err != nil {
// 		return nil, nil, nil, err
// 	}
// 	u, _, err = c.Users.Get(ctx, "santrancisco")
// 	if err != nil {
// 		return nil, nil, nil, err
// 	}
// 	return r, u, c, err
// }

// Note: These are not proper unit test, these are being used to test run the function without having to run the whole app at the moment. These will
// be converted into proper test in the future.
func TestParseGem(t *testing.T) {
	t.Log("Test Parse")
	content, err := ioutil.ReadFile("Gemfile")
	if err != nil {
		t.Errorf("Cannot open Gemfile")
	}
	lst := crawl.ParseGem(content)
	final := strings.Join(lst, ",")
	fmt.Println(final)
	// expected := "github.com/beorn7/perks,github.com/bgentry/que-go,github.com/cloudfoundry-community/go-cfenv,github.com/gogo/protobuf,github.com/golang/protobuf,github.com/google/certificate-transparency-go,github.com/jackc/pgx,github.com/matttproud/golang_protobuf_extensions,github.com/mitchellh/mapstructure,github.com/pkg/errors,github.com/prometheus/client_golang,github.com/prometheus/client_model,github.com/prometheus/common,github.com/prometheus/procfs,golang.org/x/crypto,golang.org/x/net"
	// if final != expected {
	// 	t.Errorf("ParseGoDep fails, output is not as expected")
	// }
}