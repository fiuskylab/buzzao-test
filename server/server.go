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
	s.App.Post("/config/:n", s.configEndpoint)
	s.App.Post("/process", s.processEndpoint)

	s.App.Listen(port)
}

type NumBody struct {
	Nums []int `json:"nums"`
}

type j map[string]interface{}

func (s *Server) processEndpoint(c *fiber.Ctx) error {
	if s.P == nil {
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(j{
				"error": "Worker Pool no set yet, set using the endpoint POST: /config/:n ",
			})
	}
	numBody := new(NumBody)

	if err := c.BodyParser(numBody); err != nil {
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(j{
				"error": err.Error(),
			})
	}

	d := data.NewData(s.P.Concurrency, numBody.Nums)

	s.P.AddJob(workerpool.NewJob(data.EachBatch, d))
	s.P.Run()

	return c.
		Status(http.StatusOK).
		JSON(j{
			"result": d.Sum,
		})
}

func (s *Server) configEndpoint(c *fiber.Ctx) error {
	n, err := utils.StrToInt(c.Params("n"))
	if err != nil {
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(j{
				"error": err.Error(),
			})
	}

	p := workerpool.NewPool(make([]*workerpool.Job, 0), n)
	s.P = p

	if err := s.P.SetConcurrency(n); err != nil {
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(j{
				"error": err.Error(),
			})
	}

	return c.
		Status(http.StatusOK).
		JSON(j{
			"message": fmt.Sprintf("Number of threads set to %d", s.P.Concurrency),
		})
}
