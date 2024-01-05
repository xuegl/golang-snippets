package json

import (
	"encoding/json"
	"net/http"
	"time"
)

type eventIdentifier struct {
	Event string `json:"event"`
}

type paymentPending struct {
	Event string `json:"event"`
	Data  struct {
		ID               int       `json:"id"`
		Domain           string    `json:"domain"`
		Amount           int       `json:"amount"`
		Currency         string    `json:"currency"`
		DueDate          any       `json:"due_date"`
		HasInvoice       bool      `json:"has_invoice"`
		InvoiceNumber    any       `json:"invoice_number"`
		Description      string    `json:"description"`
		PdfURL           any       `json:"pdf_url"`
		LineItems        []any     `json:"line_items"`
		Tax              []any     `json:"tax"`
		RequestCode      string    `json:"request_code"`
		Status           string    `json:"status"`
		Paid             bool      `json:"paid"`
		PaidAt           any       `json:"paid_at"`
		Metadata         any       `json:"metadata"`
		Notifications    []any     `json:"notifications"`
		OfflineReference string    `json:"offline_reference"`
		Customer         int       `json:"customer"`
		CreatedAt        time.Time `json:"created_at"`
	} `json:"data"`
}

type paymentSuccess struct {
	Event string `json:"event"`
	Data  struct {
		ID            int       `json:"id"`
		Domain        string    `json:"domain"`
		Amount        int       `json:"amount"`
		Currency      string    `json:"currency"`
		DueDate       any       `json:"due_date"`
		HasInvoice    bool      `json:"has_invoice"`
		InvoiceNumber any       `json:"invoice_number"`
		Description   string    `json:"description"`
		PdfURL        any       `json:"pdf_url"`
		LineItems     []any     `json:"line_items"`
		Tax           []any     `json:"tax"`
		RequestCode   string    `json:"request_code"`
		Status        string    `json:"status"`
		Paid          bool      `json:"paid"`
		PaidAt        time.Time `json:"paid_at"`
		Metadata      any       `json:"metadata"`
		Notifications []struct {
			SentAt  time.Time `json:"sent_at"`
			Channel string    `json:"channel"`
		} `json:"notifications"`
		OfflineReference string    `json:"offline_reference"`
		Customer         int       `json:"customer"`
		CreatedAt        time.Time `json:"created_at"`
	} `json:"data"`
}

func HandleDynamicAPI(w http.ResponseWriter, r *http.Request) {
	var (
		eventIdentifier eventIdentifier
		jsonData        json.RawMessage
	)

	if err := json.NewDecoder(r.Body).Decode(&jsonData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(jsonData, &eventIdentifier); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch eventIdentifier.Event {
	case "payment.pending":
		var paymentPending paymentPending
		if err := json.Unmarshal(jsonData, &paymentPending); err != nil {
			return
		}

		if err := json.NewEncoder(w).Encode(map[string]any{
			"event type": paymentPending.Event,
			"amount":     paymentPending.Data.Amount,
		}); err != nil {
			return
		}
		w.Header().Add("Content-Type", "application/json")
	case "payment.success":
		var paymentSuccess paymentSuccess
		if err := json.Unmarshal(jsonData, &paymentSuccess); err != nil {
			return
		}
		if err := json.NewEncoder(w).Encode(map[string]string{
			"event type": paymentSuccess.Event,
		}); err != nil {
			return
		}
		w.Header().Add("Content-Type", "application/json")
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
