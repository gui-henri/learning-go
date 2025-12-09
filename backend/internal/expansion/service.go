package expansion

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gui-henri/learning-go/internal/expansion/domain"
)

type AvaliationService interface {
	Save(ctx context.Context, data domain.AvaliacaoRequest) error
	Export(ctx context.Context, id int, format string) ([]byte, error)
	List(ctx context.Context) ([]domain.AvaliacaoRequest, error)
}

type avaliationService struct {
	repository   AvaliationRepository
	gotenbergUrl string
	templates    *template.Template
}

func NewAvaliationService(r AvaliationRepository, gotenbergUrl string, templateInternePath string) (*avaliationService, error) {

	funcMap := template.FuncMap{
		"formatDate": func(val interface{}) string {
			if val == nil {
				return ""
			}

			// 1. Se o valor for uma STRING (o caso atual do seu erro)
			if s, ok := val.(string); ok {
				if s == "" {
					return ""
				}
				// Tenta fazer o parse do formato ISO que vem do JSON (ex: 2025-12-10T03:00:00.000Z)
				t, err := time.Parse(time.RFC3339, s)
				if err == nil {
					return t.Format("02/01/2006")
				}

				// Tenta fazer o parse de data simples (ex: 2025-12-10)
				t, err = time.Parse("2006-01-02", s)
				if err == nil {
					return t.Format("02/01/2006")
				}

				// Se não conseguir converter, retorna a string original para não deixar em branco
				return s
			}

			// 2. Se o valor for *time.Time (Ponteiro)
			if t, ok := val.(*time.Time); ok {
				if t == nil || t.IsZero() {
					return ""
				}
				return t.Format("02/01/2006")
			}

			// 3. Se o valor for time.Time (Valor direto)
			if t, ok := val.(time.Time); ok {
				if t.IsZero() {
					return ""
				}
				return t.Format("02/01/2006")
			}

			return ""
		},
	}

	baseName := filepath.Base(templateInternePath)
	tmpl, err := template.New(baseName).Funcs(funcMap).ParseFiles(templateInternePath)
	if err != nil {
		return nil, fmt.Errorf("could not parse template: %w", err)
	}

	return &avaliationService{
		repository:   r,
		gotenbergUrl: gotenbergUrl,
		templates:    tmpl,
	}, nil
}

func (s *avaliationService) Save(ctx context.Context, data domain.AvaliacaoRequest) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = s.repository.Save(dataBytes)

	return err
}

func (s *avaliationService) List(ctx context.Context) ([]domain.AvaliacaoRequest, error) {
	return s.repository.GetAll()
}

func (s *avaliationService) Export(ctx context.Context, id int, format string) ([]byte, error) {

	if format != "interne" {
		return nil, fmt.Errorf("formato inválido")
	}

	var renderedHTML bytes.Buffer

	avaliation, err := s.repository.GetOne(id)

	if err != nil {
		return nil, err
	}

	if err := s.templates.ExecuteTemplate(&renderedHTML, "interne.html", avaliation); err != nil {
		return nil, fmt.Errorf("failed to render template: %w", err)
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("index.html", "index.html")
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}

	if _, err := part.Write(renderedHTML.Bytes()); err != nil {
		return nil, fmt.Errorf("failed to write html to form: %w", err)
	}

	// Optional: Add Gotenberg configurations (e.g., margins, paper size)
	// writer.WriteField("marginTop", "0")
	// writer.WriteField("marginBottom", "0")

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close multipart writer: %w", err)
	}

	url := fmt.Sprintf("%s/forms/chromium/convert/html", s.gotenbergUrl)

	req, err := http.NewRequestWithContext(ctx, "POST", url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call gotenberg: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("gotenberg failed with status %d: %s", resp.StatusCode, string(errBody))
	}

	pdfBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read pdf response: %w", err)
	}

	return pdfBytes, nil

}
