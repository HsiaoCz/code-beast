package service

import "github.com/HsiaoCz/code-beast/lenven/internal/pb"

type BookService struct {
	pb.UnimplementedLenvenServer
}

func NewBookService() *BookService {
	return &BookService{}
}
