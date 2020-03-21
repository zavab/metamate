package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/metamatex/metamate/hackernews-svc/gen/v0/sdk"
	"net/http"
)

type user struct {
	Id        *string
	Created   *int
	Delay     *int
	Karma     *int
	About     *string
	Submitted []int
}

func getSocialAccountId(c *http.Client, req sdk.GetSocialAccountsRequest) (as []sdk.SocialAccount, errs []error) {
	err := func() (err error) {
		var url string

		switch *req.Mode.Id.Kind {
		case sdk.IdKind.ServiceId:
			url = fmt.Sprintf("https://hacker-news.firebaseio.com/v0/user/%v.json", *req.Mode.Id.ServiceId.Value)
		case sdk.IdKind.Username:
			url = fmt.Sprintf("https://hacker-news.firebaseio.com/v0/user/%v.json", *req.Mode.Id.Username)
		default:
			err = errors.New(fmt.Sprintf("can't handle id %v", req.Mode.Id))

			return
		}

		rsp, err := c.Get(url)
		if err != nil {
			return
		}

		u := user{}
		err = json.NewDecoder(rsp.Body).Decode(&u)
		if err != nil {
			return
		}

		as = append(as, mapUserToSocialAccount(u))

		return
	}()
	if err != nil {
		errs = append(errs, err)
	}

	return
}

func mapUserToSocialAccount(u user) (a sdk.SocialAccount) {
	//type user struct {
	//	Id        string    x
	//	Created   int       x
	//	Delay     int
	//	Karma     int
	//	About     string    x
	//	Submitted []int     x
	//}

	a = sdk.SocialAccount{
		Id: &sdk.ServiceId{
			Value: u.Id,
		},
		AlternativeIds: []sdk.Id{
			{
				Kind:     &sdk.IdKind.Username,
				Username: u.Id,
			},
			{
				Kind: &sdk.IdKind.Url,
				Url: &sdk.Url{
					Value: sdk.String(fmt.Sprintf("https://news.ycombinator.com/user?id=%v", *u.Id)),
				},
			},
		},
		Note: &sdk.Text{
			Formatting: &sdk.FormattingKind.Html,
			Value:      u.About,
		},
		Meta: &sdk.TypeMeta{
			CreatedAt: &sdk.Timestamp{
				Kind: &sdk.TimestampKind.Unix,
				Value: &sdk.DurationScalar{
					Unit: &sdk.DurationUnit.S,
					Value: sdk.Float64(float64(*u.Created)),
				},
			},
		},
		Relations: &sdk.SocialAccountRelations{
			AuthorsStatuses: &sdk.StatusesCollection{
				Meta: &sdk.CollectionMeta{
					Count: sdk.Int32(int32(len(u.Submitted))),
				},
				Statuses: func() (ss []sdk.Status) {
					for _, s := range u.Submitted {
						ss = append(ss, sdk.Status{
							Id: &sdk.ServiceId{
								Value: sdk.String(fmt.Sprintf("%v", s)),
							},
						})
					}

					return
				}(),
			},
		},
	}

	return a
}
