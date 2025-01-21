package wallet

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ArdiSasongko/EwalletProjects-user/internal/env"
	"github.com/joho/godotenv"
)

type WalletPayload struct {
	UserID int32 `json:"user_id"`
}

type WalletResponse struct {
	UserID    int32     `json:"user_id"`
	Balance   float32   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

type WalletClient interface {
	CreateWallet(ctx context.Context, userID int32) (*WalletResponse, error)
}

type walletClient struct {
	httpClient *http.Client
}

func NewWalletClient() WalletClient {
	return &walletClient{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (w *walletClient) CreateWallet(ctx context.Context, userID int32) (*WalletResponse, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	url := env.GetEnvString("WALLET_SERVICE", "") + env.GetEnvString("WALLET_BASE_PATH", "")

	payload := WalletPayload{
		UserID: userID,
	}

	//log.Println(url)

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal wallet payload :%w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request ;%w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := w.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create wallet :%w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body :%w", err)
	}

	//log.Println("response body", string(body))
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wallet service returned error :%s", string(body))
	}

	var walletResp WalletResponse
	if err := json.Unmarshal(body, &walletResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response :%w", err)
	}
	return &walletResp, nil
}
