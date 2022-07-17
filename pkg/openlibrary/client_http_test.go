package openlibrary

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestHttpClient_Search(t *testing.T) {
	ctx := context.Background()
	type args struct {
		sreq SearchRequest
	}
	tests := []struct {
		name    string
		args    args
		want    SearchResponse
		wantErr bool
	}{
		{
			name: "aa",
			args: args{
				sreq: SearchRequest{
					Author: "Tolkien",
				},
			},
			want:    SearchResponse{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hc := NewHttpClient(&http.Client{}, "http://openlibrary.org")
			got, err := hc.Search(ctx, tt.args.sreq)
			fmt.Printf("Resp: %+v\n", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpClient.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HttpClient.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
