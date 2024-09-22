package api

import (
	"encoding/json"
	"go-oapi-test/db/sqlc"
	"go-oapi-test/tools"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ ServerInterface = (*Server)(nil)

type Server struct {
	db      *db.Queries
	jwtAuth tools.Authenticator
}

func (s *Server) DeleteBranch(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
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
			Id:           &b.ID,
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

	var branch Branch
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

	branchesInterface, success := jwt.Get("branch")
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

}

func (s *Server) GetBranchById(w http.ResponseWriter, r *http.Request, branchId int32) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) UpdateBranch(w http.ResponseWriter, r *http.Request, branchId int32) {
	//TODO implement me
	panic("implement me")
}

func NewServer(db *db.Queries, jwtAuth tools.Authenticator) Server {
	return Server{db: db, jwtAuth: jwtAuth}
}
