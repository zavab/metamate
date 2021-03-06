package boot_test

import (
	"context"
	"fmt"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamate/gen/v0/sdk"
	"github.com/metamatex/metamate/gen/v0/sdk/utils"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/metamate/pkg/v0/boot"
	"github.com/metamatex/metamate/metamate/pkg/v0/business/line"
	"github.com/metamatex/metamate/metamate/pkg/v0/business/virtual"
	"github.com/metamatex/metamate/metamate/pkg/v0/config"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

const (
	AuthSvc     = "auth"
	PipeSvc     = "pipe"
	ReqFilter   = "request-filter"
	SqlxA       = "sqlx-a"
	SqlxB       = "sqlx-b"
	ErrorA      = "error-a"
	ErrorB      = "error-b"
	PaginationA = "pagination-a"
	PaginationB = "pagination-b"
)

func TestBoot(t *testing.T) {
	c := config.DefaultConfig

	c.Log.Internal = types.InternalLogConfig{
		config.SvcRsp: map[string]string{
			//sdk.GetPostsResponseName: "{{ .Ctx.Svc.Url.Value }} : \n{{ .Ctx.GSvcReq.Type.Name }}\n{{ .Ctx.GSvcReq.Sprint }}\n\n{{ .Ctx.GSvcRsp.Type.Name }}\n{{ .Ctx.GSvcRsp.Sprint }}",
			//"*": "{{ .Ctx.Svc.Url.Value }} : {{ .Ctx.GSvcRsp.Type.Name }}",
		},
	}

	c.Virtual.Services = append(c.Virtual.Services, []types.VirtualSvc{
		{
			Id:   ErrorA,
			Name: virtual.Error,
		},
		{
			Id:   ErrorB,
			Name: virtual.Error,
		},
		{
			Id:   PaginationA,
			Name: virtual.Pagination,
		},
		{
			Id:   PaginationB,
			Name: virtual.Pagination,
		},
	}...)

	d, err := boot.NewDependencies(c, types.Version{})
	if err != nil {
		log.Fatal(err)
	}

	d.ResolveLine.(*line.Line).SetLog(false)

	ctx := context.Background()

	suffix := "-" + time.Now().Format("3:04:05PM")
	println(suffix)

	f := func(ctx context.Context, gCliReq generic.Generic) (gCliRsp generic.Generic, err error) {
		gCliReq.MustSetGeneric([]string{fieldnames.Select, fieldnames.Errors}, d.Factory.MustFromStruct(sdk.ErrorSelect{
			Message: sdk.Bool(true),
		}))

		gCliRsp = d.ServeFunc(ctx, gCliReq)

		return
	}

	//spec.TestRequestFilter(t, ctx, d.Factory, f)
	//
	//spec.TestPipe(t, ctx, d.Factory, f)

	//spec.TestDiscovery(t, ctx, d.Factory, f)

	//FTestHackernewsGetPostFeedContainsPosts(t, ctx, d.Factory, f)

	//FTestError(t, ctx, d.Factory, f)

	//FTestHackernews(t, ctx, d.Factory, f)

	FTestPagination(t, ctx, d.Factory, f)

	FTestHackernews(t, ctx, d.Factory, f)

	//FTestHackernews(t, ctx, d.Factory, f)

	//FTestHackernewsSocialAccount(t, ctx, d.Factory, f)

	//spec.TestEmptyPost(t, ctx, d.Factory, f)
	//
	//spec.TestEmptyGet(t, ctx, d.Factory, f)
	//
	//spec.TestPost(t, ctx, d.Factory, f, SqlxA)
	//
	//spec.TestPostWithNameId(t, ctx, d.Factory, f, SqlxA, suffix)
	//
	//spec.TestGetModeId(t, ctx, d.Factory, f, SqlxA, suffix)
	//
	//spec.TestGetModeCollection(t, ctx, d.Factory, f, SqlxA)
	//
	//spec.TestFilterStringIs(t, ctx, d.Factory, f, SqlxA, suffix)
	//
	//spec.TestGetModeRelation(t, ctx, d.Factory, f, SqlxA, suffix)
	//
	//spec.TestGetModeIdWithNameId(t, ctx, d.Factory, f, suffix, SqlxA)
	//
	//spec.TestGetModeIdWithServiceFilter(t, ctx, d.Factory, f, suffix, SqlxA, SqlxB)
	//
	//spec.TestPostClientAccounts(t, ctx, d.Factory, f, SqlxA)
	//
	//spec.TestAuthenticateClientAccount(t, ctx, d.Factory, f, SqlxA, AuthSvc)
	//
	//spec.TestToken(t, ctx, d.Factory, f, AuthSvc, SqlxA)
	//
	//spec.TestGetModeRelationInter(t, ctx, d.Factory, f, SqlxA, SqlxB, suffix)
	//
	//spec.TestGetModeIdWithSelfReferencingRelation(t, ctx, d.Factory, f, suffix, SqlxA)
	//
	//spec.TestGetModeIdWithRelation(t, ctx, d.Factory, f, suffix, SqlxA)
}

func FTestPagination(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	name := "TestPagination"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		t.Run("returnsPaginations", func(t *testing.T) {
			t.Parallel()

			tcs := []struct {
				mode *sdk.GetMode
			}{
				{
					mode: &sdk.GetMode{
						Kind:       &sdk.GetModeKind.Collection,
						Collection: &sdk.CollectionGetMode{},
					},
				},
				{
					mode: &sdk.GetMode{
						Kind:     &sdk.GetModeKind.Relation,
						Relation: &sdk.RelationGetMode{},
					},
				},
				{
					mode: &sdk.GetMode{
						Kind:   &sdk.GetModeKind.Search,
						Search: &sdk.SearchGetMode{},
					},
				},
			}

			for _, tc := range tcs {
				t.Run(*tc.mode.Kind, func(t *testing.T) {
					err := func() (err error) {
						getReq := sdk.GetWhateversRequest{
							ServiceFilter: &sdk.ServiceFilter{
								Id: &sdk.ServiceIdFilter{
									Value: &sdk.StringFilter{
										In: []string{PaginationA},
									},
								},
							},
							Mode: tc.mode,
						}

						gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
						if err != nil {
							return
						}

						_, ok := gGetRsp.Generic(fieldnames.Pagination)
						assert.True(t, ok)

						return
					}()
					if err != nil {
						t.Error(err)
					}
				})
			}
		})

		t.Run("addsServiceIds", func(t *testing.T) {
			t.Parallel()

			tcs := []struct {
				mode *sdk.GetMode
			}{
				{
					mode: &sdk.GetMode{
						Kind:       &sdk.GetModeKind.Collection,
						Collection: &sdk.CollectionGetMode{},
					},
				},
				{
					mode: &sdk.GetMode{
						Kind:     &sdk.GetModeKind.Relation,
						Relation: &sdk.RelationGetMode{},
					},
				},
				{
					mode: &sdk.GetMode{
						Kind:   &sdk.GetModeKind.Search,
						Search: &sdk.SearchGetMode{},
					},
				},
			}

			for _, tc := range tcs {
				t.Run(*tc.mode.Kind, func(t *testing.T) {
					err := func() (err error) {
						getReq := sdk.GetWhateversRequest{
							ServiceFilter: &sdk.ServiceFilter{
								Id: &sdk.ServiceIdFilter{
									Value: &sdk.StringFilter{
										In: []string{PaginationA},
									},
								},
							},
							Mode: tc.mode,
						}

						gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
						if err != nil {
							return
						}

						gPagination, ok := gGetRsp.Generic(fieldnames.Pagination)
						assert.True(t, ok)

						var p sdk.Pagination
						gPagination.MustToStruct(&p)

						for _, sp := range append(p.Previous, append(p.Current, p.Next...)...) {
							assert.NotNil(t, sp.Id)
						}

						return
					}()
					if err != nil {
						t.Error(err)
					}
				})
			}
		})

		t.Run("mergesPaginations", func(t *testing.T) {
			t.Parallel()

			tcs := []struct {
				mode *sdk.GetMode
			}{
				{
					mode: &sdk.GetMode{
						Kind:       &sdk.GetModeKind.Collection,
						Collection: &sdk.CollectionGetMode{},
					},
				},
				{
					mode: &sdk.GetMode{
						Kind:   &sdk.GetModeKind.Search,
						Search: &sdk.SearchGetMode{},
					},
				},
			}

			for _, tc := range tcs {
				t.Run(*tc.mode.Kind, func(t *testing.T) {
					err := func() (err error) {
						getReq := sdk.GetWhateversRequest{
							ServiceFilter: &sdk.ServiceFilter{
								Id: &sdk.ServiceIdFilter{
									Value: &sdk.StringFilter{
										In: []string{PaginationA, PaginationB},
									},
								},
							},
							Mode: tc.mode,
						}

						gGetReq := f.MustFromStruct(getReq)

						gGetRsp, err := h(ctx, gGetReq)
						if err != nil {
							return
						}

						gPagination, ok := gGetRsp.Generic(fieldnames.Pagination)
						assert.True(t, ok)

						var p sdk.Pagination
						gPagination.MustToStruct(&p)

						assert.Len(t, p.Previous, 0, fmt.Sprintf("%v:\n%v\n%v:\n%v", gGetReq.Type().Name(), gGetReq.Sprint(), gGetRsp.Type().Name(), gGetRsp.Sprint()))
						assert.Len(t, p.Current, 2, fmt.Sprintf("%v:\n%v\n%v:\n%v", gGetReq.Type().Name(), gGetReq.Sprint(), gGetRsp.Type().Name(), gGetRsp.Sprint()))
						assert.Len(t, p.Next, 2, fmt.Sprintf("%v:\n%v\n%v:\n%v", gGetReq.Type().Name(), gGetReq.Sprint(), gGetRsp.Type().Name(), gGetRsp.Sprint()))

						return
					}()
					if err != nil {
						t.Error(err)
					}
				})
			}
		})

		t.Run("distributesPages", func(t *testing.T) {
			tcs := []struct {
				mode *sdk.GetMode
			}{
				{
					mode: &sdk.GetMode{
						Kind:       &sdk.GetModeKind.Collection,
						Collection: &sdk.CollectionGetMode{},
					},
				},
				{
					mode: &sdk.GetMode{
						Kind:   &sdk.GetModeKind.Search,
						Search: &sdk.SearchGetMode{},
					},
				},
			}

			for _, tc := range tcs {
				t.Run(*tc.mode.Kind, func(t *testing.T) {
					err := func() (err error) {
						getReq := sdk.GetWhateversRequest{
							ServiceFilter: &sdk.ServiceFilter{
								Id: &sdk.ServiceIdFilter{
									Value: &sdk.StringFilter{
										In: []string{PaginationA, PaginationB},
									},
								},
							},
							Mode: tc.mode,
							Pages: []sdk.ServicePage{
								{
									Id: &sdk.ServiceId{
										ServiceName: sdk.String("discovery"),
										Value:       sdk.String(PaginationA),
									},
									Page: &sdk.Page{
										Kind: &sdk.PageKind.IndexPage,
										IndexPage: &sdk.IndexPage{
											Value: sdk.Int32(0),
										},
									},
								},
								{
									Id: &sdk.ServiceId{
										ServiceName: sdk.String("discovery"),
										Value:       sdk.String(PaginationB),
									},
									Page: &sdk.Page{
										Kind: &sdk.PageKind.IndexPage,
										IndexPage: &sdk.IndexPage{
											Value: sdk.Int32(1),
										},
									},
								},
							},
						}

						gGetReq := f.MustFromStruct(getReq)

						gGetRsp, err := h(ctx, gGetReq)
						if err != nil {
							return
						}

						gPagination, ok := gGetRsp.Generic(fieldnames.Pagination)
						assert.True(t, ok)

						var p sdk.Pagination
						gPagination.MustToStruct(&p)

						assert.ElementsMatch(t, p.Current, getReq.Pages, fmt.Sprintf("%v:\n%v\n%v:\n%v", gGetReq.Type().Name(), gGetReq.Sprint(), gGetRsp.Type().Name(), gGetRsp.Sprint()))

						assert.Len(t, p.Previous, 1, fmt.Sprintf("%v:\n%v\n%v:\n%v", gGetReq.Type().Name(), gGetReq.Sprint(), gGetRsp.Type().Name(), gGetRsp.Sprint()))
						assert.Len(t, p.Current, 2, fmt.Sprintf("%v:\n%v\n%v:\n%v", gGetReq.Type().Name(), gGetReq.Sprint(), gGetRsp.Type().Name(), gGetRsp.Sprint()))
						assert.Len(t, p.Next, 2, fmt.Sprintf("%v:\n%v\n%v:\n%v", gGetReq.Type().Name(), gGetReq.Sprint(), gGetRsp.Type().Name(), gGetRsp.Sprint()))

						return
					}()
					if err != nil {
						t.Error(err)
					}
				})
			}
		})
	})
}

func FTestError(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	name := "TestError"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			getReq := sdk.GetWhateversRequest{
				ServiceFilter: &sdk.ServiceFilter{
					Id: &sdk.ServiceIdFilter{
						Value: &sdk.StringFilter{
							In: []string{ErrorA, ErrorB},
						},
					},
				},
				Select: &sdk.GetWhateversResponseSelect{
					Errors: &sdk.ErrorSelect{
						Message: sdk.Bool(true),
					},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			gGetRsp.Print()

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}

func FTestHackernews(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	t.Run("GetSocialAccounts", func(t *testing.T) {
		t.Run("id", func(t *testing.T) {
			err := func() (err error) {
				getReq := sdk.GetSocialAccountsRequest{
					Mode: &sdk.GetMode{
						Kind: &sdk.GetModeKind.Id,
						Id: &sdk.Id{
							Kind: &sdk.IdKind.ServiceId,
							ServiceId: &sdk.ServiceId{
								Value:       sdk.String("21stio"),
								ServiceName: sdk.String("hackernews"),
							},
						},
					},
				}

				gGetReq := f.MustFromStruct(getReq)

				gGetRsp, err := h(ctx, gGetReq)
				if err != nil {
					return
				}

				var getRsp sdk.GetSocialAccountsResponse
				gGetRsp.MustToStruct(&getRsp)

				assert.Len(t, getRsp.SocialAccounts, 1)

				return
			}()
			if err != nil {
				t.Error(err)
			}
		})
	})

	//FTestHackernewsSocialAccount(t, ctx, f, h)
	//FTestHackernewsGetPostFeedContainsPosts(t, ctx, f, h)
	//FTestHackernewsGetPostsSearch(t, ctx, f, h)
	//FTestHackernewsGetSocialAccountBookmarksPosts(t, ctx, f, h)
	//FTestHackernewsGetSocialAccountAuthorsPosts(t, ctx, f, h)
	//FTestHackernewsGetPostFeedsCollection(t, ctx, f, h)
}

func FTestHackernewsSocialAccount(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	name := "TestHackernews"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			getReq := sdk.GetSocialAccountsRequest{
				ServiceFilter: &sdk.ServiceFilter{
					Id: &sdk.ServiceIdFilter{
						Value: &sdk.StringFilter{
							Is: sdk.String("hackernews"),
						},
					},
				},
				Mode: &sdk.GetMode{
					Kind: &sdk.GetModeKind.Id,
					Id: &sdk.Id{
						Kind: &sdk.IdKind.ServiceId,
						ServiceId: &sdk.ServiceId{
							Value:       sdk.String("21stio"),
							ServiceName: sdk.String("hackernews"),
						},
					},
				},
				Select: &sdk.GetSocialAccountsResponseSelect{
					SocialAccounts: &sdk.SocialAccountSelect{
						All: sdk.Bool(true),
						Relations: &sdk.SocialAccountRelationsSelect{
							AuthorsPosts: &sdk.PostsCollectionSelect{
								Errors: &sdk.ErrorSelect{
									Message: sdk.Bool(true),
								},
								Posts: &sdk.PostSelect{
									Id: &sdk.ServiceIdSelect{
										Value:       sdk.Bool(true),
										ServiceName: sdk.Bool(true),
									},
								},
							},
						},
					},
				},
				Relations: &sdk.GetSocialAccountsRelations{
					AuthorsPosts: &sdk.GetPostsCollection{
						Select: &sdk.PostsCollectionSelect{
							Errors: &sdk.ErrorSelect{
								Message: sdk.Bool(true),
							},
							Posts: &sdk.PostSelect{
								Id: &sdk.ServiceIdSelect{
									Value:       sdk.Bool(true),
									ServiceName: sdk.Bool(true),
								},
							},
						},
					},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			gGetRsp.Print()

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}

func FTestHackernewsGetPostFeedContainsPosts(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	name := "TestHackernewsGetPostFeedContainsPosts"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			getReq := sdk.GetPostsRequest{
				Mode: &sdk.GetMode{
					Kind: &sdk.GetModeKind.Relation,
					Relation: &sdk.RelationGetMode{
						Id: &sdk.ServiceId{
							ServiceName: sdk.String("hackernews"),
							Value:       sdk.String("topstories"),
						},
						Relation: &sdk.PostFeedRelationName.PostFeedContainsPosts,
					},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			gGetRsp.Print()

			getRsp := sdk.GetPostsResponse{}
			gGetRsp.MustToStruct(&getRsp)

			assert.True(t, len(getRsp.Posts) != 0)

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}

func FTestHackernewsGetPostsSearch(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	name := "TestHackernewsGetPostsSearch"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			getReq := sdk.GetPostsRequest{
				Mode: &sdk.GetMode{
					Kind: &sdk.GetModeKind.Search,
					Search: &sdk.SearchGetMode{
						Term: sdk.String("book recommendations"),
					},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			gGetRsp.Print()

			getRsp := sdk.GetPostsResponse{}
			gGetRsp.MustToStruct(&getRsp)

			assert.True(t, len(getRsp.Posts) != 0)

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}

func FTestHackernewsGetSocialAccountAuthorsPosts(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	name := "TestHackernewsGetSocialAccountAuthorsPosts"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			getReq := sdk.GetPostsRequest{
				Mode: &sdk.GetMode{
					Kind: &sdk.GetModeKind.Relation,
					Relation: &sdk.RelationGetMode{
						Id: &sdk.ServiceId{
							ServiceName: sdk.String("hackernews"),
							Value:       sdk.String("21stio"),
						},
						Relation: &sdk.SocialAccountRelationName.SocialAccountAuthorsPosts,
					},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			gGetRsp.Print()

			getRsp := sdk.GetPostsResponse{}
			gGetRsp.MustToStruct(&getRsp)

			assert.True(t, len(getRsp.Posts) != 0)

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}

func FTestHackernewsGetPostFeedsCollection(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	name := "TestHackernewsGetPostFeedsCollection"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			getReq := sdk.GetPostFeedsRequest{
				Filter: &sdk.PostFeedFilter{
					Id: &sdk.ServiceIdFilter{
						Value: &sdk.StringFilter{
							Contains: sdk.String("top"),
						},
					},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			gGetRsp.Print()

			getRsp := sdk.GetPostFeedsResponse{}
			gGetRsp.MustToStruct(&getRsp)
			assert.True(t, len(getRsp.PostFeeds) != 0)

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}

func FTestHackernewsGetSocialAccountBookmarksPosts(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	name := "TestHackernewsGetSocialAccountBookmarksPosts"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			getReq := sdk.GetPostsRequest{
				Mode: &sdk.GetMode{
					Kind: &sdk.GetModeKind.Relation,
					Relation: &sdk.RelationGetMode{
						Id: &sdk.ServiceId{
							ServiceName: sdk.String("hackernews"),
							Value:       sdk.String("peter_d_sherman"),
						},
						Relation: &sdk.SocialAccountRelationName.SocialAccountBookmarksPosts,
					},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			//gGetRsp.Print()

			getRsp := sdk.GetPostsResponse{}
			gGetRsp.MustToStruct(&getRsp)

			utils.Print(getRsp.Pagination)

			assert.True(t, len(getRsp.Posts) != 0)

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}
