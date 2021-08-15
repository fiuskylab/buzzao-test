package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fiuskylab/buzzao-test/data"
	"github.com/fiuskylab/buzzao-test/utils"
	"github.com/fiuskylab/buzzao-test/workerpool"
	fiber "github.com/gofiber/fiber/v2"
)

// Sv a
var Sv *fiber.App

type Server struct {
	App *fiber.App
	P   *workerpool.Pool
}

// NewServer a
func NewServer() *Server {
	return &Server{
		App: fiber.New(fiber.Config{
			CaseSensitive: true,
			StrictRouting: false,
			Concurrency:   256 * 1024,
			WriteTimeout:  time.Duration(time.Second * 45),
		}),
	}
}

const (
	threadLowerThanZero = "The number of threads must be higher than 0, given %d"
)

func (s *Server) Listen(port string) {

	s.App.Listen(port)
}

type NumBody struct {
	Nums []int `json:"nums"`
}

type j map[string]interface{}

