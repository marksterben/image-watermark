package usecase

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"image"
	"image-watermark/domain"
	"image-watermark/helper"
	"image/jpeg"
	"mime/multipart"
	"net/http"
	"time"
)

type Usecase struct {
	ContextTimeout time.Duration
}

func (u *Usecase) AddWatermark(ctx context.Context, src multipart.File, req domain.Request) error {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	img, _, err := image.Decode(src)
	if err != nil {
		return err
	}

	img = helper.AddWatermark(img, req.Text)

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Membuat part untuk file gambar
	part, err := writer.CreateFormFile("file", "image.jpg")
	if err != nil {
		return err
	}

	// Encode gambar ke part tersebut
	err = jpeg.Encode(part, img, nil)
	if err != nil {
		return err
	}

	// Menutup writer untuk menambahkan boundary di akhir multipart data
	err = writer.Close()
	if err != nil {
		return err
	}

	reqClient, err := http.NewRequest("POST", fmt.Sprintf("https://api.dev.aslcode.dev/resources/public/%v", req.Folder), &buf)
	if err != nil {
		return err
	}

	reqClient.Header.Add("Authorization", req.Authorization)
	reqClient.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(reqClient)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return errors.New("Gagal mengirim gambar ke API eksternal")
	}

	return nil
}
