package storage

import (
	"github.com/dairlair/tweetwatch/pkg/entity"
	"github.com/jackc/pgx"
)

func (storage *Storage) AddStream(streamInterface entity.StreamInterface) (createdStream entity.StreamInterface, err error) {
	tx, err := storage.connPool.Begin()
	if err != nil {
		return nil, pgError(err)
	}
	defer func() {
		if err != nil {
			pgRollback(tx)
		}
	}()

	createdStream, err = txAddStream(tx, streamInterface)

	if err := tx.Commit(); err != nil {
		return nil, pgError(err)
	}

	return createdStream, nil
}

func txAddStream(tx *pgx.Tx, stream entity.StreamInterface) (result entity.StreamInterface, err error) {
	const addStreamSQL = `
		INSERT INTO stream (
			topic_id
			, track
		) VALUES (
			$1, $2
		) RETURNING 
			stream_id
			, topic_id 
			, track
			, created_at
	`

	createdStream := entity.Stream{}
	if err := tx.QueryRow(
			addStreamSQL,
			stream.GetTopicID(),
			stream.GetTrack(),
		).Scan(
			&createdStream.ID,
			&createdStream.TopicID,
			&createdStream.Track,
			&createdStream.CreateAt,
		); err != nil {
		return nil, pgError(err)
	}

	result = &createdStream

	return result, nil
}

func txDeleteTopicStreams(tx *pgx.Tx, topicID int64) (streamIDs []int64, err error) {
	const sql = `DELETE FROM stream WHERE topic_id = $1 RETURNING stream_id`
	rows, err := tx.Query(sql, topicID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var topicID int64
		err := rows.Scan(&topicID)
		if err != nil {
			return nil, err
		}
		streamIDs = append(streamIDs, topicID)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return streamIDs, err
}

func txInsertTopicStreams(tx *pgx.Tx, topicID int64, streams []entity.StreamInterface) ([]entity.StreamInterface, error) {
	var createdStreams []entity.StreamInterface
	for _, stream := range streams {
		st := entity.Stream{
			TopicID: topicID,
			Track:   stream.GetTrack(),
		}
		createdStream, err := txAddStream(tx, &st)
		if err != nil {
			return nil, err
		}
		createdStreams = append(createdStreams, createdStream)
	}

	return createdStreams, nil
}

// GetStreams returns all existing streams
func (storage *Storage) GetStreams() (streams []entity.StreamInterface, err error) {
	const getStreamsSQL = `
		SELECT 
			stream_id
			, track
			, created_at
		FROM
			stream
	`

	rows, err := storage.connPool.Query(getStreamsSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var stream entity.Stream
		if err := rows.Scan(
			&stream.ID,
			&stream.Track,
			&stream.CreateAt,
		); err != nil {
			return nil, err
		}
		streams = append(streams, &stream)
	}

	return streams, nil
}

// GetStreams returns all existing streams
func (storage *Storage) GetTopicStreams(topicID int64) (streams []entity.StreamInterface, err error) {
	const getStreamsSQL = `
		SELECT 
			stream_id
			, track
			, created_at
		FROM
			stream
		WHERE topic_id = $1
	`

	rows, err := storage.connPool.Query(getStreamsSQL, topicID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var stream entity.Stream
		if err := rows.Scan(
			&stream.ID,
			&stream.Track,
			&stream.CreateAt,
		); err != nil {
			return nil, err
		}
		streams = append(streams, &stream)
	}

	return streams, nil
}

// GetStreams returns all active streams (streams with flag "is_active" = TRUE)
func (storage *Storage) GetActiveStreams() (streams []entity.StreamInterface, err error) {
	// @TODO Refactor when tweet table got is_active flag.
	return storage.GetStreams()
}