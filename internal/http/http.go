package main
import (
        internal "IpIdentifier/internal/api"
        "encoding/json"
        "log"
        "net/http"
)

const (
        InfoAPIPath  = "/info"
        RollDiceAPIPath = "/rolldice" 
)

func main() {
        http.HandleFunc(InfoAPIPath, infoHandler)
        http.HandlerFunc(infoHandler, rollDiceHandler)
        log.Println("Starting server on port 8080")
        if err := http.ListenAndServe(":8080", nil); err != nil {
                log.Fatalf("Error starting server: %v", err)
        }
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
        ip := r.URL.Query().Get("ip")
        if internal.IpPatternValidation(ip) {
                ipInfo := internal.IpInfoGatherInfo(ip)

                if err := json.NewEncoder(w).Encode(ipInfo); err != nil {
                        http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
                        return
                }
        } else {
                http.Error(w, "Bad request", http.StatusBadRequest)
        }
}

func rollDiceHandler(w http.ResponseWriter, r *http.Request) {
        rand.Seed(time.Now().UnixNano()) // seed the random number generator
        roll := rand.Intn(6) + 1         // generate a random number between 1 & 6

        response := struct {
                Outcome int `json:"outcome"`
        }{
                Outcome: roll,
        }

        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(response); err != nil {
                http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
                return
        }
}
