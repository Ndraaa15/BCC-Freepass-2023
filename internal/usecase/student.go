package usecase

import (
	"bcc-freepass-2023/internal/repository"
	"bcc-freepass-2023/model"
	"bcc-freepass-2023/pkg/bcrypt"
	errcustom "bcc-freepass-2023/pkg/error"
	"context"
	"net/http"

	"github.com/google/uuid"
)

type IStudentUsecase interface {
	CreateStudent(ctx context.Context, request model.StudentRegister) (uuid.UUID, error)
}

type StudentUsecase struct {
	studentQuerier *repository.Queries
}

func NewStudentUsecase(studentQuerier *repository.Queries) IStudentUsecase {
	return &StudentUsecase{
		studentQuerier: studentQuerier,
	}
}

func (s *StudentUsecase) CreateStudent(ctx context.Context, request model.StudentRegister) (uuid.UUID, error) {
	hashedPassword, err := bcrypt.HashPassword(request.Password)
	if err != nil {
		return uuid.Nil, errcustom.NewCustomError(http.StatusBadRequest, "[CreateStudent] : hashing password", "failed hashing user password", err)
	}

	inserUserParam := repository.InsertUserParams{
		Email:          request.Email,
		HashedPassword: hashedPassword,
		FullName:       request.FullName,
		Nim:            request.NIM,
		Major:          request.Major,
		Faculty:        request.Faculty,
		Semester:       int32(request.Semester),
		Contact:        request.Contact,
		RoleID:         1,
	}

	id, err := s.studentQuerier.InsertUser(ctx, inserUserParam)
	if err != nil {
		return uuid.Nil, errcustom.NewCustomError(http.StatusInternalServerError, "[CreateStudent] : insert user into database", "Failed insert user into database", err)
	}

	return uuid.UUID(id.Bytes), nil
}
