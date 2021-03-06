package file

import (
	"bytes"
	"compress/gzip"
	"testing"
	"time"

	"github.com/GGP1/kure/crypt"
	"github.com/GGP1/kure/pb"

	"github.com/awnumar/memguard"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	bolt "go.etcd.io/bbolt"
)

func TestFile(t *testing.T) {
	db := setContext(t)

	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)

	if _, err := gw.Write([]byte("content")); err != nil {
		t.Fatalf("Failed compressing content")
	}

	if err := gw.Close(); err != nil {
		t.Fatalf("Failed closing gzip writer")
	}

	f := &pb.File{
		Name:      "test",
		Content:   buf.Bytes(),
		CreatedAt: 0,
	}

	t.Run("Create", create(db, f))
	t.Run("Get", get(db, f.Name))
	t.Run("Get cheap", getCheap(db, f.Name))
	t.Run("List", list(db))
	t.Run("List names", listNames(db))
	t.Run("Remove", remove(db, f.Name))
}

func create(db *bolt.DB, f *pb.File) func(*testing.T) {
	return func(t *testing.T) {
		if err := Create(db, f); err != nil {
			t.Error(err)
		}
	}
}

func get(db *bolt.DB, name string) func(*testing.T) {
	return func(t *testing.T) {
		got, err := Get(db, name)
		if err != nil {
			t.Error(err)
		}

		// They aren't DeepEqual
		if got.Name != name {
			t.Errorf("Expected %s, got %s", name, got.Name)
		}
	}
}

func getCheap(db *bolt.DB, name string) func(*testing.T) {
	return func(t *testing.T) {
		got, err := GetCheap(db, name)
		if err != nil {
			t.Error(err)
		}

		// They aren't DeepEqual
		if got.Name != name {
			t.Errorf("Expected %s, got %s", name, got.Name)
		}
	}
}

func list(db *bolt.DB) func(*testing.T) {
	return func(t *testing.T) {
		files, err := List(db)
		if err != nil {
			t.Error(err)
		}

		if len(files) == 0 {
			t.Error("Expected one or more files, got 0")
		}
	}
}

func listNames(db *bolt.DB) func(*testing.T) {
	return func(t *testing.T) {
		files, err := ListNames(db)
		if err != nil {
			t.Error(err)
		}
		if len(files) == 0 {
			t.Error("Expected one or more files, got 0")
		}

		expected := "test"
		got := files[0]

		if got != expected {
			t.Errorf("Expected %s, got %s", expected, got)
		}
	}
}

func remove(db *bolt.DB, name string) func(*testing.T) {
	return func(t *testing.T) {
		if err := Remove(db, name); err != nil {
			t.Error(err)
		}
	}
}

func TestCreateErrors(t *testing.T) {
	db := setContext(t)

	if err := Create(db, &pb.File{}); err == nil {
		t.Error("Expected 'save file' error, got nil")
	}
}

func TestGetError(t *testing.T) {
	db := setContext(t)

	if _, err := Get(db, "non-existent"); err == nil {
		t.Error("Expected 'does not exist' error, got nil")
	}
}

func TestGetCheapError(t *testing.T) {
	db := setContext(t)

	if _, err := GetCheap(db, "non-existent"); err == nil {
		t.Error("Expected 'does not exist' error, got nil")
	}
}

func TestCryptErrors(t *testing.T) {
	db := setContext(t)

	name := "crypt-errors"
	if err := Create(db, &pb.File{Name: name}); err != nil {
		t.Fatal(err)
	}

	// Try to get the file with other password
	viper.Set("auth.password", memguard.NewEnclave([]byte("invalid")))

	if _, err := Get(db, name); err == nil {
		t.Error("Expected Get() to fail but it didn't")
	}
	if _, err := GetCheap(db, name); err == nil {
		t.Error("Expected GetCheap() to fail but it didn't")
	}
	if _, err := List(db); err == nil {
		t.Error("Expected List() to fail but it didn't")
	}
}

func TestProtoErrors(t *testing.T) {
	db := setContext(t)

	name := "unformatted"
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(fileBucket))
		buf := make([]byte, 64)
		encBuf, _ := crypt.Encrypt(buf)
		return b.Put([]byte(name), encBuf)
	})
	if err != nil {
		t.Fatalf("Failed writing invalid type: %v", err)
	}

	if _, err := Get(db, name); err == nil {
		t.Error("Expected Get() to fail but it didn't")
	}
	if _, err := GetCheap(db, name); err == nil {
		t.Error("Expected GetCheap() to fail but it didn't")
	}
	if _, err := List(db); err == nil {
		t.Error("Expected List() to fail but it didn't")
	}
}

func TestKeyError(t *testing.T) {
	db := setContext(t)

	if err := Create(db, &pb.File{}); err == nil {
		t.Error("Create() didn't fail")
	}
}

func TestListNameNil(t *testing.T) {
	db := setContext(t)
	err := db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket(fileBucket)
	})
	if err != nil {
		t.Fatalf("Failed deleting the file bucket: %v", err)
	}

	list, err := ListNames(db)
	if err != nil || list != nil {
		t.Errorf("Expected to receive a nil list and error, got: %v list, %v error", list, err)
	}
}

func setContext(t *testing.T) *bolt.DB {
	db, err := bolt.Open("../testdata/database", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		t.Fatalf("Failed connecting to the database: %v", err)
	}

	viper.Reset()
	// Reduce argon2 parameters to speed up tests
	auth := map[string]interface{}{
		"password":   memguard.NewEnclave([]byte("1")),
		"iterations": 1,
		"memory":     1,
		"threads":    1,
	}
	viper.Set("auth", auth)

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := "kure_file"
		tx.DeleteBucket([]byte(bucket))
		if _, err := tx.CreateBucketIfNotExists([]byte(bucket)); err != nil {
			return errors.Wrapf(err, "couldn't create %q bucket", bucket)
		}

		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := db.Close(); err != nil {
			t.Fatalf("Failed closing the database: %v", err)
		}
	})

	return db
}
