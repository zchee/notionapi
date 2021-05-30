package notionapi_test

import (
	"context"
	"github.com/jomei/notionapi"
	"reflect"
	"testing"
	"time"
)

func TestPageClient(t *testing.T) {
	timestamp, err := time.Parse(time.RFC3339, "2021-05-24T05:06:34.827Z")
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Get", func(t *testing.T) {
		tests := []struct {
			name     string
			filePath string
			id       notionapi.PageID
			want     *notionapi.Page
			wantErr  bool
			err      error
		}{
			{
				name:     "returns page by id",
				id:       "some_id",
				filePath: "testdata/page_get.json",
				want: &notionapi.Page{
					Object:         notionapi.ObjectTypePage,
					ID:             "some_id",
					CreatedTime:    timestamp,
					LastEditedTime: timestamp,
					Parent: notionapi.Parent{
						Type:       notionapi.ParentTypeDatabaseID,
						DatabaseID: "some_id",
					},
					Archived: false,
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath)
				client := notionapi.NewClient("some_token", notionapi.WithHTTPClient(c))

				got, err := client.Page.Get(context.Background(), tt.id)
				if (err != nil) != tt.wantErr {
					t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				// TODO: remove properties from comparing for a while. Have to compare with interface somehow
				got.Properties = nil
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Get() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("Create", func(t *testing.T) {
		tests := []struct {
			name     string
			filePath string
			id       notionapi.PageID
			request  *notionapi.PageCreateRequest
			want     *notionapi.Page
			wantErr  bool
			err      error
		}{
			{
				name:     "returns a new page",
				filePath: "testdata/page_create.json",
				request: &notionapi.PageCreateRequest{
					Parent: notionapi.Parent{
						Type:       notionapi.ParentTypeDatabaseID,
						DatabaseID: "f830be5eff534859932e5b81542b3c7b",
					},
					Properties: notionapi.Properties{
						"Name": notionapi.PageTitleProperty{
							Title: notionapi.Paragraph{
								{Text: notionapi.Text{Content: "hello"}},
							},
						},
					},
				},
				want: &notionapi.Page{
					Object:         notionapi.ObjectTypePage,
					ID:             "some_id",
					CreatedTime:    timestamp,
					LastEditedTime: timestamp,
					Parent: notionapi.Parent{
						Type:       notionapi.ParentTypeDatabaseID,
						DatabaseID: "some_id",
					},
					Archived: false,
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath)
				client := notionapi.NewClient("some_token", notionapi.WithHTTPClient(c))
				got, err := client.Page.Create(context.Background(), tt.request)
				if (err != nil) != tt.wantErr {
					t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				// TODO: remove properties from comparing for a while. Have to compare with interface somehow
				got.Properties = nil
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Create() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("Update", func(t *testing.T) {
		tests := []struct {
			name     string
			filePath string
			id       notionapi.PageID
			request  *notionapi.PageUpdateRequest
			want     *notionapi.Page
			wantErr  bool
			err      error
		}{
			{
				name:     "change requested properties and return the result",
				id:       "some_id",
				filePath: "testdata/page_update.json",
				request: &notionapi.PageUpdateRequest{
					Properties: notionapi.Properties{
						"SomeColumn": notionapi.RichTextProperty{
							Type: notionapi.PropertyTypeRichText,
							RichText: notionapi.Paragraph{
								{
									Type: notionapi.ObjectTypeText,
									Text: notionapi.Text{Content: "patch"},
								},
							},
						},
					},
				},
				want: &notionapi.Page{
					Object:         notionapi.ObjectTypePage,
					ID:             "some_id",
					CreatedTime:    timestamp,
					LastEditedTime: timestamp,
					Parent: notionapi.Parent{
						Type:       notionapi.ParentTypeDatabaseID,
						DatabaseID: "some_id",
					},
					Archived: false,
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath)
				client := notionapi.NewClient("some_token", notionapi.WithHTTPClient(c))
				got, err := client.Page.Update(context.Background(), tt.id, tt.request)
				if (err != nil) != tt.wantErr {
					t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				// TODO: remove properties from comparing for a while. Have to compare with interface somehow
				got.Properties = nil
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Update() got = %v, want %v", got, tt.want)
				}
			})
		}
	})
}
