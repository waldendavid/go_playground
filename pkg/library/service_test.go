package library

// import (
// 	"context"
// 	"reflect"
// 	"testing"
// )

// func Test_service_GetBooks(t *testing.T) {

// 	tests := []struct {
// 		name    string
// 		s       *service
// 		want    []Book
// 		wantErr bool
// 	}{

// 		{
// 			name: "Saving",
// 			s: &service{books: []Book{
// 				{ID: "1", Isbn: "44778854", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}},
// 				{ID: "2", Isbn: "3987654", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}},
// 			}},
// 			want: []Book{
// 				{ID: "1", Isbn: "44778854", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}},
// 				{ID: "2", Isbn: "3987654", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}},
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := tt.s.GetBooks(context.Background())
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("service.GetBooks() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("service.GetBooks() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_service_UpdateBook(t *testing.T) {
// 	type args struct {
// 		ctx  context.Context
// 		book Book
// 		id   string
// 	}
// 	tests := []struct {
// 		name    string
// 		s       *service
// 		args    args
// 		want    Book
// 		wantErr bool
// 	}{
// 		{
// 			name:    "updating",
// 			s:       &service{},
// 			args:    args{ctx: context.Background(), book: Book{ID: "1", Isbn: "999", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}}, id: "1"},
// 			want:    Book{ID: "1", Isbn: "999", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := tt.s.UpdateBook(tt.args.ctx, tt.args.book, tt.args.id)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("service.UpdateBook() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("service.UpdateBook() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_service_GetBook(t *testing.T) {
// 	type args struct {
// 		ctx context.Context
// 		id  string
// 	}
// 	tests := []struct {
// 		name    string
// 		s       *service
// 		args    args
// 		want    Book
// 		wantErr bool
// 	}{
// 		{
// 			name: "getting one book",
// 			s: &service{books: []Book{
// 				{ID: "1", Isbn: "44778854", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}},
// 				{ID: "2", Isbn: "3987654", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}},
// 			}},
// 			args:    args{ctx: context.Background(), id: "1"},
// 			want:    Book{ID: "1", Isbn: "44778854", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := tt.s.GetBook(tt.args.ctx, tt.args.id)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("service.GetBook() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("service.GetBook() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_service_DeleteBook(t *testing.T) {
// 	type args struct {
// 		ctx context.Context
// 		id  string
// 	}
// 	tests := []struct {
// 		name    string
// 		s       *service
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name:    "deleting one book",
// 			s:       &service{},
// 			args:    args{ctx: context.Background(), id: "1"},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := tt.s.DeleteBook(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
// 				t.Errorf("service.DeleteBook() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func Test_service_CreateBook(t *testing.T) {
// 	type args struct {
// 		ctx  context.Context
// 		book Book
// 	}
// 	tests := []struct {
// 		name    string
// 		s       *service
// 		args    args
// 		want    Book
// 		wantErr bool
// 	}{
// 		{
// 			name:    "creating one book",
// 			s:       &service{},
// 			args:    args{ctx: context.Background(), book: Book{Isbn: "987654321", Title: "Posted Book", Author: &Author{Firstname: "Creator", Lastname: "Poster"}}},
// 			want:    Book{Isbn: "987654321", Title: "Posted Book", Author: &Author{Firstname: "Creator", Lastname: "Poster"}},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := tt.s.CreateBook(tt.args.ctx, tt.args.book)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("service.CreateBook() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got.ID != 0 {
// 				tt.want.ID = got.ID
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("service.CreateBook() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
