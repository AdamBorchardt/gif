package image

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"
)

func (img *Image) Print() {
	writer := tabwriter.NewWriter(os.Stdout, 4, 4, 2, ' ', 0)
	defer writer.Flush()

	img.PrintTo(writer)
}

func PrintAll(images []Image) {
	writer := tabwriter.NewWriter(os.Stdout, 4, 4, 2, ' ', 0)
	defer writer.Flush()

	for _, img := range images {
		img.PrintTo(writer)
	}
}

func (img *Image) PrintTo(writer io.Writer) {
	fmt.Fprintf(writer, "%s\t", img.Id[:8])

	if img.Url == "" {
		io.WriteString(writer, "local\t")
	} else {
		io.WriteString(writer, "remote\t")
	}

	fmt.Fprintf(writer, "%v\t", img.AddedAt.Format("2006-01-02 15:04:05"))

	if len(img.Tags) > 0 {
		fmt.Fprintln(writer, strings.Join(img.Tags, ", "))
	} else {
		fmt.Fprintln(writer, "(no tags)")
	}
}