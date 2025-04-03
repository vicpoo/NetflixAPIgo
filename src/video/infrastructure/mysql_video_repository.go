// mysql_video_repository.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vicpoo/NetflixAPIgo/src/core"
	"github.com/vicpoo/NetflixAPIgo/src/video/domain"
	"github.com/vicpoo/NetflixAPIgo/src/video/domain/entities"
)

type MySQLVideoRepository struct {
	conn *sql.DB
}

func NewMySQLVideoRepository() domain.VideoRepository {
	conn := core.GetDB()
	return &MySQLVideoRepository{conn: conn}
}

func (r *MySQLVideoRepository) Save(video *entities.Video) error {
	query := `
		INSERT INTO videos (title, description, url, user_id)
		VALUES (?, ?, ?, ?)
	`
	result, err := r.conn.Exec(
		query,
		video.Title,
		video.Description,
		video.URL,
		video.UserID,
	)
	if err != nil {
		log.Println("Error al guardar video:", err)
		return fmt.Errorf("error al crear el video")
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID:", err)
		return fmt.Errorf("error al obtener ID del video")
	}
	video.ID = int(id)
	return nil
}

func (r *MySQLVideoRepository) GetByID(id int) (*entities.Video, error) {
	query := `
		SELECT id, title, description, url, user_id 
		FROM videos 
		WHERE id = ?
	`
	row := r.conn.QueryRow(query, id)

	var video entities.Video
	if err := row.Scan(
		&video.ID,
		&video.Title,
		&video.Description,
		&video.URL,
		&video.UserID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("video con ID %d no encontrado", id)
		}
		log.Println("Error al buscar video por ID:", id, "error:", err)
		return nil, fmt.Errorf("error al obtener el video")
	}
	return &video, nil
}

func (r *MySQLVideoRepository) GetAll() ([]entities.Video, error) {
	query := `
        SELECT id, title, description, url, user_id 
        FROM videos
        ORDER BY id DESC
    `
	rows, err := r.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todos los videos:", err)
		return nil, fmt.Errorf("error al listar videos")
	}
	defer rows.Close()

	var videos []entities.Video
	for rows.Next() {
		var video entities.Video
		if err := rows.Scan(
			&video.ID,
			&video.Title,
			&video.Description,
			&video.URL,
			&video.UserID,
		); err != nil {
			log.Println("Error al escanear video:", err)
			continue
		}
		videos = append(videos, video)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de leer videos:", err)
		return nil, fmt.Errorf("error al procesar los videos")
	}

	return videos, nil
}
