package handlers

//
//import (
//	"go.uber.org/zap"
//	"net/http"
//	"strings"
//)
//
//func (a Api) validateToken(w http.ResponseWriter, r *http.Request, logger zap.SugaredLogger) (err error) {
//	w.Header().Set("Content-Type", "application/json")
//	ctx := r.Context()
//
//	token := r.Header.Get("Authorization")
//	if len(token) == 0 {
//		w.WriteHeader(http.StatusBadRequest)
//		logger.Debug("Токен не предоставлен", err)
//		return err
//	}
//
//	login := r.URL.Query().Get("login")
//	if login == "" {
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	// Проверяем формат токена
//	parts := strings.SplitN(token, " ", 2)
//	if parts[0] != "Bearer" {
//		logger.Errorf("Неверный формат токена", err)
//		return err
//	}
//
//	token = parts[1]
//	err = a.userStorage.ValidateToken(ctx, login, token)
//
//	return err

//
//claims, ok := token.Claims.(jwt.MapClaims)
//if !ok {
//	fmt.Fprintf(w, "couldn't parse claims")
//	return errors.New("Token error")
//}
//
//exp := claims["exp"].(float64)
//if int64(exp) < time.Now().Local().Unix() {
//	fmt.Fprintf(w, "token expired")
//	return errors.New("Token error")
//}

//}
