package mnistloader

import (
    "encoding/binary"
    "io"
    "errors"
    "os"
)

// LoadLabels parses input label file from given path and returns list of MNIST labels
func LoadLabels(path string) ([]float64, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    return readLabels(f)
}

// LoadImages parses input image file and returns list of MNIST images
func LoadImages(path string) ([][]float64, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    return readImages(f)
}

func readLabels(r io.Reader) ([]float64, error) {
    var magic int32
    var itemCount int32
    var item byte
    if err := binary.Read(r, binary.BigEndian, &magic); err != nil || magic != 2049 {
        return nil, errors.New("mnistloader: cannot read magic number in label file")
    }
    if err := binary.Read(r, binary.BigEndian, &itemCount); err != nil {
        return nil, errors.New("mnistloader: cannot read count of labels")
    }
    items := make([]float64, itemCount)
    for i := int32(0); i < itemCount; i++ {
        if err := binary.Read(r, binary.BigEndian, &item); err != nil {
            return nil, errors.New("mnistloader: cannot read label")
        }
        items[i] = float64(item)
    }
    return items, nil
}

func readImages(r io.Reader) ([][]float64, error) {
    var magic int32
    var itemCount int32
    var rows int32
    var cols int32
    if err := binary.Read(r, binary.BigEndian, &magic); err != nil || magic != 2051 {
        return nil, errors.New("mnistloader: cannot read magic number in images file")
    }
    if err := binary.Read(r, binary.BigEndian, &itemCount); err != nil {
        return nil, errors.New("mnistloader: cannot read count of images")
    }
    if err := binary.Read(r, binary.BigEndian, &rows); err != nil {
        return nil, errors.New("mnistloader: cannot read number of rows of images")
    }
    if err := binary.Read(r, binary.BigEndian, &cols); err != nil {
        return nil, errors.New("mnistloader: cannot read number of cols of images")
    }
    items := make([][]float64, itemCount)
    for i := int32(0); i < itemCount; i++ {
        item := make([]byte, rows*cols)
        items[i] = make([]float64, rows*cols)
        if err := binary.Read(r, binary.BigEndian, item); err != nil {
            return nil, errors.New("mnistloader: cannot read image")
        }
        for j, val := range item {
            items[i][j] = float64(val)
        }
    }
    return items, nil
}
