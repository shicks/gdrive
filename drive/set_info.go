package drive

import (
	"fmt"
	"io"
	"strings"
)

type SetInfoArgs struct {
	Out         io.Writer
	Id          string
	Name        string
	Description string
	Parents     []string
	Mime        string
}

func (self *Drive) SetInfo(args SetInfoArgs) error {
	dstFile, err := self.service.Files.Get(args.Id).Fields("name", "mimeType").Do()
	if err != nil {
		return fmt.Errorf("Failed to get file: %s", err)
	}

	// Use provided file name or use filename
	if args.Name != "" {
		dstFile.Name = args.Name
	}

	// Set provided mime type or get type based on file extension
	if args.Mime != "" {
		dstFile.MimeType = args.Mime
	}

	// Set parent folders
	//remove := make([]string, 0, len(dstFile.Parents))
	//add := make([]string, 0, len(args.Parents))
	// if len(args.Parents) > 0 {
	// 	parents := make(map[string]int)
	// 	for _, p := range(args.Parents) {
	// 		parents[p] = 1
	// 	}
	// 	for _, p := range(dstFile.Parents) {
	// 		if parents[p] == 1 {
	// 			parents[p] = 0
	// 		} else {
	// 			remove = append(remove, p)
	// 		}
	// 	}
	// 	for p, v := range(parents) {
	// 		if v == 1 {
	// 			add = append(add, p)
	// 		}
	// 	}
	// }

	f, err := self.service.Files.Update(args.Id, dstFile).Fields("id", "name").AddParents(strings.Join(args.Parents, ",")).Do()
	if err != nil {
		return fmt.Errorf("Failed to update file: %s", err)
	}

	fmt.Fprintf(args.Out, "Updated %s: %s\n", f.Id, f.Name)
	return nil
}
