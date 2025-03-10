package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/koyo-os/murocami/internal/agent"
	"github.com/koyo-os/murocami/internal/config"
	"github.com/koyo-os/murocami/internal/models"
	"github.com/koyo-os/murocami/pkg/logger"
)

type Handler struct{
	Logger *logger.Logger
	Agent *agent.Agent
}

func InitHandler(cfg *config.Config) *Handler {
	logger := logger.Init()
	agent, err := agent.Init(cfg)
	if err != nil{
		logger.Infof("error get agent: %v", err)
		return nil
	}

	return &Handler{
		Logger: logger,
		Agent: agent,
	}
}

func (h Handler) Routes(mux *http.ServeMux){
	mux.HandleFunc("/webhook", h.WebHookHandler)
}

func (h *Handler) runAgent(url string) error {
	err := h.Agent.Run(url)
	if err != nil{
		h.Logger.Errorf("error run agent: %v", err)
		return err
	}

	return nil
}

func (h *Handler) WebHookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method must ve post", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil{
		http.Error(w, "cant read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var payload models.WebhookPayload
	if err = sonic.Unmarshal(body, &payload);err != nil{
		http.Error(w, "cant get payload", http.StatusInternalServerError)
		return
	}

	if payload.Ref == "refs/heads/main" {
		h.Logger.Info("starting agent...")

        go func() {
			if err = h.runAgent(payload.Repository.CloneURL);err != nil{
				h.Logger.Errorf("error to run agent: %v", err)
			}
        }()
    }

	fmt.Fprint(w, "success!")

	h.Logger.Info("github webhook recieved successfully!")
}