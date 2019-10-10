package avatar

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/leeif/pluto/config"
	perror "github.com/leeif/pluto/datatype/pluto_error"
)

type Avatar struct {
	Bucket          string
	EndPoint        string
	AccessKeyID     string
	AccessKeySecret string
	CDN             string
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

type AvatarReader struct {
	Reader    io.ReadCloser
	Ext       string
	OriginURL string
}

func (avatar *Avatar) GetRandomAvatar() (string, *perror.PlutoError) {
	ar := &AvatarReader{}
	originURL := fmt.Sprintf("https://www.gravatar.com/avatar/%s?f=y&d=identicon", randToken(8))
	resp, err := http.Get(originURL)
	if err != nil {
		return "", perror.ServerError.Wrapper(err)
	}
	ar.Reader = resp.Body
	ar.OriginURL = originURL
	contentType := resp.Header.Get("Content-type")
	if contentType == "image/png" {
		ar.Ext = "png"
	} else if contentType == "image/jpg" {
		ar.Ext = "jpg"
	} else {
		return "", perror.ServerError.Wrapper(errors.New("Not support type of avatar " + contentType))
	}
	ossURL, perr := avatar.saveAvatarImageInOSS(ar)
	if perr != nil {
		return originURL, perr
	}
	return ossURL, nil
}

func (avatar *Avatar) saveAvatarImageInOSS(reader *AvatarReader) (string, *perror.PlutoError) {

	if avatar.AccessKeyID == "" ||
		avatar.AccessKeySecret == "" ||
		avatar.Bucket == "" ||
		avatar.EndPoint == "" {
		return "", perror.ServerError.Wrapper(errors.New("aliyun config is not enough"))
	}

	client, err := oss.New(avatar.EndPoint, avatar.AccessKeyID, avatar.AccessKeySecret)
	if err != nil {
		return "", perror.ServerError.Wrapper(err)
	}
	bucket, err := client.Bucket(avatar.Bucket)
	if err != nil {
		return "", perror.ServerError.Wrapper(err)
	}
	osskey := fmt.Sprintf("avatar/%s.%s", randToken(8), reader.Ext)
	err = bucket.PutObject(osskey, reader.Reader)
	if err != nil {
		return "", perror.ServerError.Wrapper(err)
	}
	url := ""
	if avatar.CDN == "" {
		url = fmt.Sprintf("https://%s.%s/%s", avatar.Bucket, avatar.EndPoint, osskey)
	} else {
		url = fmt.Sprintf("%s/%s", avatar.CDN, osskey)
	}
	return url, nil
}

func (avatar *Avatar) SaveAvatarImageInLocal(file io.ReadCloser) (string, error) {
	return "", nil
}

func NewAvatar(config *config.Config) *Avatar {
	avatar := &Avatar{
		Bucket:          *config.Avatar.Bucket,
		EndPoint:        *config.Avatar.EndPoint,
		AccessKeyID:     *config.Avatar.AccessKeyID,
		AccessKeySecret: *config.Avatar.AccessKeySecret,
		CDN:             *config.Avatar.CDN,
	}
	return avatar
}
