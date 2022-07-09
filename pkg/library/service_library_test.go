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

func Test_service_GetBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := NewMockRepository(ctrl)
	book := Book{Isbn: "44778854", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}}

	m.
		EXPECT().
		GetBook(context.Background(), gomock.Any()).
		Return(book, nil)

	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		s       *service
		args    args
		want    Book
		wantErr bool
	}{
		{
			name:    "Getting book",
			s:       &service{repo: m},
			args:    args{ctx: context.Background(), id: "5"},
			want:    book,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetBook(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_CreateBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := NewMockRepository(ctrl)
	book := Book{Isbn: "44778854", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}}

	m.
		EXPECT().
		CreateBook(context.Background(), gomock.Any()).
		Return(book, nil)

	type args struct {
		ctx  context.Context
		book Book
	}
	tests := []struct {
		name    string
		s       *service
		args    args
		want    Book
		wantErr bool
	}{
		{
			name:    "Creating book",
			s:       &service{repo: m},
			args:    args{context.Background(), book},
			want:    book,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CreateBook(tt.args.ctx, tt.args.book)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.CreateBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_UpdateBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := NewMockRepository(ctrl)
	// todo where can i put magic numbers?
	bookInput := Book{Isbn: "44778854", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}}
	bookOutput := Book{Isbn: "123", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}}

	m.
		EXPECT().
		UpdateBook(context.Background(), bookInput, gomock.Any()).
		Return(bookOutput, nil)

	type args struct {
		ctx context.Context
		b   Book
		id  string
	}
	tests := []struct {
		name    string
		s       *service
		args    args
		want    Book
		wantErr bool
	}{
		{
			name:    "Updating book",
			s:       &service{repo: m},
			args:    args{ctx: context.Background(), b: bookInput, id: "5"},
			want:    bookOutput,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UpdateBook(tt.args.ctx, tt.args.b, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.UpdateBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.UpdateBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_DeleteBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := NewMockRepository(ctrl)

	m.
		EXPECT().
		DeleteBook(context.Background(), gomock.Any()).
		Return(nil)

	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		s       *service
		args    args
		wantErr bool
	}{
		{
			name:    "Deleting book",
			s:       &service{repo: m},
			args:    args{ctx: context.Background(), id: "5"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.DeleteBook(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("service.DeleteBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
