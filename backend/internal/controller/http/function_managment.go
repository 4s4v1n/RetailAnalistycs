package http

import (
	"APG6/internal/entity/utils"
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (h *handler) FncGrowthOfAverageCheck(w http.ResponseWriter, r *http.Request) {
	role, ok := r.Context().Value("role_id").(uint8)
	if !ok {
		WriteError(w, errors.New("invalid type of role_id"), http.StatusUnauthorized)
		return
	}

	if !r.URL.Query().Has("method") {
		WriteError(w, errors.New("query parameter method is missing"), http.StatusBadRequest)
		return
	}
	method := strings.ToUpper(r.URL.Query().Get("method"))
	if method != "PERIOD" && method != "QUANTITY" {
		WriteError(w, errors.New("query parameter method is invalid"), http.StatusUnprocessableEntity)
		return
	}

	if !r.URL.Query().Has("first") {
		WriteError(w, errors.New("query parameter first is missing"), http.StatusBadRequest)
		return
	}
	first, err := utils.ParseDate(r.URL.Query().Get("first"))
	if err != nil {
		WriteError(w, fmt.Errorf("query parameter first is invalid: %w", err), http.StatusUnprocessableEntity)
		return
	}

	if !r.URL.Query().Has("last") {
		WriteError(w, errors.New("query parameter last is missing"), http.StatusBadRequest)
		return
	}
	last, err := utils.ParseDate(r.URL.Query().Get("last"))
	if err != nil {
		WriteError(w, fmt.Errorf("query parameter last is invalid: %w", err), http.StatusUnprocessableEntity)
		return
	}

	if !r.URL.Query().Has("number") {
		WriteError(w, errors.New("query parameter number is missing"), http.StatusBadRequest)
		return
	}
	number, err := strconv.ParseInt(r.URL.Query().Get("number"), 10, 32)
	if err != nil {
		WriteError(w, fmt.Errorf("query parameter number is invalid: %w", err), http.StatusUnprocessableEntity)
		return
	}

	if !r.URL.Query().Has("coefficient") {
		WriteError(w, errors.New("query parameter coefficient is missing"), http.StatusBadRequest)
		return
	}
	coefficient, err := strconv.ParseFloat(r.URL.Query().Get("coefficient"), 64)
	if err != nil {
		WriteError(w, fmt.Errorf("query parameter coefficient is invalid: %w", err),
			http.StatusUnprocessableEntity)
		return
	}

	if !r.URL.Query().Has("max_churn_rate") {
		WriteError(w, errors.New("query parameter max_churn_rate is missing"), http.StatusBadRequest)
		return
	}
	maxChurnRate, err := strconv.ParseFloat(r.URL.Query().Get("max_churn_rate"), 64)
	if err != nil {
		WriteError(w, fmt.Errorf("query parameter max_churn_rate is invalid: %w", err),
			http.StatusUnprocessableEntity)
		return
	}

	if !r.URL.Query().Has("max_discount_share") {
		WriteError(w, errors.New("query parameter max_discount_share is missing"), http.StatusBadRequest)
		return
	}
	maxDiscountShare, err := strconv.ParseFloat(r.URL.Query().Get("max_discount_share"), 64)
	if err != nil {
		WriteError(w, fmt.Errorf("query parameter max_discount_share is invalid: %w", err),
			http.StatusUnprocessableEntity)
		return
	}

	if !r.URL.Query().Has("margin_share") {
		WriteError(w, errors.New("query parameter margin_share is missing"), http.StatusBadRequest)
		return
	}
	marginShare, err := strconv.ParseFloat(r.URL.Query().Get("margin_share"), 64)
	if err != nil {
		WriteError(w, fmt.Errorf("query parameter margin_share is invalid: %w", err),
			http.StatusUnprocessableEntity)
		return
	}

	growth, err := h.functionManger.GrowthOfAverageCheck(context.Background(), role, method, first, last,
		int32(number), coefficient, maxChurnRate, maxDiscountShare, marginShare)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	accept := r.Header.Get("Accept")
	if accept == "application/json" {
		WriteResponseJson(w, growth)
	} else if accept == "text/csv; charset=UTF-8" || accept == "text/csv" {
		WriteResponseCsv(w, growth)
	} else {
		WriteError(w, errors.New("unsupported content type: "+accept), http.StatusUnsupportedMediaType)
	}
}

func (h *handler) FncDefiningOfferIncreasingFrequencyVisits(w http.ResponseWriter, r *http.Request) {
	role, ok := r.Context().Value("role_id").(uint8)
	if !ok {
		WriteError(w, errors.New("invalid type of role_id"), http.StatusUnauthorized)
		return
	}

	if !r.URL.Query().Has("first") {
		WriteError(w, errors.New("query parameter first is missing"), http.StatusBadRequest)
		return
	}
	first, err := utils.ParseDate(r.URL.Query().Get("first"))
	if err != nil {
		WriteError(w, fmt.Errorf("query parameter first is invalid: %w", err), http.StatusUnprocessableEntity)
		return
	}

	if !r.URL.Query().Has("last") {
		WriteError(w, errors.New("query parameter last is missing"), http.StatusBadRequest)
		return
	}
	last, err := utils.ParseDate(r.URL.Query().Get("last"))
	if err != nil {
		WriteError(w, fmt.Errorf("query parameter last is invalid: %w", err), http.StatusUnprocessableEntity)
		return
	}

	if !r.URL.Query().Has("value_transaction") {
		WriteError(w, errors.New("query parameter value_transaction is missing"), http.StatusBadRequest)
		return
	}
	valueTransaction, err := strconv.ParseInt(r.URL.Query().Get("value_transaction"), 10, 32)
	if err != nil {
		WriteError(w, fmt.Errorf("query parameter value_transaction is invalid: %w", err), http.StatusUnprocessableEntity)
		return
	}

	if !r.URL.Query().Has("max_churn_rate") {
		WriteError(w, errors.New("query parameter max_churn_rate is missing"), http.StatusBadRequest)
		return
	}
	maxChurnRate, err := strconv.ParseFloat(r.URL.Query().Get("max_churn_rate"), 64)
	if err != nil {
		WriteError(w, fmt.Errorf("query parameter max_churn_rate is invalid: %w", err),
			http.StatusUnprocessableEntity)
		return
	}

	if !r.URL.Query().Has("max_discount_share") {
		WriteError(w, errors.New("query parameter max_discount_share is missing"), http.StatusBadRequest)
		return
	}
	maxDiscountShare, err := strconv.ParseFloat(r.URL.Query().Get("max_discount_share"), 64)
	if err != nil {
		WriteError(w, fmt.Errorf("query parameter max_discount_share is invalid: %w", err),
			http.StatusUnprocessableEntity)
		return
	}

	if !r.URL.Query().Has("margin_share") {
		WriteError(w, errors.New("query parameter margin_share is missing"), http.StatusBadRequest)
		return
	}
	marginShare, err := strconv.ParseFloat(r.URL.Query().Get("margin_share"), 64)
	if err != nil {
		WriteError(w, fmt.Errorf("query parameter margin_share is invalid: %w", err),
			http.StatusUnprocessableEntity)
		return
	}

	increase, err := h.functionManger.DefiningOfferIncreasingFrequencyVisits(context.Background(), role, first, last,
		int32(valueTransaction), maxChurnRate, maxDiscountShare, marginShare)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	accept := r.Header.Get("Accept")
	if accept == "application/json" {
		WriteResponseJson(w, increase)
	} else if accept == "text/csv; charset=UTF-8" || accept == "text/csv" {
		WriteResponseCsv(w, increase)
	} else {
		WriteError(w, errors.New("unsupported content type: "+accept), http.StatusUnsupportedMediaType)
	}
}

func (h *handler) FncDefiningOfferIncreasingMargin(w http.ResponseWriter, r *http.Request) {
	role, ok := r.Context().Value("role_id").(uint8)
	if !ok {
		WriteError(w, errors.New("invalid type of role_id"), http.StatusUnauthorized)
		return
	}

	if !r.URL.Query().Has("count_group") {
		WriteError(w, errors.New("query parameter count_group is missing"), http.StatusBadRequest)
		return
	}
	countGroup, err := strconv.ParseInt(r.URL.Query().Get("count_group"), 10, 32)
	if err != nil {
		WriteError(w, fmt.Errorf("query parameter count_group is invalid: %w", err), http.StatusUnprocessableEntity)
		return
	}

	if !r.URL.Query().Has("max_churn_rate") {
		WriteError(w, errors.New("query parameter max_churn_rate is missing"), http.StatusBadRequest)
		return
	}
	maxChurnRate, err := strconv.ParseFloat(r.URL.Query().Get("max_churn_rate"), 64)
	if err != nil {
		WriteError(w, fmt.Errorf("query parameter max_churn_rate is invalid: %w", err),
			http.StatusUnprocessableEntity)
		return
	}

	if !r.URL.Query().Has("max_stability_index") {
		WriteError(w, errors.New("query parameter max_stability_index is missing"), http.StatusBadRequest)
		return
	}
	maxStabilityIndex, err := strconv.ParseFloat(r.URL.Query().Get("max_stability_index"), 64)
	if err != nil {
		WriteError(w, fmt.Errorf("query parameter max_stability_index is invalid: %w", err),
			http.StatusUnprocessableEntity)
		return
	}

	if !r.URL.Query().Has("max_index_sku") {
		WriteError(w, errors.New("query parameter max_index_sku is missing"), http.StatusBadRequest)
		return
	}
	maxIndexSku, err := strconv.ParseFloat(r.URL.Query().Get("max_index_sku"), 64)
	if err != nil {
		WriteError(w, fmt.Errorf("query parameter max_index_sku is invalid: %w", err),
			http.StatusUnprocessableEntity)
		return
	}

	if !r.URL.Query().Has("margin_share") {
		WriteError(w, errors.New("query parameter margin_share is missing"), http.StatusBadRequest)
		return
	}
	marginShare, err := strconv.ParseFloat(r.URL.Query().Get("margin_share"), 64)
	if err != nil {
		WriteError(w, fmt.Errorf("query parameter margin_share is invalid: %w", err),
			http.StatusUnprocessableEntity)
		return
	}

	increase, err := h.functionManger.DefiningOfferIncreasingMargin(context.Background(), role, int32(countGroup),
		maxChurnRate, maxStabilityIndex, maxIndexSku, marginShare)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	accept := r.Header.Get("Accept")
	if accept == "application/json" {
		WriteResponseJson(w, increase)
	} else if accept == "text/csv; charset=UTF-8" || accept == "text/csv" {
		WriteResponseCsv(w, increase)
	} else {
		WriteError(w, errors.New("unsupported content type: "+accept), http.StatusUnsupportedMediaType)
	}
}
