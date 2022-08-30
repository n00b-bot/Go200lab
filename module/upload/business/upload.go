package uploadbusiness

import (
	"bytes"
	"context"
	"fmt"
	"food/common"
	"food/component/uploadprovider"
	"image"
	_ "image/jpeg"
	"io"
	"log"
	"path"
	"strings"
	"time"
)

type uploadBusiness struct {
	provider uploadprovider.Provider
}

func NewUploadBus(provider uploadprovider.Provider) *uploadBusiness {
	return &uploadBusiness{
		provider: provider,
	}
}

func (u *uploadBusiness) Upload(ctx context.Context, data []byte, folder, Filename string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)
	w, h, err := getImageDimension(fileBytes)
	log.Println(err)
	if err != nil {
		return nil, common.ErrInvalidRequest(err)
	}
	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}
	fileExt := path.Ext(Filename)
	fileName := fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)
	img, err := u.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))
	if err != nil {
		return nil, err
	}
	img.Height = h
	img.Width = w
	img.Extension = fileExt
	return img, nil

}

func getImageDimension(r io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(r)
	if err != nil {
		return 0, 0, err
	}
	return img.Width, img.Height, nil
}
