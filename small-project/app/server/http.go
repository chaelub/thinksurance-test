package server

import (
	"chaelub/thinksurance/small-project/domain/models"
	"chaelub/thinksurance/small-project/domain/repo/role"
	"chaelub/thinksurance/small-project/domain/store/user"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
)

// this server just a simple way to interact with core logic

type HTTPServerConfig struct {
	Port string `toml:"port"`
}

type loginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type loginResponse struct {
	User models.User
	Role models.Role
}

func RunServer(conf HTTPServerConfig, userStore user.UserStoreI, roleRepo role.RoleRepoI) {

	http.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "only POST requests allowed")
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			sendResponse(w, http.StatusInternalServerError, nil, err)
			return
		}

		lr := new(loginRequest)
		if err = json.Unmarshal(body, lr); err != nil {
			sendResponse(w, http.StatusInternalServerError, nil, err)
			return
		}

		user, err := userStore.GetByLogin(lr.Login)
		if err != nil {
			sendResponse(w, http.StatusInternalServerError, nil, err)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(lr.Password))
		if err != nil {
			sendResponse(w, http.StatusInternalServerError, nil, err)
			return
		}

		role, err := roleRepo.GetRoleById(user.RoleId)
		if err != nil {
			sendResponse(w, http.StatusInternalServerError, nil, err)
			return
		}

		// todo: save role with permissions into cache

		payload := loginResponse{
			User: user,
			Role: role,
		}

		sendResponse(w, http.StatusOK, payload, nil)

	})
	http.ListenAndServe(conf.Port, nil)
}

type response struct {
	Error   string      `json:"error,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

func sendResponse(w http.ResponseWriter, status int, payload interface{}, err error) {
	w.WriteHeader(status)

	resp := response{
		Payload: payload,
	}

	if err != nil {
		resp.Error = err.Error()
	}

	respBody, err := json.Marshal(resp)
	if err != nil {
		panic(fmt.Sprintf("can't marshal response: %+v", err))
	}
	_, err = w.Write(respBody)
	if err != nil {
		panic(fmt.Sprintf("can't send the response: %+v", err))
	}

}
