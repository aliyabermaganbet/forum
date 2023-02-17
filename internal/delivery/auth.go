// done clean architecture
package delivery

import (
	"errors"
	"net/http"
	"strings"

	"forum/internal/models"
	"forum/internal/service"
)

func (h *Handler) signIn(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/signin" {
		h.errorPage(writer, request, http.StatusNotFound)
		return
	}
	userId := request.Context().Value(userIdCtx).(int)
	if userId != 0 {
		h.errorPage(writer, request, http.StatusBadRequest)
		return
	}
	switch request.Method {
	case http.MethodGet:
		if err := h.render(writer, "signin.html", nil); err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		if err := request.ParseForm(); err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
		p, ok := request.Form["password"]
		if !ok {
			h.errorPage(writer, request, http.StatusBadRequest)
			return
		}
		u, ok := request.Form["username"]
		if !ok {
			h.errorPage(writer, request, http.StatusBadRequest)
			return
		}
		if IsTagEmpty(u[0]) || IsTagEmpty(p[0]) { // checkonemore time this function
			h.errorPage(writer, request, http.StatusBadRequest)
			return
		}
		user := models.User{
			Username: u[0],
			Password: p[0],
		}

		cookie, err := h.services.GenerateSessionToken(user)
		if err != nil {
			if errors.Is(err, service.ErrUserNotExist) {
				user.Erstring = "The username you entered doesn't belong to an account.Please check your username and try again"
				err := h.render(writer, "signin.html", user)
				if err != nil {
					h.errorPage(writer, request, http.StatusInternalServerError)
					return
				}
				return
			} else if errors.Is(err, service.ErrIncorrectPassword) {
				user.Erstring = "Your password is incorrect. Try again"
				err := h.render(writer, "signin.html", user)
				if err != nil {
					h.errorPage(writer, request, http.StatusInternalServerError)
					return
				}
				return
			}
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}

		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/profile", http.StatusSeeOther) // redirects to profile
	}
}

func (h *Handler) signOut(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/sign-out" {
		h.errorPage(writer, request, http.StatusNotFound)
		return
	}
	if request.Method != http.MethodGet {
		h.errorPage(writer, request, http.StatusMethodNotAllowed)
		return
	}
	user_id := request.Context().Value(userIdCtx).(int)
	if err := h.services.DeleteSessionToken(int64(user_id)); err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
	err := h.render(writer, "home.html", nil)
	if err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
}

func (h *Handler) signUp(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/signup" {
		h.errorPage(writer, request, http.StatusNotFound)
		return
	}
	userId := request.Context().Value(userIdCtx).(int)
	if userId != 0 {
		h.errorPage(writer, request, http.StatusBadRequest)
		return
	}
	switch request.Method {
	case http.MethodGet:
		if err := h.render(writer, "signup.html", nil); err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		if err := request.ParseForm(); err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
		u, ok := request.Form["username"] // you get your username here, you can also get email, password
		if !ok {
			http.Error(writer, "Username field not found", 400)
			return
		}
		username := u[0]
		e, ok := request.Form["email"] // everthng is bad request, 400 code
		if !ok {
			http.Error(writer, "Email field not found", 400)
			return
		}
		email := e[0]
		p, ok := request.Form["password"]
		if !ok {
			http.Error(writer, "Password field not found", 400)
			return
		}
		password := p[0]
		user := models.User{ // put into database by using structure
			Username: username,
			Email:    email,
			Password: password,
		}
		if IsTagEmpty(username) || IsTagEmpty(password) || IsTagEmpty(email) {
			http.Redirect(writer, request, "/signup", http.StatusBadRequest)
			return
		}

		err := h.services.CreateUser(user) // if there is an error, because scan returned error (does not match), found nothing in the database with the same username. Send it to the CreateUser function to create the new one
		if err != nil {
			if errors.Is(err, service.ErrAuth) {
				x := models.User{
					Erstring: "Please follow the instructions below to create username",
				}
				err := h.render(writer, "signup.html", x)
				if err != nil {
					h.errorPage(writer, request, http.StatusInternalServerError)
					return
				}
				return
			} else if errors.Is(err, service.ErrUserExist) {
				x := models.User{
					Erstring: "Can't create new user account. Username already exists",
				}
				writer.WriteHeader(400)
				err := h.render(writer, "signup.html", x)
				if err != nil {
					h.errorPage(writer, request, http.StatusInternalServerError)
					return
				}
				return
			} else if errors.Is(err, service.ErrEmailInvalid) {
				x := models.User{
					Erstring: "Sorry, your email address is invalid. Please try again",
				}
				writer.WriteHeader(400)
				err := h.render(writer, "signup.html", x)
				if err != nil {
					h.errorPage(writer, request, http.StatusInternalServerError)
					return
				}
				return
			} else if errors.Is(err, service.ErrEmailExist) {
				x := models.User{
					Erstring: "The email address is already taken. Please choose another one",
				}
				err := h.render(writer, "signup.html", x)
				if err != nil {
					h.errorPage(writer, request, http.StatusInternalServerError)
					return
				}
				return
			} else if errors.Is(err, service.ErrPasswordInvalid) {
				x := models.User{
					Erstring: "Please follow the instructions below to create password",
				}
				err := h.render(writer, "signup.html", x)
				if err != nil {
					h.errorPage(writer, request, http.StatusInternalServerError)
					return
				}
				return
			}
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
		http.Redirect(writer, request, "/signin", http.StatusFound)
		return

	}
}

func IsTagEmpty(s string) bool {
	cutSpace := strings.TrimSpace(s)
	if len(cutSpace) == 0 {
		return true
	}
	return false
}
