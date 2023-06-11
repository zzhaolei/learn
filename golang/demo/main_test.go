package main

import (
	"errors"
	"testing"
)

func ValidateName(name string) bool {
	if len(name) < 1 {
		return false
	} else if len(name) > 12 {
		return false
	}
	return true
}

//go:noinline
func ValidateNameNoInline(name string) bool {
	if len(name) < 1 {
		return false
	} else if len(name) > 12 {
		return false
	}
	return true
}

func (s *Server) CreateUser(name string, password string) error {
	if !ValidateName(name) {
		return errors.New("invalid name")
	}
	return nil
}

// CreateUserNoInline 使用的是禁止内联版本的 ValidateName
func (s *Server) CreateUserNoInline(name string, password string) error {
	if !ValidateNameNoInline(name) {
		return errors.New("invalid name")
	}
	return nil
}

type Server struct{}

// BenchmarkCreateUser 测试内联过的函数的性能
func BenchmarkCreateUser(b *testing.B) {
	srv := Server{}
	for i := 0; i < b.N; i++ {
		if err := srv.CreateUser("bootun", "123456"); err != nil {
			b.Logf("err: %v", err)
		}
	}
}

// BenchmarkValidateNameNoInline 测试函数禁止内联后的性能
func BenchmarkValidateNameNoInline(b *testing.B) {
	srv := Server{}
	for i := 0; i < b.N; i++ {
		if err := srv.CreateUserNoInline("bootun", "123456"); err != nil {
			b.Logf("err: %v", err)
		}
	}
}
