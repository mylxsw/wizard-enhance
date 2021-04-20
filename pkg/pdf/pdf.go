package pdf

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mylxsw/go-toolkit/file"
	"github.com/mylxsw/go-utils/str"
	"github.com/thecodingmachine/gotenberg-go-client/v7"
)

func OfficeToPDF(gotenbergServer string, storageRoot string, source string) (string, error) {
	if strings.HasSuffix(source, ".pdf") {
		return source, nil
	}

	dstPath := filepath.Join(filepath.Dir(source), "_cache")
	absPath := filepath.Join(storageRoot, dstPath)
	if !file.Exist(absPath) {
		_ = os.MkdirAll(absPath, os.ModePerm)
	}

	dstHashName := hash(source) + ".pdf"
	saveFilepath := filepath.Join(absPath, dstHashName)
	if file.Exist(saveFilepath) {
		return filepath.Join(dstPath, dstHashName), nil
	}

	c := &gotenberg.Client{Hostname: gotenbergServer}
	doc, err := gotenberg.NewDocumentFromPath(filepath.Base(source), filepath.Join(storageRoot, source))
	if err != nil {
		return "", err
	}

	officeReq := gotenberg.NewOfficeRequest(doc)
	officeReq.WaitTimeout(30.0)
	officeReq.Landscape(shouldLandscape(source))

	if err := c.Store(officeReq, saveFilepath); err != nil {
		return "", err
	}

	return filepath.Join(dstPath, dstHashName), nil
}

func shouldLandscape(source string) bool {
	return str.HasSuffixes(source, []string{".xlsx", "xls"})
}

func hash(ss ...string) string {
	sum := md5.Sum([]byte(strings.Join(ss, "-")))
	return fmt.Sprintf("%x", sum)
}
