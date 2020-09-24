package blob

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/pkg/errors"

	"github.com/pivotal/kpack/pkg/archive"
)

var unexpectedBlobTypeError = errors.New("unexpected blob file type, must be one of .zip, .tar.gz, .tar, .jar")

type Fetcher struct {
	Logger *log.Logger
}

func (f *Fetcher) Fetch(dir string, blobURL string) error {
	file, err := downloadBlob(blobURL)
	if err != nil {
		return err
	}
	defer os.RemoveAll(file.Name())

	mediaType, err := classifyFile(file)
	if err != nil {
		return err
	}

	switch mediaType {
	case "application/zip":
		info, err := file.Stat()
		if err != nil {
			return err
		}
		err = archive.ExtractZip(file, info.Size(), dir)
	case "application/x-gzip":
		err = archive.ExtractTarGZ(file, dir)
	case "application/octet-stream":
		if !archive.IsTar(file.Name()) {
			return unexpectedBlobTypeError
		}
		err = archive.ExtractTar(file, dir)
	default:
		return unexpectedBlobTypeError
	}
	if err != nil {
		return err
	}

	u, err := url.Parse(blobURL)
	if err != nil {
		return err
	}
	f.Logger.Printf("Successfully downloaded %s%s in path %q", u.Host, u.Path, dir)

	return nil
}

func downloadBlob(blobURL string) (*os.File, error) {
	resp, err := http.Get(blobURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("failed to get blob %s", blobURL)
	}

	file, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return nil, err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func classifyFile(reader io.ReadSeeker) (string, error) {
	buf := make([]byte, 512)
	_, err := reader.Read(buf)
	if err != nil {
		return "", err
	}

	_, err = reader.Seek(0, 0)
	if err != nil {
		return "", err
	}

	return http.DetectContentType(buf), nil
}
