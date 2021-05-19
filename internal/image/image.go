package image

import "github.com/jinzhu/gorm"

// Service - the struct for our image service
type Service struct {
    DB *gorm.DB
}

// Image - defines our image structure
type Image struct {
    gorm.Model
    StorageUrl string
    AspectRatio string
}

// ImageService - the interface for an image service to conform
type ImageService interface {
    GetImage(ID uint) (Image, error)
    PostImage(image Image) (Image, error)
    UpdateImage(ID uint, newImage Image) (Image, error)
    DeleteImage(ID uint)
    GetAllImages() ([]Image, error)
}

// NewService - returns a new image service
func NewService(db *gorm.DB) *Service {
    return &Service{
        DB: db,
    }
}

// GetImage - retrieves image by its ID from the DB
func (s *Service) GetImage(ID uint) (Image, error) {
    var image Image
    if result := s.DB.First(&image, ID); result.Error != nil {
        return Image{}, result.Error
    }

    return image, nil
}

// GetAllImages - retrieves all image by its ID from the DB
func (s *Service) GetAllImages() ([]Image, error) {
    var images []Image
    if result := s.DB.First(&images); result.Error != nil {
        return []Image{}, result.Error
    }

    return images, nil
}

func (s *Service) PostImage(image Image) (Image, error) {
    if result := s.DB.Save(&image); result.Error != nil {
        return Image{}, result.Error
    }

    return image, nil
}

func (s *Service) UpdateImage(ID uint, newImage Image) (Image, error) {
    image, err := s.GetImage(ID)
    if err != nil {
        return Image{}, nil
    }

    if result := s.DB.Model(&image).Updates(newImage); result.Error != nil {
        return Image{}, result.Error
    }

    return image, nil
}

func (s *Service) DeleteImage(ID uint) error {
    if result := s.DB.Delete(&Image{}, ID); result.Error != nil {
        return result.Error
    }

    return nil
}
