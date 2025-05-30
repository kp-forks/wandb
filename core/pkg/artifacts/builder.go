package artifacts

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/wandb/wandb/core/internal/hashencode"
	spb "github.com/wandb/wandb/core/pkg/service_go_proto"
	"google.golang.org/protobuf/proto"
)

type ArtifactBuilder struct {
	artifactRecord   *spb.ArtifactRecord
	isDigestUpToDate bool
}

func NewArtifactBuilder(artifactRecord *spb.ArtifactRecord) *ArtifactBuilder {
	artifactClone := proto.Clone(artifactRecord).(*spb.ArtifactRecord)
	builder := &ArtifactBuilder{
		artifactRecord: artifactClone,
	}
	builder.initDefaultManifest()
	return builder
}

func (b *ArtifactBuilder) initDefaultManifest() {
	if b.artifactRecord.Manifest != nil {
		return
	}
	b.artifactRecord.Manifest = &spb.ArtifactManifest{
		Version:       1,
		StoragePolicy: "wandb-storage-policy-v1",
		StoragePolicyConfig: []*spb.StoragePolicyConfigItem{{
			Key:       "storageLayout",
			ValueJson: "\"V2\"",
		}},
	}
}

func (b *ArtifactBuilder) AddData(name string, data any) error {
	filename, digest, size, err := WriteJSONToTempFileWithMetadata(data)
	if err != nil {
		return err
	}
	b.artifactRecord.Manifest.Contents = append(b.artifactRecord.Manifest.Contents,
		&spb.ArtifactManifestEntry{
			Path:      name,
			Digest:    digest,
			LocalPath: filename,
			Size:      size,
		})
	b.isDigestUpToDate = false
	return nil
}

func (b *ArtifactBuilder) AddFile(path string, name string) error {

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	// file size:
	var size int64
	if stat, err := file.Stat(); err == nil { // if NO error
		size = stat.Size()
	}

	defer func() {
		_ = file.Close()
	}()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	digest := hashencode.ComputeB64MD5(data)
	b.artifactRecord.Manifest.Contents = append(b.artifactRecord.Manifest.Contents,
		&spb.ArtifactManifestEntry{
			Path:      name,
			Digest:    digest,
			LocalPath: path,
			Size:      size,
		})
	b.isDigestUpToDate = false
	return nil
}

func (b *ArtifactBuilder) updateManifestDigest() {
	if b.isDigestUpToDate {
		return
	}
	manifest, err := NewManifestFromProto(b.artifactRecord.Manifest)
	if err != nil {
		panic("unable to create manifest (unexpected)")
	}
	manifestDigest := computeManifestDigest(&manifest)
	b.artifactRecord.Digest = manifestDigest
	b.isDigestUpToDate = true
}

func (b *ArtifactBuilder) GetArtifact() *spb.ArtifactRecord {
	b.updateManifestDigest()
	return b.artifactRecord
}

func computeManifestDigest(manifest *Manifest) string {
	type hashedEntry struct {
		name   string
		digest string
	}

	var entries []hashedEntry
	for name, entry := range manifest.Contents {
		entries = append(entries, hashedEntry{
			name:   name,
			digest: entry.Digest,
		})
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].name < entries[j].name
	})

	hasher := md5.New()
	_, _ = fmt.Fprintf(hasher, "wandb-artifact-manifest-v1\n")
	for _, entry := range entries {
		_, _ = fmt.Fprintf(hasher, "%s:%s\n", entry.name, entry.digest)
	}

	return hex.EncodeToString(hasher.Sum(nil))
}
