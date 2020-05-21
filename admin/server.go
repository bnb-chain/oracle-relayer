package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/binance-chain/go-sdk/types/msg"
	"github.com/gorilla/mux"

	"github.com/binance-chain/oracle-relayer/executor/bbc"
	"github.com/binance-chain/oracle-relayer/util"
)

const (
	DefaultListenAddr = "0.0.0.0:8080"
)

type Admin struct {
	Config      *util.Config
	BBCExecutor *bbc.Executor
}

func NewAdmin(config *util.Config, executor *bbc.Executor) *Admin {
	return &Admin{
		Config:      config,
		BBCExecutor: executor,
	}
}

func (admin *Admin) Endpoints(w http.ResponseWriter, r *http.Request) {
	endpoints := struct {
		Endpoints []string `json:"endpoints"`
	}{
		Endpoints: []string{
			"/skip_sequence/{sequence}/{claim_type_to_skip}/{sequence_to_skip}"},
	}

	jsonBytes, err := json.MarshalIndent(endpoints, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (admin *Admin) SkipSequence(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	seqStr := params["sequence"]
	if seqStr == "" {
		http.Error(w, "required parameter 'sequence' is missing", http.StatusBadRequest)
		return
	}

	seq, err := strconv.ParseInt(seqStr, 10, 64)
	if err != nil {
		http.Error(w, "'sequence' is invalid", http.StatusBadRequest)
		return
	}

	claimTypeToSkipStr := params["claim_type_to_skip"]
	if claimTypeToSkipStr == "" {
		http.Error(w, "required parameter 'claim_type_to_skip' is missing", http.StatusBadRequest)
		return
	}

	claimType, err := strconv.ParseInt(claimTypeToSkipStr, 10, 64)
	if err != nil {
		http.Error(w, "'claim_type_to_skip' is invalid", http.StatusBadRequest)
		return
	}
	if !msg.IsValidClaimType(msg.ClaimType(claimType)) {
		http.Error(w, "'claim_type_to_skip' is invalid", http.StatusBadRequest)
		return
	}

	seqToSkipStr := params["sequence_to_skip"]
	if seqToSkipStr == "" {
		http.Error(w, "required parameter 'sequence_to_skip' is missing", http.StatusBadRequest)
		return
	}
	seqToSkip, err := strconv.ParseInt(seqToSkipStr, 10, 64)
	if err != nil {
		http.Error(w, "'sequence_to_skip' is invalid", http.StatusBadRequest)
		return
	}

	skipClaim := msg.SkipSequenceClaim{
		ClaimType: msg.ClaimType(claimType),
		Sequence:  seqToSkip,
	}

	skipClaimBz, err := json.Marshal(skipClaim)
	if err != nil {
		http.Error(w, "marshal claim error", http.StatusInternalServerError)
		return
	}

	err = admin.BBCExecutor.Claim(msg.ClaimTypeSkipSequence, seq, string(skipClaimBz))
	if err != nil {
		http.Error(w, fmt.Sprintf("claim error, err=%s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (admin *Admin) Serve() {
	router := mux.NewRouter()

	router.HandleFunc("/", admin.Endpoints)
	router.HandleFunc("/skip_sequence/{sequence}/{claim_type_to_skip}/{sequence_to_skip}", admin.SkipSequence)

	listenAddr := DefaultListenAddr
	if admin.Config.AdminConfig != nil && admin.Config.AdminConfig.ListenAddr != "" {
		listenAddr = admin.Config.AdminConfig.ListenAddr
	}
	srv := &http.Server{
		Handler:      router,
		Addr:         listenAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	util.Logger.Infof("start admin server at %s", srv.Addr)

	err := srv.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("start admin server error, err=%s", err.Error()))
	}
}
