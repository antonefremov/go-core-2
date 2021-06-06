package filestore_test

import (
	"encoding/json"
	"fmt"
	"go-core-2/homeworks/06-gosearch-v4/pkg/crawler"
	"go-core-2/homeworks/06-gosearch-v4/pkg/storage/filestore"
	"io"
	"testing"
)

type fakeWriter struct {
	success bool
}

func (fw *fakeWriter) Write(p []byte) (n int, err error) {
	if fw.success {
		return 10, nil
	}
	return 0, fmt.Errorf("Error when saving")
}

func TestFilestore_Save(t *testing.T) {
	type args struct {
		fw fakeWriter
		d  []crawler.Document
	}
	docs := []crawler.Document{{ID: 10}, {ID: 20}}
	tests := []struct {
		name       string
		args       args
		wantErr    error
		wantExists bool
	}{
		{
			name: "Save succeeds",
			args: args{
				fw: fakeWriter{success: true},
				d:  docs,
			},
			wantErr: nil,
		}, {
			name: "Save fails",
			args: args{
				fw: fakeWriter{success: false},
				d:  docs,
			},
			wantErr: fmt.Errorf("Error when saving"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := filestore.New()

			gotErr := fs.Save(&tt.args.fw, tt.args.d)
			if tt.wantErr != nil && gotErr.Error() != tt.wantErr.Error() {
				t.Errorf("Save() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

type fakeReader struct {
	success bool
	docs    []crawler.Document
	done    bool
}

func (fr *fakeReader) Read(p []byte) (n int, err error) {
	if fr.success {
		d, _ := json.Marshal(fr.docs)
		copy(p, d)
		if fr.done {
			return 0, io.EOF
		}
		fr.done = true
		return len(d), nil
	}
	return 0, fmt.Errorf("Error when reading")
}

func TestFilestore_Retrieve(t *testing.T) {
	type args struct {
		fr fakeReader
		d  []crawler.Document
	}
	docs := []crawler.Document{{ID: 10}, {ID: 20}}
	tests := []struct {
		name       string
		args       args
		wantErr    error
		wantExists bool
	}{
		{
			name: "Read succeeds",
			args: args{
				fr: fakeReader{success: true, docs: docs},
			},
			wantErr: nil,
		}, {
			name: "Read fails",
			args: args{
				fr: fakeReader{success: false},
				d:  docs,
			},
			wantErr: fmt.Errorf("Error when reading"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := filestore.New()
			docs, gotErr := fs.Retrieve(&tt.args.fr)
			if tt.wantErr == nil {
				if gotErr != nil {
					t.Errorf("Received unexpected error = %v", gotErr.Error())
				}
				if len(docs) != len(tt.args.fr.docs) {
					t.Errorf("Retrieve() gotLen = %v, want %v", len(docs), len(tt.args.fr.docs))
				}
			}
			if tt.wantErr != nil && gotErr.Error() != tt.wantErr.Error() {
				t.Errorf("Retrieve() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
