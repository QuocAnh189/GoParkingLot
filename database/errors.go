package database

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
)

var (
	ErrRecordNotFound        = gorm.ErrRecordNotFound        // Không tìm thấy bản ghi
	ErrInvalidTransaction    = gorm.ErrInvalidTransaction    // Transaction không hợp lệ
	ErrNotImplemented        = gorm.ErrNotImplemented        // Tính năng chưa được hỗ trợ
	ErrMissingWhereClause    = gorm.ErrMissingWhereClause    // Câu lệnh cập nhật hoặc xóa thiếu điều kiện WHERE
	ErrUnsupportedRelation   = gorm.ErrUnsupportedRelation   // Quan hệ không được hỗ trợ
	ErrPrimaryKeyRequired    = gorm.ErrPrimaryKeyRequired    // Thiếu khóa chính khi cập nhật hoặc xóa
	ErrModelValueRequired    = gorm.ErrModelValueRequired    // Thiếu giá trị model
	ErrInvalidData           = gorm.ErrInvalidData           // Dữ liệu đầu vào không hợp lệ
	ErrUnsupportedDriver     = gorm.ErrUnsupportedDriver     // Driver không được hỗ trợ
	ErrRegistered            = gorm.ErrRegistered            // Model đã được đăng ký
	ErrInvalidField          = gorm.ErrInvalidField          // Trường không hợp lệ
	ErrEmptySlice            = gorm.ErrEmptySlice            // Mảng đầu vào rỗng
	ErrDryRunModeUnsupported = gorm.ErrDryRunModeUnsupported // Chế độ Dry Run không được hỗ trợ
)

// Custom PostgresSQL errors
var (
	ErrUniqueViolation     = errors.New("pq: duplicate key value violates unique constraint")            // Vi phạm ràng buộc UNIQUE
	ErrForeignKeyViolation = errors.New("pq: insert or update on table violates foreign key constraint") // Vi phạm ràng buộc khóa ngoại
	ErrCheckViolation      = errors.New("pq: new row for relation violates check constraint")            // Vi phạm ràng buộc CHECK
	ErrNotNullViolation    = errors.New("pq: null value in column violates not-null constraint")         // Cột không được NULL
)

var ErrorStatusMap = map[error]int{
	ErrRecordNotFound:        http.StatusNotFound,
	ErrInvalidTransaction:    http.StatusBadRequest,
	ErrNotImplemented:        http.StatusNotImplemented,
	ErrMissingWhereClause:    http.StatusBadRequest,
	ErrUnsupportedRelation:   http.StatusBadRequest,
	ErrPrimaryKeyRequired:    http.StatusBadRequest,
	ErrModelValueRequired:    http.StatusBadRequest,
	ErrInvalidData:           http.StatusUnprocessableEntity,
	ErrUnsupportedDriver:     http.StatusInternalServerError,
	ErrRegistered:            http.StatusConflict,
	ErrInvalidField:          http.StatusBadRequest,
	ErrEmptySlice:            http.StatusBadRequest,
	ErrDryRunModeUnsupported: http.StatusInternalServerError,
	ErrUniqueViolation:       http.StatusConflict,
	ErrForeignKeyViolation:   http.StatusBadRequest,
	ErrCheckViolation:        http.StatusBadRequest,
	ErrNotNullViolation:      http.StatusBadRequest,
}
