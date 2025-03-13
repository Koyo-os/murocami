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
	cfg *config.Config
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
		cfg: cfg,
	}
}

func (h Handler) Routes(mux *http.ServeMux){
	mux.HandleFunc("/webhook", h.WebHookHandler)
	mux.HandleFunc("/ui", h.MainPage)

	if h.cfg.UseUI {
		mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(h.cfg.StaticDir))))
	}
}

func (h *Handler) runAgent(url string) (bool, error) {
	ok, err := h.Agent.Run(url)
	if err != nil{
		h.Logger.Errorf("error run agent: %v", err)
		return ok,err
	}

	return ok, nil
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
			ok,err := h.runAgent(payload.Repository.CloneURL)
			if err != nil{
				h.Logger.Errorf("error run agent: %v",err)
			}

			if !ok {
				h.Logger.Debugf("agent for %s stopped with !ok variable, please check your code", payload.Repository.CloneURL)
				fmt.Fprintf(w, "agent for %s stopped with !ok variable, please check your code", payload.Repository.CloneURL)
			}
        }()
    }

	fmt.Fprint(w, "success!")

	h.Logger.Info("github webhook recieved successfully!")
}