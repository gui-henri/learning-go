package avaliation

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
)

type AvaliationService interface {
	Save(ctx context.Context, data AvaliacaoRequest) error
	Export(ctx context.Context, id int, format string)
}

type avaliationService struct {
	repository   avaliationRepository
	gotenbergUrl string
	templates    *template.Template
}

func NewAvaliationService(r avaliationRepository, gotenbergUrl string, templateInternePath string) (*avaliationService, error) {

	tmpl, err := template.ParseFiles(templateInternePath)
	if err != nil {
		return nil, fmt.Errorf("could not parse template: %w", err)
	}

	return &avaliationService{
		repository:   r,
		gotenbergUrl: gotenbergUrl,
		templates:    tmpl,
	}, nil
}

func (s *avaliationService) Save(ctx context.Context, data AvaliacaoRequest) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = s.repository.Save(dataBytes)

	return err
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
