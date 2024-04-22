package helper

import (
	"LoginAPI/controller"
	"LoginAPI/model"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// var wg sync.WaitGroup

type LoginResponse struct {
	Code             string            `json:"code"`
	UserInformartion *UserInformartion `json:"UserInformation"`
	Error            *Error            `json:"error"`
}
type UserInformartion struct {
	Number   string `json:"number"`
	Messages string `json:"messages"`
}
type Error struct {
	Messages string `json:"messages"`
}

var wg sync.WaitGroup

func LoginAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LoginRequest")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var RequestData model.LoginData
	_ = json.NewDecoder(r.Body).Decode(&RequestData)
	if RequestData.UserNumber == "" || RequestData.IP == "" {
		jsonData := LoginResponse{Code: "204", Error: &Error{Messages: "Data not Found"}}
		json.NewEncoder(w).Encode(jsonData)
		return
	} else if controller.CheckAlreadyExit(RequestData.UserNumber) {
		controller.UpdateData(RequestData.UserNumber)
		// defer wg.Done()
		// select {
		// case <-ctx.Done():
		// 	// If the context times out, respond without the existing user message
		// 	jsonData := LoginResponse{Code: "200", UserInformartion: &UserInformartion{Number: RequestData.UserNumber, Messages: "User-Updated"}}
		// 	json.NewEncoder(w).Encode(jsonData)
		// }

		// if wg.WaitTimeout(time.Second) {
		// 	// If the goroutine completed within 1 second, include the existing user message
		// 	jsonData := LoginResponse{Code: "200", UserInformartion: &UserInformartion{Number: RequestData.UserNumber, Messages: "Existing-User"}}
		// 	json.NewEncoder(w).Encode(jsonData)
		// } else {
		// If the goroutine did not complete within 1 second, respond without the existing user message
		// 	jsonData := LoginResponse{Code: "200", UserInformartion: &UserInformartion{Number: RequestData.UserNumber, Messages: "User-Updated"}}
		// 	json.NewEncoder(w).Encode(jsonData)
		// // }
		jsonData := LoginResponse{Code: "200", UserInformartion: &UserInformartion{Number: RequestData.UserNumber, Messages: "Existing-User"}}
		json.NewEncoder(w).Encode(jsonData)
		wg.Wait()
		return
	} else {
		val := controller.InsertIntoDB(RequestData)
		if val != nil {
			jsonData := LoginResponse{Code: "200", UserInformartion: &UserInformartion{Number: RequestData.UserNumber, Messages: "New-User"}}
			// result, err := json.Marshal(jsonData)
			// if err != nil {
			// 	panic(err)
			// }
			json.NewEncoder(w).Encode(jsonData)
			return
		}
	}

}
