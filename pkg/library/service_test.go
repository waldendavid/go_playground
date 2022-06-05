package library

import (
	"context"
	"reflect"
	"testing"
)

func Test_service_GetBooks(t *testing.T) {

	tests := []struct {
		name    string
		s       *service
		want    []Book
		wantErr bool
	}{

		{
			name: "abc",
			s:    &service{},
			want: []Book{
				{ID: "1", Isbn: "44778854", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}},
				{ID: "2", Isbn: "3987654", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetBooks(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetBooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetBooks() = %v, want %v", got, tt.want)
			}
		})
	}
}
