package api

import (
	"encoding/json"
	db "go-oapi-test/db/sqlc"
	rabbit "go-oapi-test/rabbit/events"
	"go-oapi-test/tools"
	"io"
	"net/http"
	"strconv"
	"strings"

	jwt2 "github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rabbitmq/amqp091-go"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ ServerInterface = (*Server)(nil)

type Server struct {
	db      *db.Queries
	rabbit  *amqp091.Channel
	jwtAuth tools.Authenticator
}

func (s *Server) DeleteBranch(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	defer func() { _ = r.Body.Close() }()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("could not bind request body"))
		return
	}

	var branchIds *BranchDeleteDto

	err = json.Unmarshal(bodyBytes, &branchIds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("could not bind request body " + err.Error()))
		return
	}

	err = s.db.DeleteBranches(r.Context(), *branchIds.BranchIds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}

	for _, branch := range *branchIds.BranchIds {
		body, _ := json.Marshal(rabbit.BranchRemoved{
			BranchId: int(branch),
		})

		s.rabbit.Publish("Events:BranchRemoved", "", false, false, amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) GetAllBranches(w http.ResponseWriter, r *http.Request) {
	branches, err := s.db.ListBranches(r.Context(), db.ListBranchesParams{
		Limit:  100,
		Offset: 0,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			return
		}
	}

	var result GetAllBranches200JSONResponse

	for _, b := range branches {
		result = append(result, Branch{
			Id:           &b.Id,
			Name:         &b.Name,
			MaxUsers:     &b.MaxUsers,
			CurrentUsers: &b.CurrentUsers,
			GroupIds:     &b.GroupIds,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(result)
}

func (s *Server) CreateBranch(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	defer func() { _ = r.Body.Close() }()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("could not bind request body"))
		return
	}

	var branch *Branch
	err = json.Unmarshal(bodyBytes, &branch)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("could not bind request body"))
		return
	}

	branchDb, err := s.db.CreateBranch(r.Context(), db.CreateBranchParams{
		Name:     *branch.Name,
		MaxUsers: *branch.MaxUsers,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("could not create branch"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(branchDb)
}

func (s *Server) CheckBranchLimit(w http.ResponseWriter, r *http.Request, params CheckBranchLimitParams) {
	token, err := tools.GetJWSFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	jwt, err := s.jwtAuth.ValidateJWS(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	branchesInterface, success := jwt.Claims.(jwt2.MapClaims)["branch"]
	if !success {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("could not find branch"))
		return
	}

	branchesString := branchesInterface.(string)

	branchesList := strings.Split(branchesString, ", ")

	branches := make([]int, len(branchesList))
	for i, v := range branchesList {
		branches[i], err = strconv.Atoi(v)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
	}
	if params.BranchId == nil {
		branchId := int32(branches[0])
		params.BranchId = &branchId
	}
	branch, err := s.db.GetBranchById(r.Context(), *params.BranchId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	result := params.UsersAmount+branch.CurrentUsers <= branch.MaxUsers
	_ = json.NewEncoder(w).Encode(result)
}

func (s *Server) GetBranchById(w http.ResponseWriter, r *http.Request, branchId int32) {
	branch, err := s.db.GetBranchById(r.Context(), branchId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(branch)
}

func (s *Server) UpdateBranch(w http.ResponseWriter, r *http.Request, branchId int32) {
	bodyBytes, err := io.ReadAll(r.Body)
	defer func() { _ = r.Body.Close() }()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("could not bind request body"))
		return
	}

	var updateBranchDto *UpdateBranchDto
	err = json.Unmarshal(bodyBytes, &updateBranchDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("could not bind request body"))
		return
	}

	branch, err := s.db.UpdateBranch(r.Context(), db.UpdateBranchParams{
		Name:     pgtype.Text{String: *updateBranchDto.Name, Valid: *updateBranchDto.Name != ""},
		MaxUsers: pgtype.Int4{Int32: *updateBranchDto.MaxUsers, Valid: *updateBranchDto.MaxUsers != 0},
		BranchID: branchId,
	})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(branch)
}

func NewServer(db *db.Queries, jwtAuth tools.Authenticator, rabbit *amqp091.Channel) Server {
	return Server{db: db, jwtAuth: jwtAuth, rabbit: rabbit}
}
