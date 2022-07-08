package library

import (
	"context"
	"reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func Test_service_GetBooks(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := NewMockRepository(ctrl)

	books := []Book{
		{Isbn: "44778854", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}},
		{Isbn: "3987654", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}},
	}
	m.
		EXPECT().
		GetBooks(context.Background()).
		Return(books, nil)

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *service
		args    args
		want    []Book
		wantErr bool
	}{

		{
			name:    "Getting books",
			s:       &service{repo: m},
			args:    args{ctx: context.Background()},
			want:    books,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetBooks(tt.args.ctx)
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
